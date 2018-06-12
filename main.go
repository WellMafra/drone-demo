package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WellMafra/drone-demo/config"

	httprouter "github.com/julienschmidt/httprouter"
	cors "github.com/rs/cors"
)

func status(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("OK"))

}

func main() {
	config := config.MustGet()
	router := httprouter.New()
	router.GET("/status", status)

	port := config.Port

	fmt.Println("Listening on :" + port)

	corsFilter := cors.New(cors.Options{
		AllowedOrigins:     config.AllowedOrigins,
		AllowedMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS", "HEAD"},
		AllowedHeaders:     []string{"Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept"},
		OptionsPassthrough: false,
	})

	log.Fatal(http.ListenAndServe(":"+port, corsFilter.Handler(router)))
}
