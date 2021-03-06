package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

/*
	APIサーバで受け取ったdatetimeフィールドの値をtime.Time型に変換する。
	datetimeフィールドはRFC3339のフォーマットで受信することを想定する。
*/
func UnmarshalDateTime(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		return time.Parse(time.RFC3339, v)
	case time.Time:
		return v, nil
	default:
		return time.Now(), fmt.Errorf("DateTime is invalid")
	}
}

/*
	APIサーバからJSONを返す際に、time.Time型をstringに変換する。
*/
func MarshalDateTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(strconv.Quote(t.Format(time.RFC3339))))
	})
}
