package main

import(
	"log"
	"net/http"
	"encoding/json"
	//"strconv"

	"./module"
	"github.com/ant0ine/go-json-rest/rest"
)

type postPageInput struct {
	Page json.Number
}

type postPageJSON struct {
	Result string
}

func postPage(Rw rest.ResponseWriter, req *rest.Request) {
	input := postPageInput{}
	err := req.DecodeJsonPayload(&input)
	if err != nil {
		rest.Error(Rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("input: %#v", input)

	PNumber, err := input.Page.Int64()
	if err != nil {
		rest.Error(Rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if PNumber != 0 {
		Rw.WriteJson(&postPageJSON{
			"Page number is OK",
		})
		moduleDB.DB()
	} else {
		rest.Error(Rw, "Page number is required", 400)
	}
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/page", postPage),
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server started.")

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":5200", api.MakeHandler()))
}

