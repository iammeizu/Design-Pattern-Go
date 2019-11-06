package adapter

import "testing"

func TestNewAdaptee(t *testing.T) {
	adaptee := NewAdaptee()
	s := adaptee.SpecificRequest()
	if s != "adaptee method" {
		t.Fatal("test failed")
	}
}

func TestNewAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	adapter := NewAdapter(adaptee)
	s := adapter.Request()
	if s != "adaptee method" {
		t.Fatal("test failed")
	}
}
