package robification

import (
	//"bytes"
	"encoding/json"
	//"net/http"
)

func Send(token string, external_id string, title string, message string, label_color string, label_value string, fields []Field) {
	/*
		url := "http://jrobles.net:1337/send"

		p := buildPayload()

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(p))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	*/
}

func buildPayload(fdData *fdDetailedThread) []byte {

	payload := &Payload{
		Targets: []Target{
			Target{
				Destination_Type:     "flowdock",
				Destination_Sub_Type: "new_thread",
				//Data:                 *fdData,
			},
		},
	}

	p, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	return p
}

func NewEmail(from string, subject string, body string, recipients []string) *email {
	email := &email{
		From:       from,
		Subject:    subject,
		Body:       body,
		Recipients: recipients,
	}
	return email
}

func NewFdChat(flowToken string, content string, externalUserName string) *fdChat {
	chat := &fdChat{
		Flow_Token:         flowToken,
		Content:            content,
		External_User_Name: externalUserName,
	}
	return chat
}

func NewSMS(from string, message string, recipients []string) *sms {
	sms := &sms{
		From:       from,
		Message:    message,
		Recipients: recipients,
	}
	return sms
}

func NewFdBasicThread(flowToken string, subject string, fromAddress string, source string, content string) *fdBasicThread {
	fdBasicThread := &fdBasicThread{
		Flow_Token:   flowToken,
		Subject:      subject,
		From_Address: fromAddress,
		Source:       source,
		Content:      content,
	}
	return fdBasicThread
}

func NewFdDetailedThread(flowToken string, threadID string, title string, message string, statusColor string, statusValue string, fields []Field) *fdDetailedThread {
	fdDetailedThread := &fdDetailedThread{
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
	return fdDetailedThread
}
