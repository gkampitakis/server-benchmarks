package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"
	"github.com/julienschmidt/httprouter"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BodyJSON interface{}

func ListenCtrl(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()

	fmt.Println(JSONStringify(query))

	_, err := w.Write(newJSONResponse("Responding back", query))

	if err != nil {
		panic(err)
	}
}

func PostDataCtrl(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()

	var body BodyJSON

	err := json.NewDecoder(r.Body).Decode(&body)

	fmt.Println(body)

	if err != nil {
		panic(err)
	}

	_, err = w.Write(newJSONResponse("Post response", body))

	if err != nil {
		panic(err)
	}
}

func main() {

	defer handerError()

	router := httprouter.New()
	router.GET("/listen", ListenCtrl)
	router.POST("/postdata", PostDataCtrl)

	log.Printf("Server listening on port %v", 5000)

	error := http.ListenAndServe(":5000", router)

	if error != nil {
		log.Fatal(error)
	}
}

func handerError() {

	e := recover()

	if e != nil {

		log.Fatal(e)
		return

	}

	log.Fatalln("Unhandled error, server shuting down")

}

func JSONStringify(f interface{}) string {
	value, err := json.Marshal(f)
	if err != nil {
		panic(err)
	}

	return string(value)
}

func newJSONResponse(message string, data interface{}) (res []byte) {
	res, err := json.Marshal(&Response{Message: message, Data: data})

	if err != nil {
		panic(err)
	}

	return
}
