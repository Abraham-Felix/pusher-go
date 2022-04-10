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

func (repo *DBRepo) NewPersonalMessage(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]string)
	data["message"] = r.PostForm.Get("message")
	data["username"] = r.PostForm.Get("sender")
	fmt.Println(r.PostForm)
	err := app.WsClient.Trigger("public-channel", r.PostForm.Get("receiver"), data)
	if err != nil {
		log.Println(err)
	}

	chanName := "s" + r.PostForm.Get("sender")
	privateSendLog(r.PostForm.Get("message"), chanName, r.PostForm.Get("receiver"))
}

func privateSendLog(msg string, sender string, receiver string) {
	data := make(map[string]string)
	data["message"] = msg
	data["username"] = receiver

	err := app.WsClient.Trigger("public-channel", sender, data)
	if err != nil {
		log.Println(err)
	}

}
