package robification

type fdChat struct {
	Flow_Token string
	Content    string
}

type slackChat struct {
	Token   string
	Content string
}

type apiResponse struct {
	Messages []struct {
		Status string `json:Status`
	} `json:Items`
}
