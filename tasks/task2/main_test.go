package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	recived := unpack(`a\`)
	expected := ``
	if recived != expected {
		t.Errorf("%v != %v", recived, expected)
	}

	recived = unpack(`34`)
	expected = ``
	if recived != expected {
		t.Errorf("%v != %v", recived, expected)
	}

	recived = unpack(`a4bc2d5e`)
	expected = `aaaabccddddde`
	if recived != expected {
		t.Errorf("%v != %v", recived, expected)
	}

	recived = unpack(`abcd`)
	expected = `abcd`
	if recived != expected {
		t.Errorf("%v != %v", recived, expected)
	}

	recived = unpack(`a11b`)
	expected = `aaaaaaaaaaab`
	if recived != expected {
		t.Errorf("%v != %v", recived, expected)
	}

	recived = unpack(``)
	expected = ``
	if recived != expected {
		t.Errorf("%v != %v", recived, expected)
	}

	recived = unpack(`qwe\`)
	expected = ``
	if recived != expected {
		t.Errorf("%v != %v", recived, expected)
	}

	recived = unpack(`qwe\\`)
	expected = `qwe\`
	if recived != expected {
		t.Errorf("%v != %v", recived, expected)
	}

	recived = unpack(`qwe\4\5`)
	expected = `qwe45`
	if recived != expected {
		t.Errorf("%v != %v", recived, expected)
	}

	recived = unpack(`qwe\45`)
	expected = `qwe44444`
	if recived != expected {
		t.Errorf("%v != %v", recived, expected)
	}
}
