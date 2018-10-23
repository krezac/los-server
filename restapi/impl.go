package restapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	interpose "github.com/carbocation/interpose/middleware"
	"github.com/dgrijalva/jwt-go"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/krezac/los-server/auth"
	"github.com/krezac/los-server/database"
	"github.com/krezac/los-server/models"
	"github.com/krezac/los-server/restapi/operations"
	"github.com/krezac/los-server/restapi/operations/competition"
	"github.com/krezac/los-server/restapi/operations/login"
	"github.com/krezac/los-server/restapi/operations/range_operations"
	"golang.org/x/crypto/bcrypt"
)

var logrusHandler func(http.Handler) http.Handler

//hash implements root.Hash
type hash struct{}

//Generate a salted hash for the input string
func (c *hash) Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

//Compare string to generated hash
func (c *hash) Compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}

func losAuthImpl(api *operations.LosAPI, token string, scopes []string) (*models.Principal, error) {
	api.Logger("losAuthImpl")
	principal, err := auth.HasRole(token, scopes)
	if principal != nil {
		api.Logger("have principal", principal.Name)
		if db.IsTokenInvalid(principal.Name, principal.RawToken) {
			return nil, fmt.Errorf("invalid token")
		}
	}
	return principal, err
}

func generateToken(login, password string, isLogin bool) (*models.LoginResponse, error) {
	dbu, err := db.GetUserByLogin(login, true)
	if err != nil {
		return nil, err
	}

	if isLogin { // you don't have password for refresh
		h := hash{}
		if err := h.Compare(dbu.Password, password); err != nil {
			return nil, err
		}
	}

	// user authenticated, let's generate the token
	expirationTime := time.Now().Add(time.Hour * 24 * 365) // TODO adjust the expiration
	roles := []string{"user"}                              // every logged-in user has this permission
	if dbu.RoleCompetitor {
		roles = append(roles, "competitor")
	}
	if dbu.RoleJudge {
		roles = append(roles, "judge")
	}
	if dbu.RoleDirector {
		roles = append(roles, "director")
	}
	if dbu.RoleAdmin {
		roles = append(roles, "admin")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"jti":   login,
		"iss":   auth.JwtExtraOptionsVar.JwtIssuerName,
		"exp":   expirationTime.Unix(),
		"roles": roles,
	})

	signBytes, err := ioutil.ReadFile(auth.JwtExtraOptionsVar.JwtSigningKey)
	if err != nil {
		return nil, err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, err
	}

	t, err := token.SignedString(signKey)

	if err != nil {
		return nil, err
	}

	resp := models.LoginResponse{
		Token:   t,
		ValidTo: strfmt.DateTime(expirationTime),
	}

	return &resp, nil
}

func loginUser(api *operations.LosAPI, params login.LoginUserParams) middleware.Responder {
	if params.Body == nil || params.Body.Login == "" {
		return login.NewLoginUserUnauthorized()
	}

	resp, err := generateToken(params.Body.Login, params.Body.Password, true)
	if err != nil {
		resp := models.APIResponse{
			Message: err.Error(), // TODO remove from final version (we don't want to disclose what failed)
		}
		return login.NewLoginUserUnauthorized().WithPayload(&resp)
	}

	return login.NewLoginUserOK().WithPayload(resp)
}

func logoutUser(api *operations.LosAPI, params login.LogoutUserParams, principal *models.Principal) middleware.Responder {
	err := db.InvalidateToken(principal.Name, principal.RawToken, time.Time(principal.ValidTo))
	api.Logger("InvalidateToken", err)
	if err != nil {
		resp := models.APIResponse{
			Message: err.Error(), // TODO remove from final version (we don't want to disclose what failed)
		}
		return login.NewLogoutUserUnauthorized().WithPayload(&resp)
	}

	return login.NewLogoutUserOK()
}

func refreshToken(api *operations.LosAPI, params login.RefreshTokenParams, principal *models.Principal) middleware.Responder {
	// generate new
	resp, err := generateToken(principal.Name, "", false) // no password here

	if err != nil {
		resp := models.APIResponse{
			Message: err.Error(), // TODO remove from final version (we don't want to disclose what failed)
		}
		return login.NewRefreshTokenUnauthorized().WithPayload(&resp)
	}

	err = db.InvalidateToken(principal.Name, principal.RawToken, time.Time(principal.ValidTo))
	if err != nil {
		resp := models.APIResponse{
			Message: err.Error(), // TODO remove from final version (we don't want to disclose what failed)
		}
		return login.NewRefreshTokenUnauthorized().WithPayload(&resp)
	}

	return login.NewRefreshTokenOK().WithPayload(resp)
}

func getRangeByID(api *operations.LosAPI, params range_operations.GetRangeByIDParams) middleware.Responder {
	dbr, err := db.GetRangeByID(params.RangeID)
	if err != nil {
		resp := models.APIResponse{
			Message: err.Error(),
		}
		return range_operations.NewGetRangeByIDNotFound().WithPayload(&resp)
	}

	return range_operations.NewGetRangeByIDOK().WithPayload(dbRangeToRange(dbr))
}

