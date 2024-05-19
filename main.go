package main

import (
	"log"
	"net/http"

	"github.com/OscarG356/web_server_GO/Handlers"
	"github.com/OscarG356/web_server_GO/Repository"
)

func main() {
	bd := repository.NewDataBase()

	handler := handlers.NewHandlerHero(bd)

	mux := http.NewServeMux()
	
	mux.HandleFunc("/api/{superhero}",handler.GetHero())

	log.Fatal(http.ListenAndServe(":8080", mux))
}
