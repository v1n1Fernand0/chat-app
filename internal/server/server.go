package server

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func StartServer() {
	fs := http.FileServer(http.Dir("D:/Projetos/chat-app/public"))
	http.Handle("/", fs)

	http.Handle("/ws", websocket.Handler(HandleConnections))

	go HandleMessages()

	log.Println("Servidor iniciado na porta :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
