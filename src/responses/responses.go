package responses

//Responding to Pong Requests
type Response_pong struct{
	Result string `json:"result"`
}

//Responding to Pong Requests with an error
type Response_error struct{
	Error string `json:"error"`
}

//Assinging Variable Responses that do not need to be changed.
var Responded_pong = Response_pong{"pong"}