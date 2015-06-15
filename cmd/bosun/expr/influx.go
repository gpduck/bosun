package expr

import (
	"fmt"
	"net/url"
	"time"

	"bosun.org/_third_party/github.com/MiniProfiler/go/miniprofiler"
	"bosun.org/cmd/bosun/expr/parse"
	"github.com/influxdb/influxdb/client"
	"github.com/influxdb/influxdb/influxql"
)

// Influx is a map of functions to query InfluxDB.
var Influx = map[string]parse.Func{
	"influx": {
		Args:   []parse.FuncType{parse.TypeString, parse.TypeString, parse.TypeString, parse.TypeString},
		Return: parse.TypeSeries,
		Tags:   influxTag,
		F:      InfluxQuery,
	},
}

func influxTag(args []parse.Node) (parse.Tags, error) {
	return nil, nil
}

func InfluxQuery(e *State, T miniprofiler.Timer, query, startDuration, endDuration, tagFormat string) (r *Results, err error) {
	s, err := timeInfluxRequest(e, T, query, startDuration, endDuration)
	if err != nil {
		return nil, err
	}
	_ = s
	return nil, fmt.Errorf("not done")
}

// influxQueryDuration adds time WHERE clauses to query for the given start and end durations.
func influxQueryDurations(now time.Time, query, start, end string) (string, error) {
	st, err := influxql.ParseStatement(query)
	if err != nil {
		return "", err
	}
	s, ok := st.(*influxql.SelectStatement)
	if !ok {
		return "", fmt.Errorf("influx: expected select statement")
	}
	check := func(
	influxql.WalkFunc(s.Condition, func(n influxql.Node) {
		b, ok := n.(*influxql.BinaryExpr)
		if !ok {
			return
		}
		if lv, ok := b.LHS.(*influxql.VarRef); ok{
			fmt.Printf("@%v.\n", lv.Val)
		}
	})
	return "", nil
}

func timeInfluxRequest(e *State, T miniprofiler.Timer, query, startDuration, endDuration string) (s *client.Response, err error) {
	conf := client.Config{
		URL: url.URL{
			Scheme: "http",
			Host:   e.InfluxHost,
		},
		Timeout: time.Minute,
	}
	conn, err := client.NewClient(conf)
	_ = conn
	if err != nil {
		return nil, err
	}
	/*
		T.StepCustomTiming("influx", "query", query, func() {
			getFn := func() (interface{}, error) {
				return e.tsdbContext.Query(req)
			}
			var val interface{}
			val, err = e.cache.Get(string(b), getFn)
			s = val.(opentsdb.ResponseSet).Copy()
		})
	*/
	return
}
