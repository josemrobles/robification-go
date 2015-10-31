package robification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Send(p *fdChat) {
	url := "http://jrobles.net:1337/v1/flowdock/chat"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(p.Content)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
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
	fmt.Println(res.Items[0].Name)
}

func NewFdChat(flowToken string, content string) *fdChat {
	chat := &fdChat{
		Flow_Token: flowToken,
		Content:    content,
	}
	return chat
}
