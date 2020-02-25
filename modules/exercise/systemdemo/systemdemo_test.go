package main

import "testing"

func TestEcho(t *testing.T) {
	if Echo("hello") != "hellosuffix" {
		t.Log("echo is wrong")
		t.Fail()
	}
}
