package rest

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

/*
Start the restful server
*/
func Start(host string) {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		/*
			the main api for exchange on smartraiden and lnd
		*/
		rest.Post("/api/1/register-exchange", RegisterExchange),

		/*
			other api
		*/
		rest.Get("/api/1/balance", GetBalance),
	)
	if err != nil {
		log.Printf("maker router :%s", err)
	}
	api.SetApp(router)
	log.Println("http start and listen at", host)
	err = http.ListenAndServe(host, api.MakeHandler())
	if err != nil {
		panic(err)
	}
}
