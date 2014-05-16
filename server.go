package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/cors"
	"net/http"
)

type Server struct {
	Port int64

	*martini.ClassicMartini
}

func (c *Server) mapRoutes() {
	c.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{".*"},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	c.Get("/", func() string {
		return "You want weapons? We’re in a library! Books! The best weapons in the world!"
	})

	c.Post("/collect", func(req *http.Request) string {
		key := req.FormValue("metric[key]")
		value := req.FormValue("metric[value]")
		kind := req.FormValue("metric[kind]")

		go func(key string, value string, kind string) {
			metric, err := NewMetric(kind, key, value)
			if err == nil {
				err = metric.Track()
			}
		}(key, value, kind)

		return "1"
	})
}

func (c *Server) Start() {
	fmt.Println("Running the server on port", c.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), c)
}

func NewServer(port int64) *Server {
	martini := martini.Classic()
	server := &Server{port, martini}
	server.mapRoutes()
	return server
}
