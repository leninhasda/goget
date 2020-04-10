package goget

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetMd5Hash(t *testing.T) {
	getter := New(http.DefaultClient)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	md5, err := getter.GetMd5Hash(ts.URL)
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	if md5 != "d32da7152864375f6fb7fce78877a98c" {
		t.Errorf("test failed: expected d32da7152864375f6fb7fce78877a98c, got %s", md5)
	}
}
