package main

import (
	"net/http"
)
type infoHandler struct{}

func getInfoHandler(res http.ResponseWriter, req *http.Request) {
	data := []byte("getInfo handler called")
	res.WriteHeader(200)
	res.Write(data)
}

func (h infoHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("getData function called")
	res.WriteHeader(200)
	res.Write(data)
}

 func main() {

	//Create the default mux
	mux := http.NewServeMux()

	//Handling the /v1/getInfo. The handler is a function here
	mux.HandleFunc("/v1/getInfo", getInfoHandler)

	//Handling the /v1/getData. The handler is a function here
	sInfoHandler := infoHandler{}
	mux.Handle("/v1/getData", sInfoHandler)

	//Create the server.
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	  }

	  s.ListenAndServe()
 }