package main

import "testing"

func TestGreeting(t *testing.T) {
	resultado := greeting("teste")
	if resultado != "<b>teste</b>" {
		t.Error("O resultado deve ser <b>teste</b>", resultado)
	}

}
