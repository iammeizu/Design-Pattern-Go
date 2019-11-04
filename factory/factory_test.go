package factory

import (
	"log"
	"testing"
)

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("tom")
	if s != "Hi, tom" {
		log.Fatal("Type1 test failed")
	}
}


func TestType2(t *testing.T) {
	api := NewAPI(2)
	s :=  api.Say("tom")
	if s != "Hello, tom" {
		log.Fatal("Type2 test failed")
	}
}