package main

import (
	"encoding/json"
	"log"
)

type User struct {
}

func (*User) GetCreateAction() DefinedAction {
	return &CreateUser{}
}

type CreateUser struct {
}

func (u *CreateUser) GetFromJSON(raw []byte) {
	err := json.Unmarshal(raw, &u)
	if err != nil {
		log.Println("Could not umarshal data: ", err)
	}
}

func (u *CreateUser) Process(db *DB) Response {

	return nil
}
