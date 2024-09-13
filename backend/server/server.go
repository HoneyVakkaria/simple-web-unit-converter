package server

import (
	"github/honeyvakkaria/unit-converter/converter"
	"html/template"
	"log"
	"net/http"
)

type FunctionTemplate func(string, string, string) (float64, error)

type Data struct {
	From     string
	To       string
	Amount   string
	Result   float64
	Redirect string
}

func HandleServer() {
	http.HandleFunc("/submitLength", func(w http.ResponseWriter, r *http.Request) {
		submitHandler(w, r, converter.ConvertLength)
	})
	http.HandleFunc("/submitWeight", func(w http.ResponseWriter, r *http.Request) {
		submitHandler(w, r, converter.ConvertWeight)
	})
	http.HandleFunc("/submitTemperature", func(w http.ResponseWriter, r *http.Request) {
		submitHandler(w, r, converter.ConvertTemperature)
	})
	http.HandleFunc("/weight", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/weight.html")
	})
	http.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/temperature.html")
	})
	http.HandleFunc("/length", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/length.html")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/length", http.StatusSeeOther)
	})

	log.Println("Launch server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request, convert FunctionTemplate) {
	data := Data{}
	if r.Method == http.MethodPost {
		data.From = r.FormValue("from")
		data.To = r.FormValue("to")
		data.Amount = r.FormValue("amount")
		data.Redirect = r.FormValue("redirect")
	} else {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("../frontend/result.html")
	if err != nil {
		http.Error(w, "error loading template", http.StatusInternalServerError)
		return
	}

	data.Result, err = convert(data.From, data.To, data.Amount)
	if err != nil {
		http.Error(w, "invalid amount of length value", http.StatusBadRequest)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "error executing template", http.StatusInternalServerError)
		return
	}
}
