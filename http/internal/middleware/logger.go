package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger returns a handler function that will log inf about
// the request then call the provided handler function.
func Logger(next http.HandlerFunc) http.HandlerFunc {

	// Wrap this hadnler around next one prvided.
	return func(res http.ResponseWriter, req *http.Request) {

		start := time.Now()

		// Once tge handler call proceeding this defer
		// is complete, log how long the request took.
		defer func() {
			d := time.Now().Sub(start)
			log.Printf("(%s) : %s -> %s (%s)", req.Method, req.URL.Path, req.RemoteAddr, d)
		}()

		next(res, req)
	}

}
