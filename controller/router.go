// configure_sample.goをカスタマイズしたもの。configure_sampleを直接

package controller

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/budougumi0617/sandbox_goswagger/gen/restapi/operations"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

func ConfigureAPI(api *operations.SampleAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	// TODO: この辺のイジり方はまだわかっていない
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	// 独自実装したハンドラを追加する
	api.PostAPIRegisterHandler = operations.PostAPIRegisterHandlerFunc(PostAPIRegisterHandler)
	api.GetGreetingHandler = operations.GetGreetingHandlerFunc(GetGreetingHandler)

	// 独自実装したMiddlewareを個別にセットする。
	// pathはinitHandlerCache関数内で自動生成されているので、定数などはない。
	mw := middleware.Builder(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("in specifiedMiddleware: %v", r.URL.Path)
			handler.ServeHTTP(w, r)
			log.Printf("in specifiedMiddleware after: %v", r.URL.Path)
		})
	})
	api.AddMiddlewareFor("POST", "/api/register", mw)

	api.PreServerShutdown = func() {}
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("in setupMiddleware: %v", r.URL.Path)
		handler.ServeHTTP(w, r)
		log.Printf("in setupMiddleware after: %v", r.URL.Path)
	})
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("in setupGlobalMiddleware: %v", r.URL.Path)
		handler.ServeHTTP(w, r)
		log.Printf("in setupGlobalMiddleware after: %v", r.URL.Path)
	})
}
