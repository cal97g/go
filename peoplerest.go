package main

import (
  "fmt"
  "log"
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
)

type Person struct {
  ID string `json:"id,ommitempty"`
  Firstname string `json:"firstname,ommitempty"`
  Lastname string`json:"lastname,ommitempty"`
  Address *Address `json:"address,ommitempty"`
}

type Address struct {
  City string `json:"city,ommitempty"`
  State string `json:"state,ommitempty"`
}

var people []Person


func PersonGetHandler(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  for _, item := range people {
    if item.ID == params["ID"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  http.Error(w, "Record not found.", http.StatusNotFound)
}

func PersonsGetHandler(w http.ResponseWriter, req *http.Request) {
  json.NewEncoder(w).Encode(people)
}

func PersonPostHandler(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  var person Person
  _ = json.NewDecoder(req.Body).Decode(&person)

  person.ID = params["ID"]
  people = append(people, person)
  json.NewEncoder(w).Encode(people)

}

func PersonDeleteHandler(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  for index, item := range people {
    if item.ID == params["ID"] {
      // Delete item from our slice by index if it exists
      // see https://stackoverflow.com/a/25025536/1998692
      people = append(people[:index], people[index+1:]...) //wtf
      break
    }
  }
  json.NewEncoder(w).Encode(people)
}

func ChocolatesHandler(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)

  switch req.Method {
  case "GET":
    // check map for id key, if present the user is requesting a specific record
    // https://stackoverflow.com/a/2050629/1998692
    if val, ok := params["ID"]; ok {
      // return specific record
      fmt.Println("%s", val)
    }
    // return all chocolates

  case "POST":
    // add record
  case "PUT":
    // update record
  case "DELETE":
    // delete record
  default:
    http.Error(w, "Method not allowed.", 405)
  }
}

func main() {
  router := mux.NewRouter()
  people = append(people, Person{ID: "1", Firstname: "Callam", Lastname: "Delaney", Address: &Address{City: "London", State: "London"}})
  people = append(people, Person{ID: "2", Firstname: "John", Lastname: "Doe", Address: &Address{City: "Telford", State: "Shropshire"}})

  // can we do a single function for multiple methods like Flask?
  // We can with .Methods("GET", "POST", "DELETE")
  // https://github.com/gorilla/mux
  router.HandleFunc("/people", PersonsGetHandler).Methods("GET")
  router.HandleFunc("/people/{ID}", PersonGetHandler).Methods("GET")
  router.HandleFunc("/people/{ID}", PersonPostHandler).Methods("POST")
  router.HandleFunc("/people/{ID}", PersonDeleteHandler).Methods("DELETE")
  router.HandleFunc("/chocolate/", ChocolatesHandler).Methods("GET")
  router.HandleFunc("/choclate/{ID}", ChocolatesHandler).Methods("GET", "PUT", "POST", "DELETE")
  log.Fatal(http.ListenAndServe(":12345", router))
}
