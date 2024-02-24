package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		// Handle conenction error
		if err != nil {
			fmt.Printf("%s error: %s\n", conn.RemoteAddr(), string(err.Error()))
			return
		}

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()

			// Handle message read error
			if err != nil {
				fmt.Printf("%s Unable to read message, error: %s\n", conn.RemoteAddr(), string(err.Error()))

				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				fmt.Printf("%s Unable to respond, error: %s\n", conn.RemoteAddr(), string(err.Error()))
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}
