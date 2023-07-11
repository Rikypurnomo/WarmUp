package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"unicode"

	models_server "github.com/Rikypurnomo/warmup/server/models"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func ToSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}

		out = append(out, unicode.ToLower(runes[i]))
	}

	return strings.Replace(string(out), "-", "", -1)
}

func GetTraceIdFromContext(ctx context.Context) string {
	if span := trace.SpanFromContext(ctx); span != nil {
		return span.SpanContext().TraceID().String()
	}

	return ""
}

func SetAttributeReqRes(r *http.Request, res string) {
	span := trace.SpanFromContext(r.Context())
	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	mars, _ := json.Marshal(r.Header)
	// restore body
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	if span != nil {
		span.SetAttributes(attribute.KeyValue{
			Key:   attribute.Key("request"),
			Value: attribute.StringValue(string(body)),
		})
		span.SetAttributes(attribute.KeyValue{
			Key:   attribute.Key("response"),
			Value: attribute.StringValue(res),
		})
		span.SetAttributes(attribute.KeyValue{
			Key:   attribute.Key("request_id"),
			Value: attribute.StringValue(GetTraceIdFromContext(r.Context())),
		})
		span.SetAttributes(attribute.KeyValue{
			Key:   attribute.Key("response_header"),
			Value: attribute.StringValue(string(mars)),
		})
	}
}

func ResPagination(count int64, page int64, limit int64) models_server.MetaPagination {
	var metaPagination models_server.MetaPagination
	metaPagination.CurrentPage = page
	metaPagination.TotalCount = count
	metaPagination.NextPage = page + 1
	metaPagination.PrevPage = page - 1
	metaPagination.TotalPages = (count + limit - 1) / limit
	return metaPagination
}
