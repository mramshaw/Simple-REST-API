package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

var apiVersion = "v1"

type Person struct {
    ID        string   `json:"id,       omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname, omitempty"`
    Address   *Address `json:"address,  omitempty"`
}

type Address struct {
    City  string `json:"city, omitempty"`
    State string `json:"state,omitempty"`
}

func getPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    person := getPerson(params["id"])
    // the best way to check for an empty Person
    if person.ID == "" {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(person)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(person)
}

func getPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(getPeople())
}

func createPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person Person
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createPerson(person))
}

func modifyPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person Person
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    matched, people := modifyPerson(person)
    if !matched {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(people)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(people)
}

func deletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    matched, people := deletePerson(params["id"])
    if !matched {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(people)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(people)
}

func init() {
    createPerson(Person{ID: "1", Firstname: "Fred", Lastname: "Flintstone", Address: &Address{City: "Bedrock", State: "AK"}})
    createPerson(Person{ID: "2", Firstname: "Wilma", Lastname: "Flintstone"})
    createPerson(Person{ID: "3", Firstname: "Barney", Lastname: "Rubble", Address: &Address{City: "Bedrock"}})
    createPerson(Person{ID: "4", Firstname: "Betty", Lastname: "Rubble"})
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/" + apiVersion + "/people", getPeopleEndpoint).Methods("GET")
    router.HandleFunc("/" + apiVersion + "/people/{id}", getPersonEndpoint).Methods("GET")
    router.HandleFunc("/" + apiVersion + "/people/{id}", createPersonEndpoint).Methods("POST")
    router.HandleFunc("/" + apiVersion + "/people/{id}", modifyPersonEndpoint).Methods("PUT")
    router.HandleFunc("/" + apiVersion + "/people/{id}", deletePersonEndpoint).Methods("DELETE")
    log.Print("Now listening on http://localhost:8100 ...")
    log.Fatal(http.ListenAndServe(":8100", handleCORS(router)))
}

func handleCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        origin := req.Header.Get("Origin")
        if origin != "" {
            // define the hosts we will service
            if origin == "http://localhost:3200" {
                w.Header().Set("Access-Control-Allow-Origin", origin)
            } else {
                return
            }
        }
        if req.Method == "OPTIONS" {
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
            return
        }
        next.ServeHTTP(w, req)
    })
}
