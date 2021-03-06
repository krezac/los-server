// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"html/template"
	"io"
	"log"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	swag "github.com/go-openapi/swag"
	"github.com/krezac/los-server/auth"
	"github.com/krezac/los-server/database"
	"github.com/krezac/los-server/models"
	"github.com/krezac/los-server/restapi/operations"
	"github.com/krezac/los-server/restapi/operations/competition"
	"github.com/krezac/los-server/restapi/operations/login"
	"github.com/krezac/los-server/restapi/operations/range_operations"
	"github.com/krezac/los-server/restapi/operations/user"
)

//go:generate swagger generate server --target .. --name Los --spec ../swagger/los-server.yml --principal models.Principal

var db *database.Database

func configureFlags(api *operations.LosAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		swag.CommandLineOptionsGroup{
			ShortDescription: "JWT Options",
			Options:          auth.JwtExtraOptionsVar,
		},
	}
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
	api.CompetitionCreateCompetitionHandler = competition.CreateCompetitionHandlerFunc(func(params competition.CreateCompetitionParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation competition.CreateCompetition has not yet been implemented")
	})
	api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
	})
	api.CompetitionDeleteCompetitionHandler = competition.DeleteCompetitionHandlerFunc(func(params competition.DeleteCompetitionParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation competition.DeleteCompetition has not yet been implemented")
	})
	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
	})
	api.CompetitionGetCompetitionByIDHandler = competition.GetCompetitionByIDHandlerFunc(func(params competition.GetCompetitionByIDParams) middleware.Responder {
		return getCompetitionByID(api, params) // IMPL change against generated file
	})
	api.CompetitionGetCompetitionsHandler = competition.GetCompetitionsHandlerFunc(func(params competition.GetCompetitionsParams) middleware.Responder {
		return getCompetitions(api, params) // IMPL change against generated file
	})
	api.CompetitionGetCompetitionsHTMLHandler = competition.GetCompetitionsHTMLHandlerFunc(func(params competition.GetCompetitionsHTMLParams) middleware.Responder {
		return getCompetitionsHTML(api, params) // IMPL change against generated file
	})
	api.RangeOperationsGetRangeByIDHandler = range_operations.GetRangeByIDHandlerFunc(func(params range_operations.GetRangeByIDParams) middleware.Responder {
		return getRangeByID(api, params) // IMPL change against generated file
	})
	api.RangeOperationsGetRangesHandler = range_operations.GetRangesHandlerFunc(func(params range_operations.GetRangesParams) middleware.Responder {
		return getRanges(api, params) // IMPL change against generated file
	})
	api.RangeOperationsGetRangesHTMLHandler = range_operations.GetRangesHTMLHandlerFunc(func(params range_operations.GetRangesHTMLParams) middleware.Responder {
		return getRangesHTML(api, params) // IMPL change against generated file
	})
	api.UserGetUserByNameHandler = user.GetUserByNameHandlerFunc(func(params user.GetUserByNameParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation user.GetUserByName has not yet been implemented")
	})
	api.LoginLoginUserHandler = login.LoginUserHandlerFunc(func(params login.LoginUserParams) middleware.Responder {
		return loginUser(api, params) // IMPL change against generated file
	})
	api.LoginLogoutUserHandler = login.LogoutUserHandlerFunc(func(params login.LogoutUserParams, principal *models.Principal) middleware.Responder {
		return logoutUser(api, params, principal) // IMPL change against generated file
	})
	api.LoginRefreshTokenHandler = login.RefreshTokenHandlerFunc(func(params login.RefreshTokenParams, principal *models.Principal) middleware.Responder {
		return refreshToken(api, params, principal) // IMPL change against generated file
	})
	api.CompetitionUpdateCompetitonHandler = competition.UpdateCompetitonHandlerFunc(func(params competition.UpdateCompetitonParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation competition.UpdateCompetiton has not yet been implemented")
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
