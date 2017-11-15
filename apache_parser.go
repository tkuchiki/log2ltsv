package log2ltsv

import (
	"fmt"
	"regexp"
)

type ApacheLog struct {
	remoteHost  string
	user        string
	time        string
	method      string
	uri         string
	protocol    string
	statusCode  string
	bytes       string
	referer     string
	userAgent   string
	requestTime string
}

type ApacheParser struct {
	log *ApacheLog
}

func NewApacheParser() *ApacheParser {
	return &ApacheParser{}
}

func (ap *ApacheParser) parse(line string) error {
	apacheLogRE := regexp.MustCompile(`^(\S+)\s` + // remote host
		`\S+\s+` +
		`(\S+\s+)+` + // user
		`\[([^]]+)\]\s` + // time
		`"(\S*)\s?` + // method
		`((?:[^"]*(?:\\")?)*)\s` + // URL
		`([^"]*)"\s` + // protocol
		`(\S+)\s` + // status code
		`(\S+)\s` + // bytes
		`"((?:[^"]*(?:\\")?)*)"\s` + // referer
		`"(.*)"` + // user agent
		`\s(.*)$`) // request_time

	group := apacheLogRE.FindStringSubmatch(line)

	if len(group) < 1 {
		return fmt.Errorf("invalid log format")
	}

	t, err := stringToISO8601(group[3])
	if err != nil {
		return err
	}

	ap.log = &ApacheLog{
		remoteHost:  group[1],
		user:        group[2],
		time:        t,
		method:      group[4],
		uri:         group[5],
		protocol:    group[6],
		statusCode:  group[7],
		bytes:       group[8],
		referer:     group[9],
		userAgent:   group[10],
		requestTime: group[11],
	}

	return nil
}

func (ap *ApacheParser) ParseAndOutput(line string) (string, error) {
	err := ap.parse(line)
	if err != nil {
		return "", err
	}

	values := make([]string, 0)
	values = append(values, makeValue("user", ap.log.user))
	values = append(values, makeValue("time", ap.log.time))
	values = append(values, makeValue("method", ap.log.method))
	values = append(values, makeValue("uri", ap.log.uri))
	values = append(values, makeValue("status", ap.log.statusCode))
	values = append(values, makeValue("size", ap.log.bytes))
	values = append(values, makeValue("reqtime", ap.log.requestTime))
	values = append(values, makeValue("apptime", ap.log.requestTime))

	return toLtsv(values), nil
}
