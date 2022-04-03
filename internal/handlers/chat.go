package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func (repo *DBRepo) NewMessage(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]string)
	data["message"] = r.PostForm.Get("message")
	data["username"] = r.PostForm.Get("username")
	fmt.Println(r.PostForm)
	err := app.WsClient.Trigger("public-channel", "new-message", data)
	if err != nil {
		log.Println(err)
	}
}

func (repo *DBRepo) AddChatUser(w http.ResponseWriter, r *http.Request) {
	userName := r.PostForm.Get("name")
	data := getUserList(userName)
	err := app.WsClient.Trigger("public-channel", "add-chat-user", data)
	if err != nil {
		log.Println(err)
	}
}
