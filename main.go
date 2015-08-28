package robification

import (
	"bytes"
	"encoding/json"
	"net/http"
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

func Send(token string, external_id string, title string, message string, label_color string, label_value string, fields []Field) {

	url := "http://jrobles.net:1337/send"

	payload := buildPayload(token, external_id, title, message,label_color, label_value, fields)

	p, err := json.Marshal(payload)
	if err != nil {
		panic(err)
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

func buildPayload(flowToken string, threadID string, title string, message string, statusColor string, statusValue string, fields []Field) *Payload {
	fdData := &robification_fdThread{
		Flow_Token: flowToken,
		Event:      "activity",
		Author: Author{
			Name:   "robiBot",
			Avatar: "http://img3.wikia.nocookie.net/__cb20150501175408/villains/images/b/be/Early.PNG",
		},
		Title:              message,
		External_Thread_Id: threadID,
		Thread: Thread{
			Title:  title,
			Fields: fields,
			//Body:         "some optional messaging that can be added to the thread as a sub header",
			External_Url: "https://github.com/josemrobles/robification-go",
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
