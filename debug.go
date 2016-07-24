package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

type DebugTransport struct {
	Base   http.RoundTripper
	Output *os.File
}

func NewDebugTransport(base http.RoundTripper, output *os.File) http.RoundTripper {
	return &DebugTransport{
		Base:   base,
		Output: output,
	}
}

func (t *DebugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqDump, _ := httputil.DumpRequest(req, false)
	fmt.Fprintln(t.Output, string(reqDump))
	resp, err := t.Base.RoundTrip(req)
	respDump, _ := httputil.DumpResponse(resp, false)
	fmt.Fprintln(t.Output, string(respDump))
	return resp, err
}
