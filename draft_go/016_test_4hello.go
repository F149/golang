package main

import "testing"

func TestName(t *testing.T) {
	name := GetName()
	if name != "Golang" {
		t.Error("Respone from GetName is unexped value")
	}
}
