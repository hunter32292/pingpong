package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	count      int
	serverAddr string
	targetAddr string
)

func main() {

	flag.StringVar(&serverAddr,
		"serverAddr",
		"localhost:8081",
		"Application host address")
	flag.StringVar(&targetAddr,
		"targetAddr",
		"localhost:8080",
		"The target server address")

	flag.Parse()

	handler := mux.NewRouter()
	handler.HandleFunc("/start", start)
	handler.HandleFunc("/hit", hit)
	http.Handle("/", handler)

	log.Fatal(http.ListenAndServe(serverAddr, handler))
}

func start(w http.ResponseWriter, r *http.Request) {
	go caller() // Send off on a thread to be finished
	return
}

func caller() {
	protocol := "http://"
	endpoint := "/hit"
	method := "GET"

	url := protocol + targetAddr + endpoint
	client := &http.Client{}
	for {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)+100))
		count++
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			log.Fatal(err)
		}
		_, err = client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Count: ", count)
	}

}

func hit(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT!")
	w.WriteHeader(http.StatusOK)
}
