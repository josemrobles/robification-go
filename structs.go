package robification

type chatMessage struct {
	Url   string
	Token string
	Text  string
}

type apiResponses struct {
	Responses []response `json:Items`
}

type response struct {
	Status string `json:Status`
}
