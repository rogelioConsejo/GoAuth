package auth

import "fmt"

//HttpVerb
type HttpVerb string

const (
	GET    HttpVerb = "GET"
	POST   HttpVerb = "POST"
	PUT    HttpVerb = "PUT"
	DELETE HttpVerb = "DELETE"
)

func (verb *HttpVerb)String()string{
	return fmt.Sprintf("%s", *verb)
}

//uri
type uri string

//requestBody
type requestBody string
