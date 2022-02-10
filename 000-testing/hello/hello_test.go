package hello

import "testing"

func TestDump(t *testing.T){
	want := "hello world"
	if got := Greeting(); got != want {
		t.Errorf("Greet() = %q, want %q",got,want)
	}
}