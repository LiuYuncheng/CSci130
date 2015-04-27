package static

import(
"html/template"
"net/http"
"github.com/gorilla/sessions"
"github.com/gorilla/context"
)
var globalTemp *template.Template 

var store = sessions.NewCookieStore([]byte("cool-kid-key"))

func init(){
http.HandleFunc("/", handler)
http.HandleFunc("/message", message)
http.HandleFunc("/readmsg", readmsg)
globalTemp = template.Must(template.ParseFiles("index.html", "indexform.html", "readmsg.html"))
http.ListenAndServer(":8080", context.ClearHandler(http.DefaultServeMux))
}

func handler(w http.ResponseWriter, r *http.Request){
err := globalTemp.ExecuteTemplate(w, "index.html", nil)
if err != nil{
http.Error(w, err.Error(), http.StatusInternalServerError)
}
}

func message(w http.ResponseWriter, r *http.Request){
session, _ := store.Get(r, "cool-kid-name")
session.Values["name"] = r.FormValue("name")
session.Values["msg"]=r.FormValue("msg")
session.Save(r,w)
err := globalTemp.ExecuteTemplate(w, "indexform.html",nil)
if err != nil{
http.Error(w, err.Error(), http.StatusInternalServerError)
}
}

func readmsg(w http.ResponseWriter, r *http.Request){
session, _ := store.Get(r, "cool-kid-name")
bname := session.Values["name"]
bmessage : session.Values["msg"]
session.Save(r,w)
err := globalTemp.ExecuteTemplate(w, "readmsg.html", struct{Name interface{}; Message interface{};}{bname,bmessage})
if err != nil{
http.Error(w, err.Error(), http.StatusInternalServerError)
}
}

