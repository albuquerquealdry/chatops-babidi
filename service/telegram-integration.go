package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Request() {
	text := "🔮.Babidi Chatops 🔮 | 🏟️ Aviso de alteração de infraestrutura |  📓  Enviroment: teste  🔃  | Region: teste  | 👤  User:  teste  |  🕒  Timestamp: 12344232143123 "

	resp, err := http.Get(`https://api.telegram.org/bot/sendMessage?chat_id=-866409625&text=` + text)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func TelegramRander() {
	Request()
}
