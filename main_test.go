package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getResponse(url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func TestMainHandlerWhenOk(t *testing.T) {
	response := getResponse("/cafe?count=7&city=moscow")

	assert.Equal(t, http.StatusOK, response.Code)
	require.NotEmpty(t, response.Body)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	response := getResponse("/cafe?count=7&city=omsk")

	assert.Equal(t, http.StatusBadRequest, response.Code)
	require.NotEmpty(t, response.Body)
	require.Equal(t, response.Body.String(), "wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	response := getResponse("/cafe?count=7&city=moscow")

	assert.Equal(t, http.StatusOK, response.Code)

	responseBody := response.Body.String()
	sliceCafe := strings.Split(responseBody, ",")

	require.Len(t, sliceCafe, totalCount)
}
