package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	gotenor "github.com/natsuokawai/go-tenor"
)

func main() {
	url := "https://discord.com/api/webhooks/XXXXXXXXXXXXXXXxx"
	t := gotenor.NewTenor("YYYYYYYYY")
	gifURL := getGifURL(t)

	params := &discordgo.WebhookParams{
		Username: "時報Bot", Content: timeSignalMessage(),
	}
	params.Embeds = []*discordgo.MessageEmbed{
		{Image: &discordgo.MessageEmbedImage{URL: gifURL}},
	}
	sendWebhook(url, params)
}

func timeSignalMessage() string {
	hour := time.Now().Round(time.Hour).Hour()
	return strconv.Itoa(hour) + "時です"
}

func getGifURL(t *gotenor.Tenor) string {
	data, err := t.GetRandom("cat", "", "", 3)
	if err != nil {
		fmt.Println(err)
	}
	return gotenor.GetAllGifURLS(*data)[2]

}

// discord request
func sendWebhook(url string, params *discordgo.WebhookParams) {
	j, err := json.Marshal(params)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if resp.StatusCode == 204 {
		fmt.Println("successfully sent message")
		return
	}

	fmt.Fprintln(os.Stderr, resp)
}
