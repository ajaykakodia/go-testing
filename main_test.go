//go:build integration
// +build integration

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HomePage)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

// func TestFilterUnique(t *testing.T) {
// 	input := []main.Developer{
// 		{Name: "Ajay"},
// 		{Name: "Vijay"},
// 		{Name: "Salman"},
// 		{Name: "Ranvijay"},
// 		{Name: "Raju"},
// 		{Name: "Vijay"},
// 		{Name: "Ranjeet"},
// 		{Name: "Ajay"},
// 	}

// 	expected := []string{
// 		"Ajay",
// 		"Vijay",
// 		"Salman",
// 		"Ranvijay",
// 		"Raju",
// 		"Ranjeet",
// 	}

// 	assert.Equal(t, expected, main.FilterUnique(input))
// 	// result := FilterUnique(input)
// 	// fmt.Println(result)
// 	// if !reflect.DeepEqual(result, expected) {
// 	// 	t.Fail()
// 	// }
// }
