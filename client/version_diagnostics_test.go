package client

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServerVersionRetainsLegacyDiagnostics(t *testing.T) {
	for _, components := range []string{"", `,"Components":[{"Name":"Engine","Version":"29.0","Details":{"GitCommit":"component-commit"}}]`} {
		t.Run(components, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/v1.24/version" {
					t.Errorf("unexpected request: %s", r.URL.Path)
					w.WriteHeader(404)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"Version":"1.12.6","GitCommit":"top-commit","GoVersion":"legacy-go","KernelVersion":"legacy-kernel","BuildTime":"legacy-time"` + components + `}`))
			}))
			defer server.Close()
			cli, err := New(WithHost("tcp://"+strings.TrimPrefix(server.URL, "http://")), WithAPIVersion("1.24"))
			if err != nil {
				t.Fatal(err)
			}
			defer cli.Close()
			got, err := cli.ServerVersion(t.Context(), ServerVersionOptions{})
			if err != nil {
				t.Fatal(err)
			}
			if got.GitCommit != "top-commit" || got.GoVersion != "legacy-go" || got.KernelVersion != "legacy-kernel" || got.BuildTime != "legacy-time" {
				t.Fatalf("lost top-level diagnostics: %+v", got)
			}
		})
	}
}
