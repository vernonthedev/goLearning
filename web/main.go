package main

import (
    "fmt"
    "net/http"
)

// function to handle the intial route html page
func index(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "<h1> Hello Bums. </h1>")
}

// function to handle the about route html page
func about(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "<h1> About Us. </h1>")
}
// main func to run the web server
func main(){
    http.HandleFunc("/", index)
    http.HandleFunc("/about", about)
    fmt.Println("[+] Server Starting...")
    http.ListenAndServe(":3000", nil)
}
