package api

import(
	"log"
	"net/http"
	"encoding/json"
	"strconv"

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

// https://host-name:port/api/v1/image_viewer/images/twimg/data?p={PageNum}
func API_twimg(Rw rest.ResponseWriter, req *rest.Request) {
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

// https://host-name:port/api/v1/image_viewer/images/twimg/data/search?id={UserID}&p={PageNum}
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

// https://host-name:port/api/v1/image_viewer/images/twimg/data/original?id={UserID}&img={ImageID}
func API_twimg_original(Rw rest.ResponseWriter, req *rest.Request) {
}

