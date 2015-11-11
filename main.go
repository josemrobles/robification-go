package robification

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Send(messages ...chatMessage) error {
	doop := make([]apiResponses, len(messages))
	for k, message := range messages {
		req, err := http.NewRequest("POST", message.Url, bytes.NewBuffer([]byte(message.Text)))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Token", message.Token)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		res := apiResponses{}
		json.Unmarshal([]byte(body), &res)

		doop[k] = res
		apiResponse(doop)
		println(doop)
	}
	return nil
}

func apiResponse(oof []apiResponses) {
	for _, response := range oof {
		for _, meh := range response.Responses {
			println(string(meh.Status))
		}
	}
}

func NewChatMessage(messageType string, token string, content string) *chatMessage {
	var url string = ""
	if messageType == "flowdock" {
		url = "http://jrobles.net:1337/v1/flowdock/chat!!!"
	} else {
		url = "http://jrobles.net:1337/v1/slack/chat"
	}
	chat := &chatMessage{
		Url:   url,
		Token: token,
		Text:  content,
	}
	return chat
}
