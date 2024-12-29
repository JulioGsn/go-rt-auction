package main

import (
	"log"
	"net/http"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	counter := 0
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer conn.Close()
//
// 	for {
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			break
// 		}
// 		log.Printf("Message that was sent by client: %s", p)
// 		counter++
//
// 		err = conn.WriteMessage(messageType, p)
// 		if err != nil {
// 			log.Println("write: ", err)
// 			break
// 		}
// 		log.Println(counter)
// 	}
// }

func main() {
	hub := newHub()
	go hub.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

  err := http.ListenAndServe(":8000", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
