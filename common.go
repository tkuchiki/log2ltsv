package log2ltsv

import (
	"fmt"
	"strings"
	"time"
)

var iso8601 = "2006-01-02T15:04:05-07:00"

func getMethodAndURI(req string) (string, string) {
	values := strings.Fields(req)
	if len(values) < 2 {
		return "", ""
	}

	return values[0], values[1]
}

func makeValue(label, value string) string {
	return fmt.Sprintf("%s:%s", label, value)
}

func toLtsv(values []string) string {
	return strings.Join(values, "\t")
}

func stringToISO8601(value string) (string, error) {
	t, err := time.Parse("02/Jan/2006:15:04:05 -0700", value)
	if err != nil {
		return "", err
	}

	return t.Format(iso8601), nil
}

func timeToISO8601(t time.Time) (string, error) {
	return t.Format(iso8601), nil
}

func uint64ToString(i uint64) string {
	return fmt.Sprintf("%d", i)
}
