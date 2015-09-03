package robification

type fdBasicThread struct {
	Flow_Token   string
	Subject      string
	From_Address string
	Source       string
	Content      string
}

type fdChat struct {
	Flow_Token         string
	Content            string
	External_User_Name string
}

type email struct {
	From       string
	Subject    string
	Body       string
	Recipients []string
}

type sms struct {
	From       string
	Message    string
	Recipients []string
}

type Payload struct {
	Targets []Target
}

type Target struct {
	Destination_Type     string
	Destination_Sub_Type string
	Data                 fdDetailedThread
}

type fdDetailedThread struct {
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
