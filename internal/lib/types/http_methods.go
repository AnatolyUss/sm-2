package types

type HttpMethod string

type HttpMethodEnum struct {
	CONNECT HttpMethod
	DELETE  HttpMethod
	GET     HttpMethod
	HEAD    HttpMethod
	OPTIONS HttpMethod
	PATCH   HttpMethod
	POST    HttpMethod
	PUT     HttpMethod
	TRACE   HttpMethod
}

func InitializeHttpMethod() HttpMethodEnum {
	return HttpMethodEnum{
		CONNECT: "CONNECT",
		DELETE:  "DELETE",
		GET:     "GET",
		HEAD:    "HEAD",
		OPTIONS: "OPTIONS",
		PATCH:   "PATCH",
		POST:    "POST",
		PUT:     "PUT",
		TRACE:   "TRACE",
	}
}
