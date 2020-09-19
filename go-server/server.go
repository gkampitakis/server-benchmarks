package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/loadavg"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/network"
	"github.com/mackerelio/go-osstat/uptime"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ServerStats struct {
	CPU     *cpu.Stats      `json:"cpu"`
	Memory  *memory.Stats   `json:"memory"`
	Loadavg *loadavg.Stats  `json:"loadavg"`
	Uptime  time.Duration   `json:"uptime"`
	Network []network.Stats `json:"network"`
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

func HealthCtrl(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	stats := getStats()

	_, err := w.Write(newJSONResponse("Health Status", stats))

	if err != nil {
		panic(err)
	}
}

func PostDataCtrl(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	router.GET("/health", HealthCtrl)
	router.POST("/postdata", PostDataCtrl)

	log.Printf("Go Server listening on port %v 🚀", 5000)

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

func getStats() *ServerStats {
	c, err := cpu.Get()
	m, err := memory.Get()
	n, err := network.Get()
	u, err := uptime.Get()
	l, err := loadavg.Get()

	if err != nil {
		panic(err)
	}

	return &ServerStats{
		Memory:  m,
		CPU:     c,
		Network: n,
		Uptime:  u,
		Loadavg: l,
	}
}
