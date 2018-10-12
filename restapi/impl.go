package restapi

import (
	"net/http"
	"strings"

	interpose "github.com/carbocation/interpose/middleware"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/krezac/los-server/auth"
	"github.com/krezac/los-server/models"
	"github.com/krezac/los-server/restapi/operations"
	"github.com/krezac/los-server/restapi/operations/range_operations"
	"github.com/rakyll/statik/fs"
)

var logrusHandler func(http.Handler) http.Handler

func losAuthImpl(api *operations.LosAPI, token string, scopes []string) (*models.Principal, error) {
	api.Logger("HasRoleAuth handler called")
	return auth.HasRole(token, scopes)
	/*
		p := models.Principal{
			Name:  "gogo",
			Roles: []string{"admin"},
		}
		return &p, nil
	*/
}

func getRangeByID(params range_operations.GetRangeByIDParams, principal *models.Principal) middleware.Responder {
	dbr, err := db.GetRangeByID(params.RangeID)
	if err != nil {
		resp := models.APIResponse{
			Message: err.Error(),
		}
		return range_operations.NewGetRangeByIDNotFound().WithPayload(&resp)
	}

	r := models.Range{
		ID:        dbr.ID,
		Name:      dbr.Name,
		Latitude:  dbr.Latitude,
		Longitude: dbr.Longitude,
		Active:    dbr.Active,
	}

	return range_operations.NewGetRangeByIDOK().WithPayload(&r)
}

func getRanges(params range_operations.GetRangesParams, principal *models.Principal) middleware.Responder {
	dbRanges, err := db.GetRanges()
	if err != nil {
		resp := models.APIResponse{
			Message: err.Error(),
		}
		return range_operations.NewGetRangesInternalServerError().WithPayload(&resp)
	}

	ranges := []*models.Range{}
	for _, dbr := range dbRanges {
		r := models.Range{
			ID:        dbr.ID,
			Name:      dbr.Name,
			Latitude:  dbr.Latitude,
			Longitude: dbr.Longitude,
			Active:    dbr.Active,
		}
		ranges = append(ranges, &r)
	}

	return range_operations.NewGetRangesOK().WithPayload(ranges)
}

// this is middleware to serve swagger-ui UI
func swaggerUiMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Shortcut helpers for swagger-ui
		if r.URL.Path == "/swagger-ui" || r.URL.Path == "/swaggerui" || r.URL.Path == "/api/help" {
			http.Redirect(w, r, "/swagger-ui/", http.StatusFound)
			return
		}
		// Serving ./swagger-ui/
		if strings.Index(r.URL.Path, "/swagger-ui/") == 0 {
			statikFS, err := fs.New()
			if err != nil {
				panic(err)
			}
			staticServer := http.FileServer(statikFS)
			http.StripPrefix("/swagger-ui/", staticServer).ServeHTTP(w, r)

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