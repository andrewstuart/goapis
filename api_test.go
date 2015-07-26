package apis

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		q := r.URL.Query()

		if q["test"][0] != "test" {
			t.Errorf("Test header was not present")
		}

		fmt.Fprintln(w, "")
	}))

	defer ts.Close()

	c := NewClient(ts.URL)

	c.DefaultParams(Query{
		"test": "test",
	})

	c.Get("", nil)
}
