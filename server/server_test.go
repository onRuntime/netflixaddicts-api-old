package server

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func Test_GetSheets(t *testing.T) {
	url := "http://localhost:8080/api/sheets"
	req, _ := http.NewRequest("GET", url, nil)
	log.Print(doRequest(req))
}

func Test_PostSheet(t *testing.T) {
	url := "http://localhost:8080/api/sheet"
	payload := []byte(`{"name": "stranger-things", "title": "Stranger Things", "image": "", "note": 0, "styles": [], "synopsis": "Meilleure série srx"}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	log.Print(doRequest(req))
}

func Test_GetSheet(t *testing.T) {
	url := "http://localhost:8080/api/sheet/stranger-things"
	req, _ := http.NewRequest("GET", url, nil)
	log.Print(doRequest(req))
}

func Test_DeleteSheet(t *testing.T) {
	url := "http://localhost:8080/api/sheet/stranger-things"
	req, _ := http.NewRequest("DELETE", url, nil)
	log.Print(doRequest(req))
}

func doRequest(r *http.Request) string {
	r.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
