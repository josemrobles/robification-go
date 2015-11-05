package robification

type chatMessage struct {
	Url   string
	Token string
	Text  string
}

type apiResponse struct {
	Messages []struct {
		Status string `json:Status`
	} `json:Items`
}
