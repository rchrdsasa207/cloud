package main

type basicAction struct {
	Object string `json:"Object"`
	Action string `json:"Action"`
}

func (db *DB) ServePOST(raw []byte, Authorization string) []byte {
	return nil
}
