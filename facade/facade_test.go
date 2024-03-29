package facade

import (
	"testing"
)

var expect = "A module running\nB module running"

func TestFacadeAPI(t *testing.T) {
	api := NewAPI()
	ret := api.Test()
	if ret != expect{
		t.Fatalf("expect \n%s,\n return %s", expect, ret)
	}
}
