package controller

import (
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"testing"
)

func TestServer_Start(t *testing.T) {
	type req struct {
		method, path string
		headers      map[string]string
		query, body  string
	}
	type resp struct {
		status int
		body   string
	}
	tests := []struct {
		name string
		req  req
		want resp
	}{
		{
			name: "GetHello",
			req: req{
				method:  "GET",
				path:    "/hello",
				headers: map[string]string{},
				query:   "?name=budougumi0617",
				body:    "",
			},
			want: resp{
				status: http.StatusOK,
				body:   "{\"name\":\"budougumi0617\"}\n",
			},
		},
		{
			name: "GetHelloWithOutQuery",
			req: req{
				method: "GET",
				path:   "/hello",
			},
			want: resp{
				status: http.StatusOK,
				body:   "{}\n",
			},
		},
		{
			name: "PostAPIRegister",
			req: req{
				method: "POST",
				path:   "/api/register",
				headers: map[string]string{
					"Content-Type": "application/json",
				},
				query: "?hoge_id=20",
				body:  `{"name": "budougumi0617", "lastName":"taro"}`,
			},
			want: resp{
				status: http.StatusOK,
			},
		},
		{
			name: "PostAPIRegister_NotEnoughQuery",
			req: req{
				method: "POST",
				path:   "/api/register",
				headers: map[string]string{
					"Content-Type": "application/json",
				},
				query: "",
				body:  `{"name":"budougumi", "="lastName":"taro"}`,
			},
			want: resp{
				status: http.StatusUnprocessableEntity,
				body:   `{"code":602,"message":"hoge_id in query is required"}`,
			},
		},
		{
			name: "PostAPIRegister_NotEnoughBody",
			req: req{
				method: "POST",
				path:   "/api/register",
				headers: map[string]string{
					"Content-Type": "application/json",
				},
				query: "?hoge_id=20",
				body:  `{"lastName":"taro"}`,
			},
			want: resp{
				status: http.StatusUnprocessableEntity,
				body:   `{"code":602,"message":"name in body is required"}`,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			l, err := net.Listen("tcp", "127.0.0.1:0")
			if err != nil {
				t.Fatal(err)
			}
			s := NewServer(l)
			t.Cleanup(func() {
				if err := s.Server.Close(); err != nil {
					t.Logf("Server.Server.Close() failed: %v", err)
				}
			})

			go func() {
				if err := s.Start(); err != http.ErrServerClosed {
					t.Logf("Server.Start() faled termination: %v", err)
				}
			}()
			url := "http://" + l.Addr().String() + tt.req.path + tt.req.query
			req, err := http.NewRequest(tt.req.method, url, strings.NewReader(tt.req.body))
			for k, v := range tt.req.headers {
				req.Header.Add(k, v)
			}

			if err != nil {
				t.Fatalf("create request failed %v", err)
			}
			got, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("http.Get failed: %v", err)
			}
			body, err := ioutil.ReadAll(got.Body)
			if err != nil {
				t.Fatalf("ioutil.ReadAll failed: %v", err)
			}
			if got.StatusCode != tt.want.status {
				t.Errorf("%q: want %d, but got %d", tt.name, tt.want.status, got.StatusCode)
			}

			if string(body) != tt.want.body {
				t.Errorf("%q: want %q, but got %q", tt.name, tt.want.body, body)
			}
		})
	}
}
