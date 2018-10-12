// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"strings"

	"github.com/krezac/los-server/database"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/rakyll/statik/fs"

	interpose "github.com/carbocation/interpose/middleware"
	"github.com/krezac/los-server/models"
	"github.com/krezac/los-server/restapi/operations"
	"github.com/krezac/los-server/restapi/operations/range_operations"
	"github.com/krezac/los-server/restapi/operations/user"
	_ "github.com/krezac/los-server/swaggeruistatik" // this is to make sure the initializer gets called
)

//go:generate swagger generate server --target .. --name Los --spec ../swagger/los-server.yml

var db *database.Database

func configureFlags(api *operations.LosAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.LosAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	// database access
	var err error
	db, err = database.NewMysqlDatabase()
	if err != nil {
		panic(err) // TODO do better handling
	}

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "api_key" header is set
	/* TODO removed as not in use?
	api.APIKeyAuth = func(token string) (interface{}, error) {
		// TODO return nil, errors.NotImplemented("api key auth (api_key) api_key from header param [api_key] has not yet been implemented")
		return "have it", nil
	}
	*/

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
		return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
	})
	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams) middleware.Responder {
		return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
	})
	api.RangeOperationsGetRangeByIDHandler = range_operations.GetRangeByIDHandlerFunc(func(params range_operations.GetRangeByIDParams) middleware.Responder {
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
	})

	api.RangeOperationsGetRangesHandler = range_operations.GetRangesHandlerFunc(func(params range_operations.GetRangesParams) middleware.Responder {
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
	})
	api.UserGetUserByNameHandler = user.GetUserByNameHandlerFunc(func(params user.GetUserByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation user.GetUserByName has not yet been implemented")
	})
	api.UserLoginUserHandler = user.LoginUserHandlerFunc(func(params user.LoginUserParams) middleware.Responder {
		return middleware.NotImplemented("operation user.LoginUser has not yet been implemented")
	})
	api.UserLogoutUserHandler = user.LogoutUserHandlerFunc(func(params user.LogoutUserParams) middleware.Responder {
		return middleware.NotImplemented("operation user.LogoutUser has not yet been implemented")
	})
	api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(func(params user.UpdateUserParams) middleware.Responder {
		return middleware.NotImplemented("operation user.UpdateUser has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {

	logViaLogrus := interpose.NegroniLogrus()
	return logViaLogrus(uiMiddleware(handler))
}

// this is middleware to serve swagger-ui UI
func uiMiddleware(handler http.Handler) http.Handler {
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
