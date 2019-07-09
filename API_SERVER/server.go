package main

import(
	"log"
	"net/http"
	"encoding/json"
	//"strconv"

	"./module"
	"github.com/ant0ine/go-json-rest/rest"
)

type postHelloInput struct {
	Page json.Number
}

type postHelloOutput struct {
	Result string
}

func postHello(w rest.ResponseWriter, req *rest.Request) {
	input := postHelloInput{}
	err := req.DecodeJsonPayload(&input)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// var PageNum int 
	// PageNum, _ = strconv.Atoi(input.Page)

	log.Printf("input: %#v", input)
	log.Printf("input.page: %d", input.Page)

	PNumber, _ := input.Page.Int64()

	if PNumber != 0 {
		w.WriteJson(&postHelloOutput{
			"Page number is ",//+input.Page,
		})
		moduleDB.DB()
	} else {
		rest.Error(w, "Page number is required", 400)
	}
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/hello", postHello),
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server started.")

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":5200", api.MakeHandler()))
}

