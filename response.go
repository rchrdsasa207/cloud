package main

import (
	"encoding/json"
	"log"
)

type Response interface {
	Byte() []byte
	Status() int
}

type BasicRespose struct {
	status int
	Bytes  []byte
}

func (r *BasicRespose) Byte() []byte {
	return r.Bytes
}

func (r *BasicRespose) Status() int {
	return r.status
}

func NewResponse(status int, message string) *BasicRespose {
	j, err := json.Marshal(struct {
		Status  int    `json:"Status"`
		Message string `json:"Message"`
	}{status, message})
	if err != nil {
		log.Fatal("Cannot marshal the response: ", err)
	}
	return &BasicRespose{status: status, Bytes: j}
}
