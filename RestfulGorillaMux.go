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

var people []Person

func getPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range people {
        if item.ID == params["id"] {
            w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    // If no match found, return 'empty' Person
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(&Person{})
}

func getPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(people)
}

func createPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person Person
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(people)
}

func modifyPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person Person
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            people = append(people, person)
            break
        }
    }
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(people)
}

func deletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
    }
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(people)
}

func init() {
    people = append(people, Person{ID: "1", Firstname: "Fred", Lastname: "Flintstone", Address: &Address{City: "Bedrock", State: "AK"}})
    people = append(people, Person{ID: "2", Firstname: "Wilma", Lastname: "Flintstone"})
    people = append(people, Person{ID: "3", Firstname: "Barney", Lastname: "Rubble", Address: &Address{City: "Bedrock"}})
    people = append(people, Person{ID: "4", Firstname: "Betty", Lastname: "Rubble"})
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/" + apiVersion + "/people", getPeopleEndpoint).Methods("GET")
    router.HandleFunc("/" + apiVersion + "/people/{id}", getPersonEndpoint).Methods("GET")
    router.HandleFunc("/" + apiVersion + "/people/{id}", createPersonEndpoint).Methods("POST")
    router.HandleFunc("/" + apiVersion + "/people/{id}", modifyPersonEndpoint).Methods("PUT")
    router.HandleFunc("/" + apiVersion + "/people/{id}", deletePersonEndpoint).Methods("DELETE")
    log.Print("Now listening on http://localhost:8100 ...")
    log.Fatal(http.ListenAndServe(":8100", router))
}
