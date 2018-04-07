package main

import (
	"fmt"
	"log"
	"net/http"

	httprouter "github.com/julienschmidt/httprouter"
	cors "github.com/rs/cors"
)

func status(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("OK"))
}

func main() {
	router := httprouter.New()
	router.GET("/status", status)

	port := "8080"

	fmt.Println("Listening on :" + port)

	corsFilter := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS", "HEAD"},
		AllowedHeaders:     []string{"Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept"},
		OptionsPassthrough: false,
	})

	log.Fatal(http.ListenAndServe(":"+port, corsFilter.Handler(router)))
}
