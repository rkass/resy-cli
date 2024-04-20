package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"time"

	"moul.io/http2curl"

	"github.com/spf13/viper"
)

func getAuthHeaders() *http.Header {
	apiKey := viper.GetString("resy_api_key")
	authToken := viper.GetString("resy_auth_token")
	return &http.Header{
		"authorization":         {fmt.Sprintf(`ResyAPI api_key="%s"`, apiKey)},
		"x-resy-auth-token":     {authToken},
		"x-resy-universal-auth": {authToken},
	}
}

type Req struct {
	QueryParams map[string]string
	Body        []byte
}

type loggingTransport struct{}

func (s *loggingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	bytes, _ := httputil.DumpRequestOut(r, true)

	resp, err := http.DefaultTransport.RoundTrip(r)
	// err is returned after dumping the response

	respBytes, _ := httputil.DumpResponse(resp, true)
	bytes = append(bytes, respBytes...)

	fmt.Printf("%s\n", bytes)

	return resp, err
}


func template(method string, contentType string) func(string, *Req) ([]byte, int, error) {
	return func(url string, p *Req) ([]byte, int, error) {
		req, _ := http.NewRequest(method, url, bytes.NewReader(p.Body))
		req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
		req.Header.Add("origin", "https://resy.com")
		req.Header.Add("referrer", "https://resy.com")
		req.Header.Add("x-origin", "https://resy.com")
		req.Header.Add("cache-control", "no-cache")
		req.Header.Add("accept-encoding", "application/json")
		req.Header.Add("acept", "*/*")
		req.Header.Add("connection", "keep-alive")
		client := &http.Client{Timeout: 3 * time.Second}
		authHeaders := getAuthHeaders()
		if contentType != "" {
			req.Header.Add("content-type", contentType)
		}
		for key, val := range *authHeaders {
			req.Header.Add(key, val[0])
		}
		if p.QueryParams != nil {
			query := req.URL.Query()
			for key, val := range p.QueryParams {
				query.Add(key, val)
			}
			req.URL.RawQuery = query.Encode()
		}
		
		command, _ := http2curl.GetCurlCommand(req)
		fmt.Println(command)
		fmt.Println("\n\n\n");

		res, err := client.Do(req)

		if err != nil {
			return nil, 500, err
		}
		if res == nil {
			return nil, 0, nil
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, res.StatusCode, err
		}

		return body, res.StatusCode, nil
	}
}

func PostJSON(url string, p *Req) ([]byte, int, error) {
	return template(http.MethodPost, "application/json")(url, p)
}

func PostForm(url string, p *Req) ([]byte, int, error) {
	return template(http.MethodPost, "application/x-www-form-urlencoded")(url, p)
}

func Get(url string, p *Req) ([]byte, int, error) {
	return template(http.MethodGet, "")(url, p)
}
