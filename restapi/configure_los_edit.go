// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/krezac/los-server/database"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/krezac/los-server/models"
	"github.com/krezac/los-server/restapi/operations"
	"github.com/krezac/los-server/restapi/operations/competition"
	"github.com/krezac/los-server/restapi/operations/range_operations"
	"github.com/krezac/los-server/restapi/operations/user"
)

//go:generate swagger generate server --target .. --name Los --spec ../swagger/los-server.yml --principal models.Principal

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
	api.Logger = log.Printf // IMPL change against generated file

	// database access
	var err error
	db, err = database.NewMysqlDatabase()
	if err != nil {
		panic(err) // TODO do better handling
	}

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.HTMLProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		htmlRes := data.(*models.HTMLResponse)
		tmpl, err := template.ParseFiles(htmlRes.Template)
		if err != nil {
			return err
		}
		return tmpl.Execute(w, htmlRes.Payload)
	})

	api.LosAuthAuth = func(token string, scopes []string) (*models.Principal, error) {
		return losAuthImpl(api, token, scopes) // IMPL change against generated file
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
	})
	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
	})
	api.CompetitionGetCompetitionByIDHandler = competition.GetCompetitionByIDHandlerFunc(func(params competition.GetCompetitionByIDParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation competition.GetCompetitionByID has not yet been implemented")
	})
	api.CompetitionGetCompetitionsHandler = competition.GetCompetitionsHandlerFunc(func(params competition.GetCompetitionsParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation competition.GetCompetitions has not yet been implemented")
	})
	api.CompetitionGetCompetitionsHTMLHandler = competition.GetCompetitionsHTMLHandlerFunc(func(params competition.GetCompetitionsHTMLParams) middleware.Responder {
		return middleware.NotImplemented("operation competition.GetCompetitionsHTML has not yet been implemented")
	})
	api.RangeOperationsGetRangeByIDHandler = range_operations.GetRangeByIDHandlerFunc(func(params range_operations.GetRangeByIDParams, principal *models.Principal) middleware.Responder {
		return getRangeByID(api, params, principal) // IMPL change against generated file
	})

	api.RangeOperationsGetRangesHandler = range_operations.GetRangesHandlerFunc(func(params range_operations.GetRangesParams, principal *models.Principal) middleware.Responder {
		return getRanges(api, params, principal) // IMPL change against generated file
	})
	api.RangeOperationsGetRangesHTMLHandler = range_operations.GetRangesHTMLHandlerFunc(func(params range_operations.GetRangesHTMLParams) middleware.Responder {
		return getRangesHTML(api, params) // IMPL change against generated file
	})
	api.UserGetUserByNameHandler = user.GetUserByNameHandlerFunc(func(params user.GetUserByNameParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation user.GetUserByName has not yet been implemented")
	})
	api.UserLoginUserHandler = user.LoginUserHandlerFunc(func(params user.LoginUserParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation user.LoginUser has not yet been implemented")
	})
	api.UserLogoutUserHandler = user.LogoutUserHandlerFunc(func(params user.LogoutUserParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation user.LogoutUser has not yet been implemented")
	})
	api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(func(params user.UpdateUserParams, principal *models.Principal) middleware.Responder {
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

	return logrusMiddleware(staticContentMiddleware(handler)) // IMPL change against generated file
}
