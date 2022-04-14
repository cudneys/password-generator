package models

type Response struct {
	Request  RequestParams `json:"request"`
	Password string        `json:"password"`
	Error    string        `json:"error,omitempty"`
}
