package log2ltsv

import (
	//"github.com/k0kubun/pp"
	"github.com/satyrius/gonx"
)

type NginxParser struct {
	parser *gonx.Parser
}

var nginxLogFormat = `$remote_addr - $remote_user [$time_local] "$request" $status $body_bytes_sent "$http_referer" "$http_user_agent" "$http_x_forwarded_for"`
var nginxVars = []string{"remote_user", "status", "body_bytes_sent"}
var nginxVarsTable = map[string]string{
	"remote_user":     "user",
	"time_local":      "time",
	"status":          "status",
	"body_bytes_sent": "size",
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
	value, err := entry.Field("time_local")
	if err != nil {
		return "", err
	}
	t, err := stringToISO8601(value)
	if err != nil {
		return "", err
	}
	values = append(values, makeValue("time", t))

	for _, v := range nginxVars {
		value, err := entry.Field(v)
		if err != nil {
			return "", err
		}
		label := nginxVarsTable[v]
		values = append(values, makeValue(label, value))
	}
	req, err := entry.Field("request")
	if err != nil {
		return "", err
	}
	method, uri := getMethodAndURI(req)
	values = append(values, makeValue("method", method))
	values = append(values, makeValue("uri", uri))
	values = append(values, makeValue("apptime", "0"))

	return toLtsv(values), nil
}
