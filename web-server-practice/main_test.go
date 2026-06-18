package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHome(t *testing.T) {
	w := httptest.NewRecorder()
	handleHome(w, nil)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code, expected %v but got %v\n body: %s\n", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("ASCII WEB SOLVER\n")

	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return got: %q, expected: %q", w.Body.String(), expectedMessage)
	}

}

func TestHandleHelp(t *testing.T) {
	w := httptest.NewRecorder()
	handleHelp(w, nil)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code, expected %v but got %v\n body: %s\n", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("ASK QUESTIONS\n")

	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected: %q", w.Body.String(), expectedMessage)
	}

}

func TestHandleHelloParameterized(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?user=TestMan", nil)
	w := httptest.NewRecorder()

	handleHelloParameterized(w, req)

	desiredCode := http.StatusOK

	if w.Code != desiredCode {
		t.Errorf("bad response code, expected %v but got %v\n body: %s\n", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello, TestMan!\n")

	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleHelloParameterizedNoParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello/", nil)

	w := httptest.NewRecorder()

	handleHelloParameterized(w, req)

	desiredCode := http.StatusOK

	if w.Code != desiredCode {
		t.Errorf("bad response code, expected %v but got %v\n body: %s\n", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello, User!\n")

	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleHelloParameterizedWrongParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?bull=crap", nil)

	w := httptest.NewRecorder()

	handleHelloParameterized(w, req)

	desiredCode := http.StatusOK

	if w.Code != desiredCode {
		t.Errorf("bad response code, expected %v but got %v\n body: %s\n", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello, User!\n")

	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}
