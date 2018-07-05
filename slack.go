package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// 記得調整 ID
const webhookURL = `https://hooks.slack.com/services/______________/______________`

var (
	iconEmoji string // ':ghost:'
	iconURL   string // 'https://slack.com/img/icons/app-57.png'
	text      string // 'Hello, world!'
	username  string // 'BOT'
	toUser    string
	toChannel string
)

// main 之前自動執行, 將 cli 參數帶入變數中
func init() {
	flag.StringVar(&username, "name", "", "Display name of webhook message")
	flag.StringVar(&username, "n", "", "Display name of webhook message (shorthand)")

	flag.StringVar(&toChannel, "channel", "", "Slack channel to post, or leave blank to default channel of webhook")
	flag.StringVar(&toChannel, "c", "", "Slack channel to post, or leave blank to default channel of webhook (shorthand)")

	flag.StringVar(&toUser, "user", "", "Message send to specific user")
	flag.StringVar(&toUser, "u", "", "Message send to specific user (shorthand")

	flag.StringVar(&iconEmoji, "emoji", "", "Emoji tag to display as avatar, e.g. ghost")
	flag.StringVar(&iconURL, "image", "", "The URL of image to display as avatar, e.g. https://png.icons8.com/color/50/000000/anonymous-mask.png")

	flag.Parse()
}

/**
 * 生成 incoming webhook payload, 參考 incoming webhook API doc
 *
 * @reference https://kkbox.slack.com/services/B7BS940Q1#service_setup
 */
func buildPayload() map[string]string {
	payload := make(map[string]string)

	var buffer [1024]byte
	n, err := os.Stdin.Read(buffer[:])

	if err != nil {
		fmt.Println("ERROR: ", err)
		return payload
	}

	// 將文字使用 ``` 包裝, 維持 format, 方便將 cli output 呈現
	if n > 0 {
		payload["text"] = "```\n" + string(buffer[:]) + "\n```"
	}

	// channel - 兩種格式: #channel / @username
	if len(toChannel) > 0 {
		payload["channel"] = "#" + toChannel
	} else if len(toUser) > 0 {
		payload["channel"] = "@" + toUser
	}

	// 自訂 message 頭像, 這邊不確定 emoji / url 文件上沒標示哪個優先級高
	if len(iconEmoji) > 0 {
		payload["icon_emoji"] = ":" + iconEmoji + ":"
	}

	if len(iconURL) > 0 {
		payload["icon_url"] = iconURL
	}

	// 自訂 message 送出者名稱
	if len(username) > 0 {
		payload["username"] = username
	}

	return payload
}

func main() {
	payload := buildPayload()

	// Map -> JSON
	jsonPayload, _ := json.Marshal(payload)

	// Form Data
	data := url.Values{}
	data.Set("payload", string(jsonPayload))

	// POST Request
	client := &http.Client{}
	r, _ := http.NewRequest("POST", webhookURL, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	// Send Request
	resp, _ := client.Do(r)
	fmt.Println(resp.Status)
}
