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

type postResultJSON struct {
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
		Rw.WriteJson(&postResultJSON{
			"Page number is OK",
		})
		moduleDB.DB()
	} else {
		rest.Error(Rw, "Page number is required", 400)
	}
}

func postUserID(Rw rest.ResponseWriter, req *rest.Request) {
	var Tinput string
	err := req.DecodeJsonPayload(&Tinput)
	if err != nil {
		rest.Error(Rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("input: ", Tinput)
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/page", postPage),
		rest.Post("/userID", postUserID),
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server started.")

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":5200", api.MakeHandler()))
}

