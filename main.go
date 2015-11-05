package robification

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func Send(p *fdChat) error {
	url := "http://jrobles.net:1337/v1/flowdock/chat"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(p.Content)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Token", p.Flow_Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	res := apiResponse{}
	json.Unmarshal([]byte(body), &res)

	// Hard-coding one response for now...
	if res.Messages[0].Status == "200 OK" {
		return nil
	}
	return errors.New(res.Messages[0].Status)

}

func NewFdChat(flowToken string, content string) *fdChat {
	chat := &fdChat{
		Flow_Token: flowToken,
		Content:    content,
	}
	return chat
}

func NewSlackChat(token string, content string) *slackChat {
	chat := &slackChat{
		Token:   token,
		Content: content,
	}
	return chat
}
