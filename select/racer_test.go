package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))

		fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(25 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 19*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}
