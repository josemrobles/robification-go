package robification

type fdChat struct {
	Flow_Token string
	Content    string
}

type apiResponse struct {
	Items []struct {
		Name string `json:Name`
	} `json:Items`
}
