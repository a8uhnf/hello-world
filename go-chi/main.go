package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a8uhnf/hello-world/go-chi/pkg"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var promMetrics *prometheus.HistogramVec

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
}

func middlewareOne(fn http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		fn.ServeHTTP(w, r)
		t2 := time.Now()
		mp := make(map[string]string)
		mp["method"] = r.Method
		mp["endpoint"] = r.URL.Path
		// mp["duration"] = strconv.FormatInt(t2.Sub(t1).Nanoseconds(), 10)
		fmt.Println(mp)
		_, err := promMetrics.CurryWith(mp)
		if err != nil {
			fmt.Println("------------")
			log.Println(err.Error())
		}
		ctx := r.Context()
		chiCtx := chi.RouteContext(ctx)
		fmt.Println(chi.URLParam(r, "id"))
		fmt.Println("request URI", r.RequestURI)
		fmt.Println(chiCtx.RoutePattern())

		promMetrics.WithLabelValues(r.Method, chiCtx.RoutePattern()).Observe(float64(t2.Sub(t1).Nanoseconds()))
		log.Println("xxxxx", r.Method, t2.Sub(t1).Nanoseconds())
	})
}

func main() {
	r := chi.NewRouter()
	promMetrics = pkg.RegisterMetrics("test_metrics")
	prometheus.MustRegister(promMetrics)

	r.Use(middlewareOne)
	r.Use(middleware.Logger)
	r.Method("GET", "/", Handler(customHandler))
	r.Method("GET", "/metrics", promhttp.Handler())
	r.Group(func(rr chi.Router) {
		rr.Method("GET", "/hello-1", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte("hello-1"))
		}))
		rr.Method("GET", "/hello-2", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte("hello-2"))
		}))
		rr.Method("GET", "/hello/{id}", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte("hello id"))
		}))
	})
	http.ListenAndServe(":3333", r)
}

func customHandler(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")

	if q != "" {
		return errors.New(q)
	}
	w.Write([]byte("foo"))
	return nil
}
