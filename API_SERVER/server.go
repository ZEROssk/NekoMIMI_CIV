package main

import(
	"log"
	"net/http"

	"./api"
	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("api/v1/image_viewer/images/twimg/data?p={PageNum}", postPage),
		rest.Post("api/v1/image_viewer/images/twimg/data/search?id={UserID}&p={PageNum}", postTwiID_Page),
		rest.Post("api/v1/image_viewer/images/twimg/data/original?id={UserID}&img={ImageID}", postTwiID_Page),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started.")

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":5200", api.MakeHandler()))
}

