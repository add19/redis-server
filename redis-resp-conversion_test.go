package main

import (
	"testing"
)

func TestTestFunc(t *testing.T) {

	var j String = "hello world"
	got := Serialization(j)
	want := "+hello world\r\n"

	if got != want {
		t.Errorf("got : %q want : %q", got, want)
	}

}

func TestIntegerSer(t *testing.T) {
	var i Integer = 1
	got := Serialization(i)
	want := ":1\r\n"

	if got != want {
		t.Errorf("got : %q want : %q", got, want)
	}
}

func TestBulkString(t *testing.T) {
	var i BulkString
	var str String = "Some random bulk string"
	i.Value = &str
	got := Serialization(i)
	want := "$23\r\nSome random bulk string\r\n"

	if got != want {
		t.Errorf("got : %q want : %q", got, want)
	}
}

func TestEmptyBulkString(t *testing.T) {
	var i BulkString
	var str String = ""
	i.Value = &str
	got := Serialization(i)
	want := "$0\r\n\r\n"

	if got != want {
		t.Errorf("got : %q want : %q", got, want)
	}
}

func TestNilString(t *testing.T) {
	var i BulkString
	i.Value = nil
	got := Serialization(i)
	want := "$-1\r\n"

	if got != want {
		t.Errorf("got : %q want : %q", got, want)
	}
}

func TestIntegerArr(t *testing.T) {
	var k IntegerArray = []Integer{2, 3, 4}
	got := Serialization(k)
	want := "*3\r\n:2\r\n:3\r\n:4\r\n"

	if got != want {
		t.Errorf("got : %q want : %q", got, want)
	}
}

func TestStringArr(t *testing.T) {
	var l StringArray = []string{"hello", "world", "!!!"}
	got := Serialization(l)
	want := "*3\r\n$5\r\nhello\r\n$5\r\nworld\r\n$3\r\n!!!\r\n"

	if got != want {
		t.Errorf("got : %q want : %q", got, want)
	}
}

func TestSpacedStringArr(t *testing.T) {
	var l StringArray = []string{"hello world", "!!!"}
	got := Serialization(l)
	want := "*2\r\n$11\r\nhello world\r\n$3\r\n!!!\r\n"

	if got != want {
		t.Errorf("got : %q want : %q", got, want)
	}
}

func TestError(t *testing.T) {
	var r RedisError = RedisError{msg: "Error"}
	got := Serialization(r)
	want := "-Error\r\n"

	if got != want {
		t.Errorf("got : %q want : %q", got, want)
	}
}

func TestNilResponse(t *testing.T) {
	var n NilResponse
	got := Serialization(n)
	want := "*-1\r\n"

	if got != want {
		t.Errorf("got : %q want : %q", got, want)
	}
}

func TestDeserialization(t *testing.T) {
	got := Deserialization(":200\r\n")
	want := []string{"200"}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got : %q want : %q", got, want)
		}
	}

}

func TestDeserializationErrorMsg(t *testing.T) {
	got := Deserialization("-Error\r\n")
	want := []string{"Error"}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got : %q want : %q", got, want)
		}
	}

}

func TestDeserializationBulk(t *testing.T) {
	got := Deserialization("$11\r\nhello world\r\n")
	want := []string{"hello world"}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got : %q want : %q", got, want)
		}
	}

}

func TestDeserializationArray(t *testing.T) {
	got := Deserialization("*3\r\n$5\r\nhello\r\n$5\r\nworld\r\n$3\r\n!!!\r\n")
	want := []string{"hello", "world", "!!!"}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got : %q want : %q", got, want)
		}
	}
}

func TestDeserializationNumeric(t *testing.T) {
	got := Deserialization("*3\r\n:2\r\n:3\r\n:4\r\n")
	want := []string{"2", "3", "4"}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got : %q want : %q", got, want)
		}
	}
}

func TestDeserializationEmptyString(t *testing.T) {
	got := Deserialization("$0\r\n\r\n")
	want := []string{""}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got : %q want : %q", got, want)
		}
	}
}

func TestDeserializationNil(t *testing.T) {
	got := Deserialization("$-1\r\n")
	want := []string{}
	if len(want) != len(got) && len(got) != 0 {
		t.Errorf("got : %q want : %q", got, want)
	}
}
