package config

import (
	"fmt"
	"time"

	"gorm.io/gorm/logger"
)

func RedisToAddr() string {
	return fmt.Sprintf("%s:%s", appConfig.Redis.Host, appConfig.Redis.Port)
}

func PasswordRedis() string {
	return appConfig.Redis.Password
}

func PGToAddr() string {
	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%s", appConfig.Postgresql.Host, appConfig.Postgresql.Port, appConfig.Postgresql.Username, appConfig.Postgresql.Password, appConfig.Postgresql.Database, appConfig.Postgresql.SslMode, appConfig.Postgresql.Timeout)
	return args
}

func GetDbName() string {
	return appConfig.Postgresql.Database
}

func LogLevePG() logger.LogLevel {
	switch appConfig.Postgresql.LogLevel {
	case "error":
		return logger.Error
	case "info":
		return logger.Info
	}
	return logger.Silent
}

func LogLevel() string {
	return appConfig.Config.LogLevel
}

func IsLogTypeJson() bool {
	return appConfig.Config.LogType == LogTypeJson
}

func IsLogTypeText() bool {
	return appConfig.Config.LogType == LogTypeText
}

func IsDev() bool {
	return appConfig.Config.Env == Development
}

func IsStaging() bool {
	return appConfig.Config.Env == Staging
}

func IsProduction() bool {
	return appConfig.Config.Env == Production
}

func TimeLocation() (*time.Location, error) {
	return time.LoadLocation(appConfig.Config.LocationTime)
}

func GetHost() string {
	return fmt.Sprintf("%s:%s", appConfig.Config.Host, appConfig.Config.Port)
}

func ServerReadTimeoutInSecond() time.Duration {
	return time.Duration(appConfig.Config.ReadTimeoutSecond * int(time.Second))
}

func ServerWriteTimeoutInSecond() time.Duration {
	return time.Duration(appConfig.Config.WriteTimeoutSecond * int(time.Second))
}

func ServerIdleTimeoutInSecond() time.Duration {
	return time.Duration(appConfig.Config.IdleTimeoutSecond * int(time.Second))
}

func GetMaxAgeDuration() time.Duration {
	return time.Duration(appConfig.Cors.MaxAge)
}

func MessageError(msg string) string {
	if msg == "" {
		return MessageErrorDefault
	}
	return msg
}

func MessageSuccess(msg string) string {
	if msg == "" {
		return MessageSuccessDefault
	}
	return msg
}

func HttpReqTimeout() time.Duration {
	return time.Duration(appConfig.HTTPReq.Timeout)
}

func HttpReqRetry() int {
	return appConfig.HTTPReq.Retry
}

func HttpReqDebug() bool {
	return appConfig.HTTPReq.Debug
}

func AdapterAzHost() string {
	return fmt.Sprintf("%s:%s", appConfig.AdapterAz.Host, appConfig.AdapterAz.Port)
}

func IsEnabledPG() bool {
	return appConfig.Postgresql.Enabled
}

func IsEnabledRedis() bool {
	return appConfig.Redis.Enabled
}

func IsEnabledAdapterAz() bool {
	return appConfig.AdapterAz.Enable
}

func NameServices() string {
	return appConfig.Config.NameServices
}

func IsEnabledOpenTelemetry() bool {
	return appConfig.OpenTelemetry.Enabled
}

func OtelToAddr() string {
	return fmt.Sprintf("%s:%s", appConfig.OpenTelemetry.Host, appConfig.OpenTelemetry.Port)
}
