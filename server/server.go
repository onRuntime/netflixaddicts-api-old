package server

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/netflixaddicts/Go-API/data"
	"github.com/netflixaddicts/Go-API/router"
	"log"
	"net/http"
	"os"
)

type Server struct {
	router *mux.Router
	Data   *data.Data
}

func New() *Server {
	return &Server{}
}

func (s *Server) Initialize() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	s.Data = data.New()
	s.Data.Connect(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_ADDR"))
	defer s.Data.Close()

	s.router = mux.NewRouter().StrictSlash(false)
	api := s.router.PathPrefix("/api").Subrouter()
	s.handleRoutes(api)
}

func (s *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.router))
}

func (s *Server) handleRoutes(api *mux.Router) {
	api.HandleFunc("/sheets", s.handleRequest(router.GetSheets)).Methods("GET")

	api.HandleFunc("/sheet", s.handleRequest(router.PostSheet)).Methods("POST")
	api.HandleFunc("/sheet/{name}", s.handleRequest(router.GetSheet)).Methods("GET")
	api.HandleFunc("/sheet/{name}", s.handleRequest(router.PatchSheet)).Methods("PATCH")
	api.HandleFunc("/sheet/{name}", s.handleRequest(router.DeleteSheet)).Methods("DELETE")
}

type RequestHandler func(d *data.Data, w http.ResponseWriter, r *http.Request)

func (s *Server) handleRequest(handler RequestHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(s.Data, w, r)
	}
}
