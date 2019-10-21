package main

import(
	."fmt"
	"log"
	"net/http"
	"strconv"

	"./package_useDB"
	"./package_checkDB"
	"github.com/ant0ine/go-json-rest/rest"
)

type ResultJSONhome struct {
	PLimit	string		`json:"PageLimit"`
	PNum	string		`json:"PageNumber"`
	List	[]ImgJSON	`json:"Thumbnail"`
}

type ResultJSONsearch struct {
	TwiID	string		`json:"TwitterID"`
	PLimit	string		`json:"PageLimit"`
	PNum	string		`json:"PageNumber"`
	List	[]ImgJSON	`json:"Thumbnail"`
}

type ImgJSON struct {
	TwiID	string		`json:"TwitterID"`
	FName	string		`json:"FileName"`
}

var get_by_n int = 30

// https://host-name:port/api/v1/twimg/page?p={PageNum}
func API_twimg(Rw rest.ResponseWriter, req *rest.Request) {
	v := req.URL.Query()
	PNum, err := strconv.Atoi(v.Get("p"))
	if err != nil {
		rest.Error(Rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if PNum != 0 {
		start := (get_by_n*PNum)-get_by_n

		content, Pl := useDB.DB_home(
			Sprintf("%d", PNum),
			Sprintf("%d", start),
			Sprintf("%d", get_by_n),
		)

		var a int
		if (Pl % get_by_n) != 0 {
			a = (Pl / get_by_n) + 1
		} else {
			a = Pl / get_by_n
		}

		list := []ImgJSON{}
		for i := 0; i < len(content); i++ {
			list = append(list, ImgJSON{content[i][0], content[i][1]})
		}

		result := ResultJSONhome{}
		result.PLimit	= Sprintf("%d", a)
		result.PNum		= Sprintf("%d", PNum)
		result.List		= list

		Rw.WriteJson(&result)
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
		start := (get_by_n*PNum)-get_by_n

		content, Pl := useDB.DB_search(
			twiID,
			Sprintf("%d", start),
			Sprintf("%d", get_by_n),
		)

		var a int
		if (Pl % get_by_n) != 0 {
			a = (Pl / get_by_n) + 1
		} else {
			a = Pl / get_by_n
		}

		list := []ImgJSON{}
		for i := 0; i < len(content); i++ {
			list = append(list, ImgJSON{content[i][0], content[i][1]})
		}

		result := ResultJSONsearch{}
		result.TwiID	= twiID
		result.PLimit	= Sprintf("%d", a)
		result.PNum		= Sprintf("%d", PNum)
		result.List		= list

		Rw.WriteJson(&result)
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
		content := useDB.DB_origin(twiID, img)

		r := map[string]ImgJSON{}
		r["Image"] = ImgJSON{content[0], content[1]}

		Rw.WriteJson(r)
	} else {
		rest.Error(Rw, "FileName & TwitterID is required", 400)
	}
}

func main() {
	useDB.LoginDB()
	checkDB.CheckDB()

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

