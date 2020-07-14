package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// This method is used as a goroutine to handle REST APIs
func handleRequests() {
	logger.Println("starting API server")
	// creates a new instance of a mux router
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/v1.0/queue/list", api1)
	r.HandleFunc("/api/v1.0/queue/detail", api2)
	r.HandleFunc("/api/v1.0/queue/service", api3)
	r.HandleFunc("/api/v1.0/queue/enqueue", api4).Methods("POST")
	r.HandleFunc("/api/v1.0/queue/renege/{id}", api5).Methods("DELETE")
	r.HandleFunc("/api/v1.0/SystemInfo", api6)
	r.HandleFunc("/", allOther)
	port := ":10000"
	logger.Println("API server started listening at port" + port)
	log.Fatal(http.ListenAndServe(port, r))
}

// This method is for Listing Customers in Queue
func api1(w http.ResponseWriter, r *http.Request) {
	logger.Printf("Endpoint Hit: /api/v1.0/queue/list")
	s1Struct := selection1(&PQ, false)
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	enc.Encode(s1Struct)
}

// This method is for Listing Customers details in Queue
func api2(w http.ResponseWriter, r *http.Request) {
	logger.Printf("Endpoint Hit: /api/v1.0/queue/detail")
	s2Struct := selection2(&PQ, false)
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	enc.Encode(s2Struct)
}

// This method is for Servicing Customer Request
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

// This method is for Enqueueing Customer Request
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
	if cr.CustomerName == "" && cr.Description == "" && cr.PriorityWeight == 0 {
		temp := Selection4ErrorStruct{Error: "INVALID_PARAMETERS", Msg: "no body found in PUT request"}
		enc.Encode(temp)
		return
	}
	cr.EnqueueTime = tempTime

	s4Struct, err := selection4(&PQ, &cr, false)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		enc.Encode(Selection4ErrorStruct{Error: "MAX_CAPACITY_REACHED", Msg: err.Error()})
	} else {
		enc.Encode(s4Struct)
	}
}

// This method is for Reneging Customer Request
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

// This method is for getting System Information
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

// Method to handle all other requests
func allOther(w http.ResponseWriter, r *http.Request) {
	logger.Printf("Endpoint Hit: allOther")
	fmt.Fprintf(w, "Try other routers such as: ")
	fmt.Fprintf(w, "/api/v1.0/queue/list")
	fmt.Fprintf(w, "/api/v1.0/queue/detail")
	fmt.Fprintf(w, "/api/v1.0/queue/service")
	fmt.Fprintf(w, "/api/v1.0/queue/enqueue")
	fmt.Fprintf(w, "/api/v1.0/queue/renege/{id}")
	fmt.Fprintf(w, "/api/v1.0/SystemInfo")
}
