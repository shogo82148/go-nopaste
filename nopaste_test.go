package nopaste

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {
	dir, err := ioutil.TempDir("", "nopaste")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	nopaste := New(&Config{
		Root:    "",
		DataDir: dir,
	})
	ts := httptest.NewServer(nopaste)
	defer ts.Close()

	client := ts.Client()

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	want := "hello"
	resp, err := client.PostForm(ts.URL, url.Values{
		"text": []string{want},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	loc := strings.TrimLeft(resp.Header.Get("Location"), "/")
	b, err := ioutil.ReadFile(filepath.Join(dir, loc) + ".txt")
	if err != nil {
		t.Fatal(err)
	}

	got := string(b)
	if got != want {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestWithBackslash(t *testing.T) {
	dir, err := ioutil.TempDir("", "nopaste")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	nopaste := New(&Config{
		Root:    "",
		DataDir: dir,
	})
	ts := httptest.NewServer(nopaste)
	defer ts.Close()

	client := ts.Client()

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	_, err = client.Get(`/..\..\..\evil`)
	if err == nil {
		t.Fatalf("backslash should not be accepted")
	}
}
