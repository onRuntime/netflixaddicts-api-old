package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/netflixaddicts/Go-API/data"
	"github.com/netflixaddicts/Go-API/structs"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	router *mux.Router
	Data   *data.Data
}

func New() *Server {
	return &Server{}
}

func (s *Server) Initialize() {
	log.Print("\n\n -==  NetflixAddicts Rest-API.go ==- \n\n")

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	s.Data = data.New()
	s.Data.Connect(os.Getenv("DB_ADDR"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	defer s.Data.Close()

	s.router = mux.NewRouter().StrictSlash(false)
	api := s.router.PathPrefix("/api").Subrouter()
	s.handleRoutes(api)
}

func (s *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.router))
}

func (s *Server) handleRoutes(api *mux.Router) {
	api.HandleFunc("/sheets", s.handleRequest(GetSheets)).Methods("GET")

	api.HandleFunc("/sheet", s.handleRequest(PostSheet)).Methods("POST")
	api.HandleFunc("/sheet/{name}", s.handleRequest(GetSheet)).Methods("GET")
	api.HandleFunc("/sheet/{name}", s.handleRequest(PatchSheet)).Methods("PATCH")
	api.HandleFunc("/sheet/{name}", s.handleRequest(DeleteSheet)).Methods("DELETE")
}

func GetSheets(d *data.Data, w http.ResponseWriter, r *http.Request) []byte {
	return sendJson(w, http.StatusOK, d.Sheets)
}

func PostSheet(d *data.Data, w http.ResponseWriter, r *http.Request) []byte {
	sheet := &structs.Sheet{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(sheet); err != nil {
		return sendError(w, http.StatusBadRequest, err.Error())
	}
	defer r.Body.Close()

	if d.Sheets[sheet.Name] != nil {
		return sendError(w, http.StatusBadRequest, fmt.Sprintf("A sheet named '%s' already exists in database.", sheet.Name))
	}
	sheet.ID = len(d.Sheets) + 1
	sheet.CreatedAt = time.Now()
	sheet.UpdatedAt = time.Now()

	d.Sheets[sheet.Name] = sheet
	// TODO: Make DB insert query

	return sendJson(w, http.StatusCreated, nil)
}

func GetSheet(d *data.Data, w http.ResponseWriter, r *http.Request) []byte {
	vars := mux.Vars(r)
	name := vars["name"]

	sheet, ok := d.Sheets[name]
	if !ok {
		return sendError(w, http.StatusNotFound, fmt.Sprintf("No sheet were found with name '%s'", name))
	}
	return sendJson(w, http.StatusOK, sheet)
}

func PatchSheet(d *data.Data, w http.ResponseWriter, r *http.Request) []byte {
	return sendJson(w, http.StatusOK, nil)
}

func DeleteSheet(d *data.Data, w http.ResponseWriter, r *http.Request) []byte {
	vars := mux.Vars(r)
	name := vars["name"]

	_, ok := d.Sheets[name]
	if !ok {
		return sendError(w, http.StatusNotFound, fmt.Sprintf("No sheet were found with name '%s'", name))
	}
	delete(d.Sheets, name)
	// TODO: Make DB remove query

	return sendJson(w, http.StatusOK, nil)
}

type RequestHandler func(d *data.Data, w http.ResponseWriter, r *http.Request) []byte

func (s *Server) handleRequest(handler RequestHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s - %s - %s] Request received...", r.Method, r.RequestURI, r.RemoteAddr)
		res := handler(s.Data, w, r)
		log.Printf("[%s - %s - %s] Request handled, responded with '%s'", r.Method, r.RequestURI, r.RemoteAddr, res)
	}
}

func sendError(w http.ResponseWriter, code int, message string) []byte {
	return sendJson(w, code, map[string]string{"error": message})
}

func sendJson(w http.ResponseWriter, code int, payload interface{}) []byte {
	response, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if payload != nil {
		_, _ = w.Write(response)
	}
	return response
}
