package log2ltsv

import (
	"github.com/satyrius/gonx"
)

type NginxParser struct {
	parser *gonx.Parser
}

var nginxLogFormat = `$remote_addr - $remote_user [$time_local] "$request" $status $body_bytes_sent "$http_referer" "$http_user_agent" "$http_x_forwarded_for" $request_time $upstream_response_time`
var nginxVars = []string{"remote_user", "time_local", "request", "status", "body_bytes_sent", "request_time", "upstream_response_time"}
var nginxVarsTable = map[string]string{
	"remote_user":            "user",
	"time_local":             "time",
	"status":                 "status",
	"body_bytes_sent":        "size",
	"request_time":           "reqtime",
	"upstream_response_time": "apptime",
}

func NewNginxParser() Parser {
	return &NginxParser{
		parser: gonx.NewParser(nginxLogFormat),
	}
}

func (np *NginxParser) ParseAndOutput(data string) (string, error) {
	entry, err := np.parser.ParseString(data)
	if err != nil {
		return "", err
	}

	values := make([]string, 0)

	for _, v := range nginxVars {
		var value string
		var err error

		switch v {
		case "time_local":
			value, err = entry.Field("time_local")
			if err != nil {
				return "", err
			}
			t, err := stringToISO8601(value)
			if err != nil {
				return "", err
			}
			values = append(values, makeValue("time", t))
		case "request":
			req, err := entry.Field("request")
			if err != nil {
				return "", err
			}
			method, uri := getMethodAndURI(req)
			values = append(values, makeValue("method", method))
			values = append(values, makeValue("uri", uri))
		default:
			value, err = entry.Field(v)
			if err != nil {
				return "", err
			}
			label := nginxVarsTable[v]
			values = append(values, makeValue(label, value))
		}
	}

	return toLtsv(values), nil
}
