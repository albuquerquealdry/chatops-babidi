package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/joho/godotenv"
)

func TelegramRander(commitHash string, branch string, user string, timestamp string) {
	godotenv.Load(".env")
	//token_telegram := os.Getenv("API_TOKEN")
	text := fmt.Sprintf("	🔮Babidi Chatops🔮\n🏟️ Commit : %s\n📓Branch : %s\n👤User : %s\n🕒Timestamp : %s ", commitHash, branch, user, timestamp)
	called := fmt.Sprintf("https://api.telegram.org/bot5757162669:AAEXlI_N-mNAitaKzBO5w7TTo7N3Cb73z0E/sendMessage?chat_id=-866409625&text=%s", url.QueryEscape(text))
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
