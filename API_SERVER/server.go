package main

import(
	."fmt"
	"log"
	"net/http"
	"strconv"

	"./module"
	"github.com/ant0ine/go-json-rest/rest"
)

type ResultJSON struct {
	Result string
}

type testJSON struct {
	Result []string
}

var get_by_n int = 5

// https://host-name:port/api/v1/twimg/page?p={PageNum}
func API_twimg(Rw rest.ResponseWriter, req *rest.Request) {
	v := req.URL.Query()
	PNum, err := strconv.Atoi(v.Get("p"))
	if err != nil {
		rest.Error(Rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if PNum != 0 {
		//json := Sprintf("Page number is %d", PNum)
		start := (get_by_n*PNum)-get_by_n

		content := useDB.DB_home(
			Sprintf("%d", PNum),
			Sprintf("%d", start),
			Sprintf("%d", get_by_n),
		)
		//SendJSON(Rw, json + content)
		Rw.WriteJson(&testJSON{content})
	} else {
		rest.Error(Rw, "Page number is required", 400)
	}
}

// https://host-name:port/api/v1/twimg/search?tid={TwiID}&p={PageNum}
func API_twimg_search(Rw rest.ResponseWriter, req *rest.Request) {
	v := req.URL.Query()
	twiID := v.Get("tid")

	PNum, err := strconv.Atoi(v.Get("p"))
	if err != nil {
		rest.Error(Rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if PNum != 0 && twiID != "" {
		json := Sprintf("Page number is %d TwiID is %s", PNum, twiID)
		start := (get_by_n*PNum)-get_by_n

		content := useDB.DB_search(
			twiID,
			Sprintf("%d", start),
			Sprintf("%d", get_by_n),
		)
		SendJSON(Rw, json + content)
	} else {
		rest.Error(Rw, "Page number & TwitterID is required", 400)
	}
}

// https://host-name:port/api/v1/twimg/original?tid={TwiID}&fname={FileName}
func API_twimg_original(Rw rest.ResponseWriter, req *rest.Request) {
	v := req.URL.Query()
	twiID := v.Get("tid")
	img := v.Get("fname")

	if twiID != "" && img != "" {
		json := Sprintf("TwitterID is %s FileName is %s", twiID, img)

		content := useDB.DB_origin(twiID, img)
		SendJSON(Rw, json + content)
	} else {
		rest.Error(Rw, "FileName & TwitterID is required", 400)
	}
}

func SendJSON(Rw rest.ResponseWriter, j string) {
	Rw.WriteJson(&ResultJSON{j})
}

func main() {
	useDB.Login_DB()
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/api/v1/twimg/page", API_twimg),
		rest.Get("/api/v1/twimg/search", API_twimg_search),
		rest.Get("/api/v1/twimg/original", API_twimg_original),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started.")

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":5200", api.MakeHandler()))
}

