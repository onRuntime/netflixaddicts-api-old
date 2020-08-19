package router

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/netflixaddicts/Go-API/data"
	"github.com/netflixaddicts/Go-API/structs"
	"net/http"
	"time"
)

func GetSheets(d *data.Data, w http.ResponseWriter, r *http.Request) {
	sendJson(w, http.StatusOK, d.Sheets)
}

func PostSheet(d *data.Data, w http.ResponseWriter, r *http.Request) {
	sheet := &structs.Sheet{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(sheet); err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if d.Sheets[sheet.Name] != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("A sheet named '%s' already exists in database.", sheet.Name))
		return
	}
	sheet.ID = len(d.Sheets) + 1
	sheet.CreatedAt = time.Now()
	sheet.UpdatedAt = time.Now()

	d.Sheets[sheet.Name] = sheet
	// TODO: Make DB insert query

	sendJson(w, http.StatusCreated, nil)
}

func GetSheet(d *data.Data, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	sheet, ok := d.Sheets[name]
	if !ok {
		sendError(w, http.StatusNotFound, fmt.Sprintf("No sheet were found with name '%s'", name))
		return
	}
	sendJson(w, http.StatusOK, sheet)
}

func PatchSheet(d *data.Data, w http.ResponseWriter, r *http.Request) {}

func DeleteSheet(d *data.Data, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	_, ok := d.Sheets[name]
	if !ok {
		sendError(w, http.StatusNotFound, fmt.Sprintf("No sheet were found with name '%s'", name))
		return
	}
	delete(d.Sheets, name)
	// TODO: Make DB remove query

	sendJson(w, http.StatusOK, nil)
}

func sendError(w http.ResponseWriter, code int, message string) {
	sendJson(w, code, map[string]string{"error": message, "code": string(rune(code))})
}

func sendJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if payload != nil {
		_, _ = w.Write(response)
	}
}
