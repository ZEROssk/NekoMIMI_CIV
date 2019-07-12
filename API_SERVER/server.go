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
	PNums := req.PathParam("Page")
	log.Println("Page Number is ", PNums)

	PNumi, err := strconv.Atoi(PNums)
	if err != nil {
		rest.Error(Rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if PNumi != 0 {
		Rw.WriteJson(&Result_JSON{
			"Page number is "+PNums,
		})
	} else {
		rest.Error(Rw, "Page number is required", 400)
	}
}

// https://host-name:port/api/v1/twimg/data/search/{UserID}/{PageNum}
func API_twimg_search(Rw rest.ResponseWriter, req *rest.Request) {
      input := TwiID_Page_Input{}
      err := req.DecodeJsonPayload(&input)
      if err != nil {
              rest.Error(Rw, err.Error(), http.StatusInternalServerError)
              return
      }

      log.Printf("input: %#v", input)

      PNum, _ := input.Page.Int64()
      var s string
      s = strconv.Itoa(PNum)

      if input.UserID && input.Page == "" {
              rest.Error(Rw, "UserID & Page number is required", 400)
      } else if input.Page == "" {
              Rw.WriteJson(&Result_JSON{
                      input.UserID,
              })
              moduleDB.DB()
      } else if PNum != 0 {
              Rw.WriteJson(&Result_JSON{
                      input.UserID+" "+s,
              })
              moduleDB.DB()
      } else {
              rest.Error(Rw, "Page number is required", 400)
      }
}

// https://host-name:port/api/v1/twimg/data/original/{UserID}/{ImageID}
func API_twimg_original(Rw rest.ResponseWriter, req *rest.Request) {
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/api/v1/twimg/data/page/:Page", API_twimg),
		rest.Post("/api/v1/twimg/data/search/:UserID/:PageNum", API_twimg_search),
		rest.Post("/api/v1/twimg/data/original/:UserID/:ImageID", API_twimg_original),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started.")

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":5200", api.MakeHandler()))
}

