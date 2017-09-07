package log2ltsv

import (
	"github.com/Songmu/axslogparser"
)

type ApacheParser struct {
	parser axslogparser.Apache
}

func NewApacheParser() *ApacheParser {
	return &ApacheParser{
		parser: axslogparser.Apache{},
	}
}

func (ap *ApacheParser) ParseAndOutput(line string) (string, error) {
	l, err := ap.parser.Parse(line)
	if err != nil {
		return "", err
	}

	values := make([]string, 0)

	t, err := timeToISO8601(l.Time)
	if err != nil {
		return "", err
	}
	values = append(values, makeValue("time", t))

	values = append(values, makeValue("user", l.User))
	values = append(values, makeValue("status", uint64ToString(uint64(l.Status))))
	values = append(values, makeValue("size", uint64ToString(l.Size)))

	method, uri := getMethodAndURI(l.Request)
	values = append(values, makeValue("method", method))
	values = append(values, makeValue("uri", uri))
	values = append(values, makeValue("apptime", "0"))

	return toLtsv(values), nil

	return "", nil
}
