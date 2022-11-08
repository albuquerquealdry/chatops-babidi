package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func TelegramCommit(repository string, commitHash string, branch string, user string, timestamp string) {
	var token = os.Getenv("TOKEN_TELEGRAM")
	text := fmt.Sprintf("	🔮Babidi Chatops🔮\n \n Action:Commit🎈 \n-------------------------------------\n \n📦 Repository : %s \n🏟️ Commit : %s\n📓Branch : %s\n👤User : %s\n🕒Timestamp : %s ", repository, commitHash, branch, user, timestamp)
	called := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=-673484328&text=%s", token, url.QueryEscape(text))
	resp, err := http.Get(called)

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

func TelegramPull(repository string, commitHash string, branch string, user string, timestamp string) {
	var token = os.Getenv("TOKEN_TELEGRAM")
	text := fmt.Sprintf("	🔮Babidi Chatops🔮\n \n Action:Pull Request🔄 \n-------------------------------------\n📦 Repository : %s \n⭐ OriginBranch : %s\n🌌 FutureBranch : %s\n👤User : %s\n🕒Timestamp : %s ", repository, commitHash, branch, user, timestamp)
	called := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=-673484328&text=%s", token, url.QueryEscape(text))
	resp, err := http.Get(called)

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
