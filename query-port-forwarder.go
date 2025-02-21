package query2port

import (
	"context"
	"net"
	"net/http"
)

type Config struct {
	QueryParamName string `yaml:"queryParamName"`
}

func CreateConfig() *Config {
	return &Config{}
}

type QueryPortForwarder struct {
	next           http.Handler
	name           string
	queryParamName string
}

func New(_ context.Context, next http.Handler, c *Config, name string) (http.Handler, error) {
	return &QueryPortForwarder{
		next:           next,
		name:           name,
		queryParamName: c.QueryParamName,
	}, nil
}

func (qpf *QueryPortForwarder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	port := r.URL.Query().Get(qpf.queryParamName)
	if len(port) == 0 || len(port) > 5 {
		qpf.next.ServeHTTP(w, r)
		return
	}
	r.URL.Host = net.JoinHostPort(r.URL.Hostname(), port)
	qpf.next.ServeHTTP(w, r)
}
