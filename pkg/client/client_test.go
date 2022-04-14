package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type example struct {
	A int    `json:"a"`
	B string `json:"b"`
}

func Test_Auth(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	expectedUsername := "my-username"
	expectedPassword := "my-password"
	var actualUsername string
	var actualPassword string
	var hasBasicAuth bool

	mux.HandleFunc("/testing", func(w http.ResponseWriter, r *http.Request) {
		actualUsername, actualPassword, hasBasicAuth = r.BasicAuth()
	})

	client := New(expectedUsername, expectedPassword, OptionWithBaseURL(server.URL))
	err := client.Get(context.TODO(), "/testing", nil)
	require.NoError(t, err)

	assert.True(t, hasBasicAuth)
	assert.Equal(t, expectedUsername, actualUsername)
	assert.Equal(t, expectedPassword, actualPassword)
}

func Test_Get(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	expected := example{
		A: 123,
		B: "hello world!",
	}

	var actualMethod string

	mux.HandleFunc("/testing", func(w http.ResponseWriter, r *http.Request) {
		actualMethod = r.Method
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expected)
	})

	client := New("", "", OptionWithBaseURL(server.URL))
	var actual example
	err := client.Get(context.TODO(), "/testing", &actual)
	require.NoError(t, err)

	assert.Equal(t, http.MethodGet, actualMethod)
	assert.Equal(t, expected, actual)
}

func Test_Post(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	expectedRequest := example{
		A: 123,
		B: "hello world!",
	}
	expectedResponse := example{
		A: 123,
		B: "hello world!",
	}

	var actualRequest example
	var actualResponse example
	var actualMethod string

	mux.HandleFunc("/testing", func(w http.ResponseWriter, r *http.Request) {
		actualMethod = r.Method
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expectedResponse)
		_ = json.NewDecoder(r.Body).Decode(&actualRequest)
	})

	client := New("", "", OptionWithBaseURL(server.URL))
	err := client.Post(context.TODO(), "/testing", &actualResponse, expectedRequest)
	require.NoError(t, err)

	assert.Equal(t, http.MethodPost, actualMethod)
	assert.Equal(t, expectedRequest, actualRequest)
	assert.Equal(t, expectedResponse, actualResponse)
}

func Test_Patch(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	expectedRequest := example{
		A: 123,
		B: "hello world!",
	}
	expectedResponse := example{
		A: 123,
		B: "hello world!",
	}

	var actualRequest example
	var actualResponse example
	var actualMethod string

	mux.HandleFunc("/testing", func(w http.ResponseWriter, r *http.Request) {
		actualMethod = r.Method
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expectedResponse)
		_ = json.NewDecoder(r.Body).Decode(&actualRequest)
	})

	client := New("", "", OptionWithBaseURL(server.URL))
	err := client.Patch(context.TODO(), "/testing", &actualResponse, expectedRequest)
	require.NoError(t, err)

	assert.Equal(t, http.MethodPatch, actualMethod)
	assert.Equal(t, expectedRequest, actualRequest)
	assert.Equal(t, expectedResponse, actualResponse)
}

func Test_Delete(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	var actualMethod string
	mux.HandleFunc("/testing", func(w http.ResponseWriter, r *http.Request) {
		actualMethod = r.Method
	})

	client := New("", "", OptionWithBaseURL(server.URL))
	err := client.Delete(context.TODO(), "/testing")
	require.NoError(t, err)
	assert.Equal(t, http.MethodDelete, actualMethod)
}

func Test_Retries(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	expected := example{
		A: 123,
		B: "hello world!",
	}

	var count int
	mux.HandleFunc("/testing", func(w http.ResponseWriter, r *http.Request) {
		if count < 3 {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(expected)
		}
		count++
	})

	client := New("", "", OptionWithBaseURL(server.URL))
	var actual example
	err := client.Get(context.TODO(), "/testing", &actual)
	require.NoError(t, err)
	assert.Equal(t, 4, count)
	assert.Equal(t, expected, actual)
}
