package httprouter

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"testing"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(":8010", router))
}
