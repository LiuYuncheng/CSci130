package main

import(
"fmt"
"html/template"
"log"
"net/http"
"os"
"io/ioutil"
)

func main(){
http.HandleFunc("/", root)
http.HandleFunc("/results", showResults)
fmt.Println("listening...")
err:=http.ListenAndServe(GetPort(),nil)
if err != nil{
log.Fatal("ListenAndServe: ", err)
return
}
}

func root(w http.ResponseWriter, r *http.Request){
rootForm, err:=ioutil.ReadFile("pages/rootForm.html");
if err != nil{
fmt.Fprint(w, "404 Not Found")
}
fmt.Fprint(w,string(rootForm))
}

var results, _ = ioutil.ReadFile("pages/rootForm.html");
var resultsTemplate = template.Must(template.New("results").Parse(string(results)))

func showResults(w http.ResponseWriter, r *http.Request){
strEntered:=r.FormValue("str")
if strEntered=="Yuncheng"{
err:=resultsTemplate.Execute(w,"Right!")
if err != nil{
http.Error(w, err.Error(),http.StatusInternalServerError)
}
}else{
err:=resultsTemplate.Execute(w,"wrong")
if err != nil{
http.Error(w, err.Error(), http.StatusInternalServerError)
}
}
}

func GetPort() string{
var port=os.Geteny("PORT")
if port==""{
port="4747"
fmt.Println("no port environment variable detected defaulting to " + port)
}
return ":" + port
