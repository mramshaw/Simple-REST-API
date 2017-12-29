package main

import (
    // native packages
    "encoding/json"
    "log"
    "net/http"
    // local packages (cannot be installed, TRAITS really)
    "./api"
    "./api/people"
    // GitHub packages
    "github.com/gorilla/mux"
)

var apiVersion = "v1"

func getPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    person := people.GetPerson(params["id"])
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
    json.NewEncoder(w).Encode(people.GetPeople())
}

func createPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person people.Person
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(people.CreatePerson(person))
}

func modifyPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person people.Person
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    matched, people := people.ModifyPerson(person)
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
    matched, people := people.DeletePerson(params["id"])
    if !matched {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(people)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(people)
}

func init() {
    people.CreatePerson(people.Person{ID: "1", Firstname: "Fred", Lastname: "Flintstone", Address: &people.Address{City: "Bedrock", State: "AK"}})
    people.CreatePerson(people.Person{ID: "2", Firstname: "Wilma", Lastname: "Flintstone"})
    people.CreatePerson(people.Person{ID: "3", Firstname: "Barney", Lastname: "Rubble", Address: &people.Address{City: "Bedrock"}})
    people.CreatePerson(people.Person{ID: "4", Firstname: "Betty", Lastname: "Rubble"})
}

func main() {
    router := mux.NewRouter()

    // Health Check
    router.HandleFunc("/ping", api.HealthCheck).Methods("GET")

    // API
    router.HandleFunc("/" + apiVersion + "/people", getPeopleEndpoint).Methods("GET")
    router.HandleFunc("/" + apiVersion + "/people/{id}", getPersonEndpoint).Methods("GET")
    router.HandleFunc("/" + apiVersion + "/people/{id}", createPersonEndpoint).Methods("POST")
    router.HandleFunc("/" + apiVersion + "/people/{id}", modifyPersonEndpoint).Methods("PUT")
    router.HandleFunc("/" + apiVersion + "/people/{id}", deletePersonEndpoint).Methods("DELETE")

    log.Print("Now listening on http://localhost:8100 ...")
    log.Fatal(http.ListenAndServe(":8100", api.HandleCORS(router)))
}
