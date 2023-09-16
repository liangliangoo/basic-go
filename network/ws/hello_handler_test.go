package ws

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"testing"
)

func TestHelloHandlerWs(t *testing.T) {

	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
