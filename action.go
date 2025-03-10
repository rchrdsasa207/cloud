package main

import (
	"encoding/json"
	"log"
)

type BasicAction struct {
	Object string `json:"Object"`
	Action string `json:"Action"`
}

type Object interface {
	GetCreateAction() DefinedAction
}

type DefinedAction interface {
	Process(*DB) Response
	GetFromJSON([]byte)
}

func (db *DB) ServePOST(raw []byte, Authorization string) Response {
	req := &BasicAction{}
	err := json.Unmarshal(raw, &req)
	if err != nil {
		log.Println("Could not decode json from the client: ", err)
		return NewResponse(400, "Could not decode json: "+err.Error())
	}
	switch req.Object {
	case "User":
		break
	default:
		log.Println("Object not found: ", req.Object)
		return NewResponse(400, "Unexpected object provided")
	}
	switch req.Action {
	case "Create":
		break
	default:
		log.Println("Action not found: ", req.Object)
		return NewResponse(400, "Unexpected action provided")
	}
	return NewResponse(500, "Unexpected error")
}
