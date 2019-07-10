package main

import(
	"log"
	"net/http"
	"encoding/json"
	//"strconv"

	"./module"
	"github.com/ant0ine/go-json-rest/rest"
)

type Page_Input struct {
	Page json.Number
}

type TwiID_Page_Input struct {
	Page json.Number
	UserID string
}

type Result_JSON struct {
	Result string
}

func postPage(Rw rest.ResponseWriter, req *rest.Request) {
	input := Page_Input{}
	err := req.DecodeJsonPayload(&input)
	if err != nil {
		rest.Error(Rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("input: %#v", input)

	PNum, err := input.Page.Int64()
	if err != nil {
		rest.Error(Rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if PNum != 0 {
		Rw.WriteJson(&Result_JSON{
			"Page number is OK",
		})
		moduleDB.DB()
	} else {
		rest.Error(Rw, "Page number is required", 400)
	}
}

func postTwiID_Page(Rw rest.ResponseWriter, req *rest.Request) {
	input := TwiID_Page_Input{}
	err := req.DecodeJsonPayload(&input)
	if err != nil {
		rest.Error(Rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("input: %#v", input)

	PNum, _ := input.Page.Int64()

	if input.UserID && input.Page == nil {
		rest.Error(Rw, "UserID & Page number is required", 400)
	} else if input.Page == nil {
		Rw.WriteJson(&Result_JSON{
			input.UserID,
		})
		moduleDB.DB()
	} else if PNum != 0 {
		Rw.WriteJson(&Result_JSON{
			input.UserID+" "+input.Page,
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
		rest.Post("api/Page", postPage),
		rest.Post("api/TwiID_Page", postTwiID_Page),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started.")

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":5200", api.MakeHandler()))
}

