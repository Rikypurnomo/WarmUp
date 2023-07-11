package logger

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Rikypurnomo/warmup/config"
	"github.com/Rikypurnomo/warmup/pkg/util"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func SetupLogger() {
	level, err := logrus.ParseLevel(config.LogLevel())
	if err != nil {
		level = logrus.WarnLevel
	}

	logger = &logrus.Logger{
		Out:   os.Stdout,
		Hooks: make(logrus.LevelHooks),
		Level: level,
	}

	if config.IsLogTypeJson() {
		logger.Formatter = &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
	} else {
		logger.Formatter = &logrus.TextFormatter{
			ForceColors:     true,
			TimestampFormat: "2006-01-02 15:04:05",
		}
	}
}

func AddHook(hook logrus.Hook) {
	logger.Hooks.Add(hook)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

func Debugrf(r *http.Request, format string, args ...interface{}) {
	httpRequestLogEntry(r).Debugf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Errorln(args ...interface{}) {
	logger.Errorln(args...)
}

func Errorrf(r *http.Request, format string, args ...interface{}) {
	httpRequestLogEntry(r).Errorf(format, args...)
}

func ErrorWithFieldsf(fields logrus.Fields, format string, args ...interface{}) {
	logger.WithFields(fields).Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

func Inforf(r *http.Request, format string, args ...interface{}) {
	httpRequestLogEntry(r).Infof(format, args...)
}

func InfoWithFieldsf(fields logrus.Fields, format string, args ...interface{}) {
	logger.WithFields(fields).Infof(format, args...)
}

func ProxyInfo(aclName string, downstreamHost string, r *http.Request, responseStatus int, rw http.ResponseWriter) {
	logger.WithFields(logrus.Fields{
		"type":            "proxy",
		"downstream_host": downstreamHost,
		"api_name":        aclName,
		"request":         httpRequestFields(r),
		"response":        httpResponseFields(responseStatus, rw),
	}).Info("proxy")
}

func ApiLogger(r *http.Request, responseStatus int, responseBody, requestString, traceId string) {
	var res interface{}
	err := json.Unmarshal([]byte(responseBody), &res)
	if err != nil {
		res = responseBody
	}
	logger.WithFields(logrus.Fields{
		"type":    "api",
		"method":  r.Method,
		"headers": r.Header,
		"request": httpRequestFields(r),
		"response": logrus.Fields{
			"status": responseStatus,
			"body":   res,
		},
		"request_id": requestString,
		"trace_id":   traceId,
	}).Info("api")
}

func httpRequestFields(r *http.Request) logrus.Fields {
	requestHeaders := map[string]string{}
	for k := range r.Header {
		normalizedKey := util.ToSnake(k)
		if normalizedKey == "authorization" {
			continue
		}

		requestHeaders[normalizedKey] = r.Header.Get(k)

	}
	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
	}
	// restore body
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return logrus.Fields{
		"uri":     r.URL.String(),
		"query":   r.URL.Query(),
		"method":  r.Method,
		"headers": requestHeaders,
		"body":    string(body),
	}
}

func httpResponseFields(responseStatus int, rw http.ResponseWriter) logrus.Fields {
	responseHeaders := map[string]string{}
	for k := range rw.Header() {
		responseHeaders[util.ToSnake(k)] = rw.Header().Get(k)

	}
	return logrus.Fields{
		"status":  responseStatus,
		"headers": responseHeaders,
	}
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}

func httpRequestLogEntry(r *http.Request) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"request_method": r.Method,
		"request_host":   r.Host,
		"request_url":    r.URL.String(),
	})
}
