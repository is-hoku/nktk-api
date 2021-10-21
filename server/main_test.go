package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/is-hoku/nktk-api/server/model"
	"github.com/is-hoku/nktk-api/server/router"
)

func TestGetByID(t *testing.T) {
	router := router.NewRouter()

	t.Run("is_200", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/1", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		if http.StatusOK != recorder.Code {
			t.Error("A status code is not 200.")
		}
		body := recorder.Body.String()
		var item model.Item
		err := json.Unmarshal([]byte(body), &item)
		if err != nil {
			t.Fatal(err)
		}
		if item.ID == 0 || item.Text == "" || item.Class == "" {
			t.Error("A body is not included.")
		}
	})

	t.Run("is_404", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/0", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		if http.StatusNotFound != recorder.Code {
			t.Error("A status code is not 404.")
		}
	})
}

func TestRandom(t *testing.T) {
	router := router.NewRouter()

	t.Run("is_200", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/random", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		if http.StatusOK != recorder.Code {
			t.Error("A status code is not 200.")
		}
		body := recorder.Body.String()
		var item model.Item
		err := json.Unmarshal([]byte(body), &item)
		if err != nil {
			t.Fatal(err)
		}
		if item.ID == 0 || item.Text == "" || item.Class == "" {
			t.Error("A body is not included.")
		}
	})

	t.Run("is_404", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/NotFoundパス", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		if http.StatusNotFound != recorder.Code {
			t.Error("A status code is not 404.")
		}
	})
}
