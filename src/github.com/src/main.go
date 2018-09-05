package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	//http.ListenAndServe(":8000", http.FileServer(http.Dir("C:\\Users\\Test\\Go\\src\\lemonade-app\\public")))
	templates := populateTemplates()
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		requestedFile := request.URL.Path[1:]
		temp := templates.Lookup(requestedFile + ".html")

		if temp != nil {
			err := temp.Execute(responseWriter, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			responseWriter.WriteHeader(http.StatusNotFound)
		}
	})

	http.Handle("/img/", http.FileServer(http.Dir("C:\\Users\\Test\\Go\\src\\lemonade-app\\public")))
	http.Handle("/css/", http.FileServer(http.Dir("C:\\Users\\Test\\Go\\src\\lemonade-app\\public")))

	http.ListenAndServe(":8000", nil)

}

func populateTemplates() *template.Template {
	temp := template.New("templates")
	const basePath = "C:\\Users\\Test\\Go\\src\\lemonade-app\\templates"
	template.Must(temp.ParseGlob(basePath + "/*.html"))
	return temp
}
