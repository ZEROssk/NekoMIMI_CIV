package main

import(
	"log"
	"net/http"
	"strconv"

	//"./module"
	"github.com/ant0ine/go-json-rest/rest"
)

type Result_JSON struct {
	Result string
}

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
		Rw.WriteJson(&Result_JSON{
			"Page number is "+page,
		})
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
		Rw.WriteJson(&Result_JSON{
			"Page number is "+page+" TwitterID is "+twiID,
		})
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
		Rw.WriteJson(&Result_JSON{
			"TwitterID is "+twiID+" UserID is "+imgID,
		})
	} else {
		rest.Error(Rw, "ImageID & TwitterID is required", 400)
	}
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/api/v1/twimg/data/page/:PageNum", API_twimg),
		rest.Post("/api/v1/twimg/data/search/:TwiID/:PageNum", API_twimg_search),
		rest.Post("/api/v1/twimg/data/original/:TwiID/:ImageID", API_twimg_original),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started.")

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":5200", api.MakeHandler()))
}

