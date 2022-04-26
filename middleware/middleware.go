/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-04-26 10:27:36
 * @LastEditTime: 2022-04-26 14:29:03
 * @FilePath: /graphs-Rishabh-Mishra/middleware/middleware.go
 * @description: type some description
 */

package middleware

import (
	"log"
	"net/http"
	"time"
)

// our base middleware implementation.
type service func(http.Handler) http.Handler

//ServiceLoader chain load middleware services.
func ServiceLoader(h http.Handler, svcs ...service) http.Handler {
	for _, svc := range svcs {
		h = svc(h)
	}
	return h
}

//RequestMetrics middleware for request length metrics.
func RequestMetrics(l *log.Logger) service {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			h.ServeHTTP(w, r)
			l.Printf("%s request to %s took %vns.", r.Method, r.URL.Path, time.Now().Sub(start).Nanoseconds())
		})
	}
}
