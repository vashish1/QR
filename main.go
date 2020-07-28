package main

import (
	"image/png"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gorilla/mux"
)

var port = os.Getenv("PORT")
func main() {
	r:=mux.NewRouter()

	r.HandleFunc("/QR", homeHandler).Methods("GET")
	r.HandleFunc("/QR/Generated", viewCodeHandler).Methods("POST","GET")
	http.Handle("/",r)

	http.ListenAndServe(":"+port, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./index.html")
	t.Execute(w, "")
}

func viewCodeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	text := r.FormValue("test")
	link := r.FormValue("link")
	if strings.ToLower(link)=="no"{
		link=""
	}
	Data:= text+ " \n "+ link +"\n \n ~ Made by " +name + " :-)"
	qrCode, _ := qr.Encode(Data, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 128, 128)

	png.Encode(w, qrCode)
}