func getRanges(api *operations.LosAPI, params range_operations.GetRangesParams) middleware.Responder {
	dbRanges, err := db.GetRanges(*params.ActiveOnly)
	if err != nil {
		resp := models.APIResponse{
			Message: err.Error(),
		}
		return range_operations.NewGetRangesInternalServerError().WithPayload(&resp)
	}

	ranges := []*models.Range{}
	for _, dbr := range dbRanges {
		ranges = append(ranges, dbRangeToRange(&dbr))
	}

	return range_operations.NewGetRangesOK().WithPayload(ranges)
}

func getRangesHTML(api *operations.LosAPI, params range_operations.GetRangesHTMLParams) middleware.Responder {
	dbRanges, err := db.GetRanges(*params.ActiveOnly)
	if err != nil {
		resp := models.APIResponse{
			Message: err.Error(),
		}
		return range_operations.NewGetRangesInternalServerError().WithPayload(&resp)
	}

	ranges := []*models.Range{}
	for _, dbr := range dbRanges {
		ranges = append(ranges, dbRangeToRange(&dbr))
	}

	htmlRes := models.HTMLResponse{
		Template: "templates/ranges_html.gotmpl",
		Payload:  ranges,
	}
	return range_operations.NewGetRangesHTMLOK().WithPayload(&htmlRes)
}

func getCompetitionByID(api *operations.LosAPI, params competition.GetCompetitionByIDParams) middleware.Responder {
	dbc, err := db.GetCompetitionByID(params.CompetitionID)
	if err != nil {
		resp := models.APIResponse{
			Message: err.Error(),
		}
		return competition.NewGetCompetitionByIDNotFound().WithPayload(&resp)
	}

	return competition.NewGetCompetitionByIDOK().WithPayload(dbCompetitionToCompetition(dbc))
}

func getCompetitions(api *operations.LosAPI, params competition.GetCompetitionsParams) middleware.Responder {
	dbCompetitions, err := db.GetCompetitions(params.RangeID, *params.ActiveOnly)
	if err != nil {
		resp := models.APIResponse{
			Message: err.Error(),
		}
		return range_operations.NewGetRangesInternalServerError().WithPayload(&resp)
	}

	competitions := []*models.Competition{}
	for _, dbc := range dbCompetitions {
		competitions = append(competitions, dbCompetitionToCompetition(&dbc))
	}

	return competition.NewGetCompetitionsOK().WithPayload(competitions)
}

func getCompetitionsHTML(api *operations.LosAPI, params competition.GetCompetitionsHTMLParams) middleware.Responder {
	dbCompetitions, err := db.GetCompetitions(params.RangeID, *params.ActiveOnly)
	if err != nil {
		resp := models.APIResponse{
			Message: err.Error(),
		}
		return range_operations.NewGetRangesInternalServerError().WithPayload(&resp)
	}

	competitions := []*models.Competition{}
	for _, dbc := range dbCompetitions {
		competitions = append(competitions, dbCompetitionToCompetition(&dbc))
	}

	htmlRes := models.HTMLResponse{
		Template: "templates/competitions_html.gotmpl",
		Payload:  competitions,
	}
	return competition.NewGetCompetitionsHTMLOK().WithPayload(&htmlRes)
}

func dbRangeToRange(dbr *database.Range) *models.Range {
	return &models.Range{
		ID:        dbr.ID,
		Name:      dbr.Name,
		Latitude:  dbr.Latitude,
		Longitude: dbr.Longitude,
		URL:       dbr.URL,
		Active:    dbr.Active,
	}
}

func dbCompetitionToCompetition(dbc *database.Competition) *models.Competition {
	return &models.Competition{
		ID:        dbc.ID,
		Name:      dbc.Name,
		EventDate: strfmt.Date(dbc.EventDate),
		//RangeID:   dbc.RangeID,
		Category: &models.CompetitionCategory{
			ID:   dbc.CategoryID,
			Code: dbc.CategoryCode,
			Name: dbc.CategoryName,
		},
		Type: &models.CompetitionType{
			ID:   dbc.TypeID,
			Code: dbc.TypeCode,
			Name: dbc.TypeName,
		},
		Range: &models.Range{
			ID:   dbc.RangeID,
			Name: dbc.RangeName,
		},
		//Active: dbc.Active,
	}
}

// this is middleware to serve swagger-ui UI
func staticContentMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Shortcut helpers for swagger-ui
		if r.URL.Path == "/swagger-ui" || r.URL.Path == "/swaggerui" || r.URL.Path == "/api/help" {
			http.Redirect(w, r, "/swagger-ui/", http.StatusFound)
			return
		}
		// Serving ./swagger-ui/
		if strings.Index(r.URL.Path, "/swagger-ui/") == 0 {
			fs := http.FileServer(http.Dir("swagger-ui"))
			http.StripPrefix("/swagger-ui/", fs).ServeHTTP(w, r)
			return
		} else if strings.Index(r.URL.Path, "/static/") == 0 {
			fs := http.FileServer(http.Dir("static"))
			http.StripPrefix("/static/", fs).ServeHTTP(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func logrusMiddleware(handler http.Handler) http.Handler {
	if logrusHandler == nil {
		logrusHandler = interpose.NegroniLogrus()
	}
	return logrusHandler(handler)
}
