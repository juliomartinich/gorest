package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func restCustomers(w http.ResponseWriter, r *http.Request){
    // entry del api rest "Customers"

    switch r.Method {
    case "GET":
        fmt.Println("Endpoint Hit: returnAllArticles GET")
        //fmt.Fprintf(w, "GET! %v\n", r.URL.Path)
        fmt.Fprintf(w, "{ name:John, age:30, car:null }")
    case "POST":
        fmt.Println("Endpoint Hit: returnAllArticles POST")
        fmt.Fprintf(w, "POST! %v\n", r.URL.Path)
    default:
        fmt.Println("Endpoint Hit: returnAllArticles default")
        fmt.Fprintf(w, "default! %v\n", r.URL.Path)
    }
}

func restLibs(w http.ResponseWriter, r *http.Request){
    // entry del api rest "libs"

    fmt.Fprintf(w, "Welcome %v %v", r.URL.Path, r.Method)
    fmt.Println("Endpoint Hit: libs")
}

func main() {

    http.HandleFunc("/", homePage)
    http.HandleFunc("/customers", restCustomers)
    http.HandleFunc("/libs", restLibs)

    log.Fatal(http.ListenAndServe(":10000", nil))
}


