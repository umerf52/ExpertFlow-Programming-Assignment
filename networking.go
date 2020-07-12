package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	logger.Println("starting API server")
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/api/v1.0/queue/list", api1)
	myRouter.HandleFunc("/api/v1.0/queue/detail", api2)
	myRouter.HandleFunc("/api/v1.0/queue/service", api3)
	myRouter.HandleFunc("/api/v1.0/queue/enqueue", api4).Methods("POST")
	myRouter.HandleFunc("/api/v1.0/queue/renege/{id}", api5).Methods("DELETE")
	myRouter.HandleFunc("/api/v1.0/SystemInfo", api6)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	port := ":10000"
	logger.Println("API server started listening at port" + port)
	log.Fatal(http.ListenAndServe(port, myRouter))
}

func api1(w http.ResponseWriter, r *http.Request) {
	logger.Printf("Endpoint Hit: /api/v1.0/queue/list")
	s1Struct := selection1(&PQ, false)
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	enc.Encode(s1Struct)
}

func api2(w http.ResponseWriter, r *http.Request) {
	logger.Printf("Endpoint Hit: /api/v1.0/queue/detail")
	s2Struct := selection2(&PQ, false)
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	enc.Encode(s2Struct)
}

func api3(w http.ResponseWriter, r *http.Request) {
	logger.Printf("Endpoint Hit: /api/v1.0/queue/service")
	w.Header().Set("Content-Type", "application/json")
	s3Struct, errorStruct, err := selection3(&PQ, false)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	if err != nil {
		enc.Encode(errorStruct)
	} else {
		enc.Encode(s3Struct)
	}
}

func api4(w http.ResponseWriter, r *http.Request) {
	tempTime := time.Now()
	logger.Printf("Endpoint Hit: /api/v1.0/queue/enqueue")
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	enc.SetIndent("", "    ")

	// get the body of our POST request
	reqBody, _ := ioutil.ReadAll(r.Body)
	cr := CustomerRequest{}
	json.Unmarshal(reqBody, &cr)
	cr.EnqueueTime = tempTime

	s4Struct, err := selection4(&PQ, &cr, false)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		enc.Encode(Selection4ErrorStruct{Error: "MAX_CAPACITY_REACHED", Msg: err.Error()})
	} else {
		enc.Encode(s4Struct)
	}
}

func api5(w http.ResponseWriter, r *http.Request) {
	logger.Printf("Endpoint Hit: /api/v1.0/queue/renege/")
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	enc.SetIndent("", "    ")

	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	idStr := vars["id"]
	idInt, _ := strconv.Atoi(idStr)

	s5Struct, err := selection5(&PQ, idInt, false)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		enc.Encode(ErrorStruct{Msg: err.Error()})
	} else {
		enc.Encode(s5Struct)
	}
}

func api6(w http.ResponseWriter, r *http.Request) {
	logger.Printf("Endpoint Hit: /api/v1.0/SystemInfo")
	w.Header().Set("Content-Type", "application/json")
	s6Struct, errorStruct, err := selection6(&PQ, false)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	if err != nil {
		enc.Encode(errorStruct)
	} else {
		enc.Encode(s6Struct)
	}
}
