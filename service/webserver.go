package service

import (
	"fmt"
	"log"
	"net/http"
)

func homePage( w http.ResponseWriter, r *http.Request ) {
	fmt.Fprint( w, "Hello World!\n" )
}

func StartWebServer( port string ) {

	log.Println( "Starting HTTP Service at " +  port )

	r := NewRouter()
	log.Fatal( http.ListenAndServe( ":" + port, r ) )
}
