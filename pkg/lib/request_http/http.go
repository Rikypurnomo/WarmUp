package request_http

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

func createNewRequest() *resty.Request {
	client := resty.New().
		SetRetryCount(5).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(5 * time.Second).
		SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
			return 0, errors.New("quota exceeded")
		}).
		// SetTimeout(5 * time.Second).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		EnableTrace()

	req := client.R()
	return req
}

// HTTPGet func
func HTTPGet(context context.Context, url string, header map[string]string) (int, []byte, resty.TraceInfo, error) {
	_, parentSpan := tracer.Start(context, "HttpRequest.HTTPGet")
	defer parentSpan.End()

	parentSpan.AddEvent(fmt.Sprintf("GetHttRequest %s", url))
	req := createNewRequest().SetHeaders(header)
	resp, err := req.Execute(http.MethodGet, url)
	trace := resp.Request.TraceInfo()

	if err != nil {
		parentSpan.RecordError(err)
		return resp.StatusCode(), nil, trace, err
	}
	return resp.StatusCode(), resp.Body(), trace, nil
}

// HTTPPost func
func HTTPPost(context context.Context, url string, jsondata interface{}, header map[string]string) (int, []byte, resty.TraceInfo, error) {
	_, parentSpan := tracer.Start(context, "HttpRequest.HTTPPost")
	defer parentSpan.End()

	parentSpan.AddEvent(fmt.Sprintf("GetHttRequest %s", url))
	req := createNewRequest().SetHeaders(header).SetBody(jsondata)
	resp, err := req.Execute(http.MethodPost, url)
	trace := resp.Request.TraceInfo()

	if err != nil {
		parentSpan.RecordError(err)
		return resp.StatusCode(), nil, trace, err
	}
	return resp.StatusCode(), resp.Body(), trace, nil
}

func HttpPostFormData(context context.Context, url string, formData map[string]string, header map[string]string) (int, []byte, resty.TraceInfo, error) {
	_, parentSpan := tracer.Start(context, "HttpRequest.HttpPostFormData")
	defer parentSpan.End()

	parentSpan.AddEvent(fmt.Sprintf("GetHttRequest %s", url))
	req := createNewRequest().SetHeaders(header).SetFormData(formData)
	resp, err := req.Execute(http.MethodPost, url)
	trace := resp.Request.TraceInfo()

	if err != nil {
		parentSpan.RecordError(err)
		return resp.StatusCode(), nil, trace, err
	}
	return resp.StatusCode(), resp.Body(), trace, nil
}

// HTTPPutWithHeader func
func HttpPut(context context.Context, url string, jsondata interface{}, header map[string]string) (int, []byte, resty.TraceInfo, error) {
	_, parentSpan := tracer.Start(context, "HttpRequest.HttpPut")
	defer parentSpan.End()

	parentSpan.AddEvent(fmt.Sprintf("GetHttRequest %s", url))
	req := createNewRequest().SetHeaders(header).SetBody(jsondata)
	resp, err := req.Execute(http.MethodPut, url)
	trace := resp.Request.TraceInfo()

	if err != nil {
		parentSpan.RecordError(err)
		return resp.StatusCode(), nil, trace, err
	}
	return resp.StatusCode(), resp.Body(), trace, nil
}

func HttpPatch(context context.Context, url string, jsondata interface{}, header map[string]string) (int, []byte, resty.TraceInfo, error) {
	_, parentSpan := tracer.Start(context, "HttpRequest.HttpPatch")
	defer parentSpan.End()

	parentSpan.AddEvent(fmt.Sprintf("GetHttRequest %s", url))
	req := createNewRequest().SetHeaders(header).SetBody(jsondata)
	resp, err := req.Execute(http.MethodPatch, url)
	trace := resp.Request.TraceInfo()

	if err != nil {
		parentSpan.RecordError(err)
		return resp.StatusCode(), nil, trace, err
	}
	return resp.StatusCode(), resp.Body(), trace, nil
}

// HTTPDeleteWithHeader func
func HttpDelete(context context.Context, url string, jsondata interface{}, header map[string]string) (int, []byte, resty.TraceInfo, error) {
	_, parentSpan := tracer.Start(context, "HttpRequest.HttpDelete")
	defer parentSpan.End()

	parentSpan.AddEvent(fmt.Sprintf("GetHttRequest %s", url))
	req := createNewRequest().SetHeaders(header).SetBody(jsondata)
	resp, err := req.Execute(http.MethodDelete, url)
	trace := resp.Request.TraceInfo()

	if err != nil {
		parentSpan.RecordError(err)
		return resp.StatusCode(), nil, trace, err
	}
	return resp.StatusCode(), resp.Body(), trace, nil
}

// SendHttpRequest ..
func SendHttpRequest(context context.Context, method string, url string, header map[string]string, body interface{}) ([]byte, int, resty.TraceInfo, error) {
	var data []byte
	var err error
	var trace resty.TraceInfo
	var statusCode int
	switch method {
	case http.MethodGet:
		statusCode, data, trace, err = HTTPGet(context, url, header)
	case http.MethodPost:
		statusCode, data, trace, err = HTTPPost(context, url, body, header)
	case http.MethodPut:
		statusCode, data, trace, err = HttpPut(context, url, body, header)
	case http.MethodDelete:
		statusCode, data, trace, err = HttpDelete(context, url, body, header)
	default:
		err = errors.New("method not found")
	}
	return data, statusCode, trace, err
}

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
