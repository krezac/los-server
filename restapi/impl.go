package restapi

import (
	"net/http"
	"strings"

	interpose "github.com/carbocation/interpose/middleware"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/krezac/los-server/auth"
	"github.com/krezac/los-server/database"
	"github.com/krezac/los-server/models"
	"github.com/krezac/los-server/restapi/operations"
	"github.com/krezac/los-server/restapi/operations/range_operations"
)

var logrusHandler func(http.Handler) http.Handler

func losAuthImpl(api *operations.LosAPI, token string, scopes []string) (*models.Principal, error) {
	api.Logger("HasRoleAuth handler called")
	return auth.HasRole(token, scopes)
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
	dbRanges, err := db.GetRanges()
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
	dbRanges, err := db.GetRanges()
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
