package utils

type Response struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Meta       Meta        `json:"meta"`
	Errors     interface{} `json:"errors"`
	StatusCode int         `json:"status_code"`
}

type Meta struct {
	Trace Trace `json:"trace"`
}

type Trace struct {
	TraceID   string `json:"trace_id"`
	RequestID string `json:"request_id"`
}
