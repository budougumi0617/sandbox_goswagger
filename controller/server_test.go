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
			// curl localhost:18080/hello?name=budougumi0617
			// {"name":"budougumi0617"}
			// curl -D - -X POST -H "Content-Type: application/json" localhost:18080/api/register?hoge_id=20 -d '{"name": "budougumi0617", "lastName":"taro"}'
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
			// curl -D - -X POST -H "Content-Type: application/json" localhost:18080/api/register?hoge_id=20 -d '{"name": "budougumi0617", "lastName":"taro"}'
			name: "GetHelloWithOutQuery",
			req: req{
				method:  "GET",
				path:    "/hello",
				headers: map[string]string{},
				query:   "?name=budougumi0",
				body:    "",
			},
			want: resp{
				status: http.StatusOK,
				body:   "{}\n",
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
				t.Logf("%q: Server.Server.Close() tamenated", tt.name)
			}()
			url := "http://" + l.Addr().String() + tt.req.path + tt.req.query
			req, err := http.NewRequest(tt.req.method, url, strings.NewReader(tt.req.body))
			if err != nil {
				t.Fatalf("create request failed %v", err)
			}
			cli := &http.Client{}
			got, err := cli.Do(req)
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
