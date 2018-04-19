package params

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

type data struct {
	X   string `http:"x"`
	Zip int    `http:"z;zip"`
}

func TestValidate(t *testing.T) {
	ts := []struct {
		name  string
		value string
		want  bool
	}{
		{name: "zip", value: "0000000", want: true},
		{name: "zip", value: "00000000", want: false},
	}

	for _, tc := range ts {
		got := validate(tc.name, tc.value)
		if got != tc.want {
			t.Errorf("fail to validate: got %t, want %t", got, tc.want)
		}
	}
}

func TestUnpack(t *testing.T) {
	ts := []struct {
		url        string
		raiseError bool
		want       data
	}{
		{url: "http://localhost/search?x=hoge&z=0000000", raiseError: false, want: data{"hoge", 0000000}},
		{url: "http://localhost/search?x=hoge&z=00000000", raiseError: true, want: data{"", 0}},
	}

	for _, tc := range ts {
		url, _ := url.Parse(tc.url)
		var req http.Request
		req.URL = url

		var got data

		if err := Unpack(&req, &got); err != nil {
			if tc.raiseError == true {
				continue
			}

			t.Errorf("fail to validate: %v", err)
			continue
		}

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%q: got %v, want %v", tc.url, got, tc.want)
		}
	}
}
