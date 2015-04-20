package main

import (
"fmt"
"html/template"
"log"
"net/http"
"os"
"path"
)

func main() {
fs := http.FileServer(http.Dir("public"))
http.Handle("/tmpfiles/", http.StripPrefix("/tempfiles/", fs))

http.HandleFunc("/", ServeTemplate)

fmt.Println("Listening...")
err := http.ListenAndServe(GetPort(), nil)
if err != nil {
log.Fatal("ListenAndServe: ", err)
return
}
}


func GetPort() string {
var port = os.Getenv("PORT")

if port == "" {
port = "4747"
fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
}
return ":" + port
}

func ServeTemplate(w http.ResponseWriter, r *http.Request){
lp := path.Join("templates", "layout.html") //templates/layout.html 		// http://golang.org/pkg/path/#Join		
fp := path.Join("templates", r.URL.Path) //templates/r.URL.Path[1:]


info, err := os.Stat(fp)  // http://golang.org/pkg/os/#Stat
if err != nil {
if os.IsNotExist(err) { //http://golang.org/pkg/os/#IsNotExist
http.NotFound(w, r) // http://golang.org/pkg/net/http/#NotFound
return
}
}


if info.IsDir() {  // http://golang.org/pkg/os/#FileMode.IsDir
http.NotFound(w, r)
return
}



templates, _ := template.ParseFiles(lp, fp)  // http://golang.org/pkg/html/template/#ParseFiles // http://golang.org/pkg/html/template/#Template.ParseFiles
if err != nil{
fmt.Println(err)
http.Error(w, "500 Internal Server Error", 500)
return
}

templates.ExecuteTemplate(w, "layout", nil)  // http://golang.org/pkg/html/template/#Template.ExecuteTemplate
}

//TO RUN THIS FILE
//http://localhost:4747/indexnew.html
