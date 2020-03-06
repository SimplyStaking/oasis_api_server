package responses

//Responding to Pong Requests
type Response_pong struct{
	Result string `json:"result"`
}

//Responding to Pong Requests with an error
type Response_error struct{
	Error string `json:"error"`
}

type Response_Conns struct{
	Results []string `json:"result"`
}

//Assinging Variable Responses that do not need to be changed.
var Responded_pong = Response_pong{"pong"}