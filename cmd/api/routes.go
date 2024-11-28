package main

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	// Use the requirePermission() middleware on each of the /v1/books** endpoints,
	// passing in the required permission code as the first parameter.
	router.HandlerFunc(http.MethodGet, "/v1/books", app.requirePermission("books:read", app.listBooksHandler))
	router.HandlerFunc(http.MethodPost, "/v1/books", app.requirePermission("books:write", app.createBookHandler))
	router.HandlerFunc(http.MethodGet, "/v1/books/:id", app.requirePermission("books:read", app.showBookHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/books/:id", app.requirePermission("books:write", app.updateBookHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/books/:id", app.requirePermission("books:write", app.deleteBookHandler))
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	// Register a new GET /debug/vars endpoint pointing to the expvar handler.
	router.Handler(http.MethodGet, "/debug/vars", expvar.Handler())

	return app.metrics(app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router)))))
}
