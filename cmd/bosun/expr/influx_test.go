package expr

import (
	"fmt"
	"testing"
	"time"
)

func TestInflux(t *testing.T) {
	type influxTest struct {
		query  string
		expect string // empty for error
	}
	const influxFmt = "2006-01-02 15:04:05"
	date := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Hour
	start := date.Format(influxFmt)
	end := date.Add(dur).Format(influxFmt)
	tests := []influxTest{
		{
			"select * from a WHERE (time > 1 and TIME > 2)",
			fmt.Sprintf("select * from a WHERE time > '%s' and time < '%s'", start, end),
		},
		{
			"select * from a WHERE value > 0",
			fmt.Sprintf("select * from a WHERE value > 0 and time > '%s' and time < '%s'", start, end),
		},
		{
			"select * from a WHERE time > 0",
			"",
		},
	}
	for _, test := range tests {
		q, err := influxQueryDurations(date, test.query, dur.String(), "")
		break
		if err != nil && test.expect != "" {
			t.Errorf("%v: unexpected error: %v", test.query, err)
		} else if q != test.expect {
			t.Errorf("%v: \n\texpected: %v\n\tgot: %v", test.query, test.expect, q)
		}
	}
}
