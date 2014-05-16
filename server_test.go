package main

import (
	"bytes"
	"fmt"
	"github.com/bmizerany/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestSuccessfulGetHome(t *testing.T) {
	server := NewServer(80)

	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		panic(err)
	}
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, `You want weapons? We’re in a library! Books! The best weapons in the world!`, fmt.Sprintf("%s", response.Body))
}

func TestPostMetricWithoutData(t *testing.T) {
	server := NewServer(80)

	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	request, err := http.NewRequest("POST", "/collect", nil)
	if err != nil {
		panic(err)
	}
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, `1`, fmt.Sprintf("%s", response.Body))
}

func TestSuccessfulPostMetric(t *testing.T) {
	server := NewServer(80)

	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	payload := url.Values{
		"metric[key]":   {"collect.test.counter"},
		"metric[value]": {"10"},
		"metric[kind]":  {"counter"},
	}
	content := bytes.NewBufferString(payload.Encode())
	request, err := http.NewRequest("POST", "/collect", content)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(payload.Encode())))
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, `1`, fmt.Sprintf("%s", response.Body))
}
