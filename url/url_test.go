package url

import "testing"

func TestUrlIsInternal(t *testing.T) {
	tests := []struct {
		want bool
		url  string
		host string
	}{
		{true, "http://a.com", "a.com"},
		{true, "http://a.com", "a.com:80"},
		{true, "http://a.com:80", "a.com"},
		{true, "http://a.com:80", "a.com:80"},
		{true, "http://a.com/asdf", "a.com"},
		{true, "https://a.com", "a.com"},
		{true, "https://a.com/grh", "http://a.com/adge"},
		{true, "https://a.com:80/grh", "http://a.com/adge"},
		{true, "https://a.com/grh", "http://a.com:80/adge"},
		{true, "/", "a.com"},
		{true, "/asf", "a.com"},

		{false, "http://b.com", "a.com"},
		{false, "http://a.com:8080", "a.com"},
		{false, "http://a.com", "a.com:8080"},
	}

	for _, tt := range tests {
		if tt.want != UrlIsInternal(tt.url, tt.host) {
			t.Errorf("Url %s, host %s; expected %t but got %t", tt.url, tt.host, tt.want, !tt.want)
		}
	}
}
