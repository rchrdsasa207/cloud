package main

type Response interface {
	Byte() []byte
	Status() int
}
