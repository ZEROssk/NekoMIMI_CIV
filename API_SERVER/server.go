package main

import(
	"log"
	"net/http"
	"strconv"

	"./module"
	"github.com/ant0ine/go-json-rest/rest"
)

type ResultJSON struct {
	Result string
}

// https://host-name:port/api/v1/twimg/data/page?p={PageNum}
func API_twimg(Rw rest.ResponseWriter, req *rest.Request) {
	v := req.URL.Query()
	page := v.Get("p")

	m := 5

	PNum, err := strconv.Atoi(page)
	if err != nil {
		rest.Error(Rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if PNum != 0 {
		json := "Page number is "+page
		e := m*PNum
		b := e-m
		es := strconv.Itoa(e)
		bs := strconv.Itoa(b)

		useDB.DB_home(page, bs, es)
		SendJSON(Rw, json)
	} else {
		rest.Error(Rw, "Page number is required", 400)
	}
}

// https://host-name:port/api/v1/twimg/data/search?tid={TwiID}&p={PageNum}
func API_twimg_search(Rw rest.ResponseWriter, req *rest.Request) {
	v := req.URL.Query()
	twiID := v.Get("tid")
	page := v.Get("p")

	PNum, err := strconv.Atoi(page)
	if err != nil {
		rest.Error(Rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if PNum != 0 && twiID != "" {
		json := "Page number is "+page+" TwitterID is "+twiID

		useDB.DB_search(twiID, page)
		SendJSON(Rw, json)
	} else {
		rest.Error(Rw, "Page number & TwitterID is required", 400)
	}
}

// https://host-name:port/api/v1/twimg/data/original?tid={TwiID}&fname={FileName}
func API_twimg_original(Rw rest.ResponseWriter, req *rest.Request) {
	v := req.URL.Query()
	twiID := v.Get("tid")
	img := v.Get("fname")

	if twiID != "" && img != "" {
		json := "TwitterID is "+twiID+" UserID is "+img

		useDB.DB_origin(twiID, img)
		SendJSON(Rw, json)
	} else {
		rest.Error(Rw, "FileName & TwitterID is required", 400)
	}
}

func SendJSON(Rw rest.ResponseWriter, j string) {
	Rw.WriteJson(&ResultJSON{j})
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/api/v1/twimg/data/page", API_twimg),
		rest.Get("/api/v1/twimg/data/search", API_twimg_search),
		rest.Get("/api/v1/twimg/data/original", API_twimg_original),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started.")

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":5200", api.MakeHandler()))
}

