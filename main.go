package robification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Payload struct {
	Targets []Target
}

type Target struct {
	Destination_Type     string
	Destination_Sub_Type string
	Data                 robification_fdThread
}

type robification_emailMessage struct {
	From       string
	Subject    string
	Body       string
	Recipients []string
}

type robification_fdThread struct {
	Flow_Token         string
	Event              string
	Author             Author
	Title              string
	External_Thread_Id string
	Thread             Thread
}

type Thread struct {
	Title        string
	Fields       []Field
	Body         string
	External_Url string
	Status       ThreadStatus
}

type ThreadStatus struct {
	Color string
	Value string
}

type Field struct {
	Label string
	Value string
}

type Author struct {
	Name   string
	Avatar string
}

func send(token string, external_id string, title string, label_color string, label_value string, fields []Field) {

	url := "http://jrobles.net:1337/send"

	payload := buildPayload(token, external_id, title, label_color, label_value, fields)

	p, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(p))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func buildPayload(flowToken string, threadID string, title string, statusColor string, statusValue string, fields []Field) *Payload {
	fdData := &robification_fdThread{
		Flow_Token: flowToken,
		Event:      "activity",
		Author: Author{
			Name:   "Wallace",
			Avatar: "https://d2cxspbh1aoie1.cloudfront.net/avatars/local/87dee60e5d045f2fa8371a2f3f45b919a0e9cd0209e3689a46951411fef18681/120",
		},
		Title:              title,
		External_Thread_Id: threadID,
		Thread: Thread{
			Title:  title,
			Fields: fields,
			//Body:         "some optional messaging that can be added to the thread as a sub header",
			External_Url: "http://jrobles.net",
			Status: ThreadStatus{
				Color: statusColor,
				Value: statusValue,
			},
		},
	}

	payload := &Payload{
		Targets: []Target{
			Target{
				Destination_Type:     "flowdock",
				Destination_Sub_Type: "new_thread",
				Data:                 *fdData,
			},
		},
	}
	return payload
}
