package main

import(
	"log"
	"net/http"
	"strconv"

	"./module"
	"github.com/ant0ine/go-json-rest/rest"
)

// type Result_JSON struct {
// 	Result string
// }

// https://host-name:port/api/v1/twimg/data/page/{PageNum}
func API_twimg(Rw rest.ResponseWriter, req *rest.Request) {
	page := req.PathParam("PageNum")
	log.Println("Page Number is ", page)

	PNum, err := strconv.Atoi(page)
	if err != nil {
		rest.Error(Rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if PNum != 0 {
		json := "Page number is "+page

		useDB.DB_home()
		Send_JSON(Rw, json)
	} else {
		rest.Error(Rw, "Page number is required", 400)
	}
}

// https://host-name:port/api/v1/twimg/data/search/{TwiID}/{PageNum}
func API_twimg_search(Rw rest.ResponseWriter, req *rest.Request) {
	twiID := req.PathParam("TwiID")
	page := req.PathParam("PageNum")
	log.Println("TwitterID is ", twiID)
	log.Println("Page Number is ", page)

	PNum, err := strconv.Atoi(page)
	if err != nil {
		rest.Error(Rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if PNum != 0 && twiID != "" {
		json := "Page number is "+page+" TwitterID is "+twiID

		useDB.DB_search()
		Send_JSON(Rw, json)
	} else {
		rest.Error(Rw, "Page number & TwitterID is required", 400)
	}
}

// https://host-name:port/api/v1/twimg/data/original/{TwiID}/{FileName}
func API_twimg_original(Rw rest.ResponseWriter, req *rest.Request) {
	twiID := req.PathParam("TwiID")
	imgID := req.PathParam("ImageID")
	log.Println("TwitterID is ", twiID)
	log.Println("ImageID is ", imgID)

	if twiID != "" && imgID != "" {
		json := "TwitterID is "+twiID+" UserID is "+imgID

		useDB.DB_origin()
		Send_JSON(Rw, json)
	} else {
		rest.Error(Rw, "ImageID & TwitterID is required", 400)
	}
}

func Send_JSON(Rw rest.ResponseWriter, j string) {
	Rw.WriteJson(j)
	//Rw.WriteJson(&Result_JSON{json})
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/api/v1/twimg/data/page/:PageNum", API_twimg),
		rest.Get("/api/v1/twimg/data/search/:TwiID/:PageNum", API_twimg_search),
		rest.Get("/api/v1/twimg/data/original/:TwiID/:ImageID", API_twimg_original),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started.")

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":5200", api.MakeHandler()))
}

