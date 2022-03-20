package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	id, _ := strconv.Atoi(r.PostForm.Get("id"))
	userName := r.PostForm.Get("name")
	clients[id] = userName
	data := make(map[string][]string)
	data["connect_users"] = getUserList()
	fmt.Println(getUserList())
	err := app.WsClient.Trigger("public-channel", "update-users", data)
	if err != nil {
		log.Println(err)
	}
}
