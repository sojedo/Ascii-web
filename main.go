package main

import(
	"net/http"
	"html/template"
)
func homehandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}
    tmpl.Execute(w, nil)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
    text := r.FormValue("text")
	if text == "" {
    http.Error(w, "Bad Request: text is empty", http.StatusBadRequest)
    return
}
    banner := r.FormValue("banner")
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
    http.Error(w, "Bad Request: invalid banner", http.StatusBadRequest)
    return
}
    
    result, err := Generate(text, "banners/"+banner+".txt")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, "Template not found", http.StatusNotFound)
        return
    }
    tmpl.Execute(w, result)
}

func main() {
	http.HandleFunc("/", homehandler)
	http.HandleFunc("/ascii-art", asciiHandler)
	http.ListenAndServe(":8080", nil)
}