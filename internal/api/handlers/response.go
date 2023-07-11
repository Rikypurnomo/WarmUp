package handlers

import (
	"bytes"
	"encoding/json"

	"github.com/Rikypurnomo/warmup/config"
	"github.com/Rikypurnomo/warmup/pkg/logger"
	"github.com/Rikypurnomo/warmup/pkg/util"
	models_server "github.com/Rikypurnomo/warmup/server/models"
	"github.com/gin-gonic/gin"
)

func errorResponse(ctx *gin.Context, code int, messageError string) {
	response(ctx, code, models_server.Response{
		Meta: models_server.Meta{
			Status:  false,
			Code:    code,
			Message: config.MessageError(messageError),
		},
	})
}

func successResponse(ctx *gin.Context, code int, messageSuccess string, data interface{}) {
	response(ctx, code, models_server.Response{
		Meta: models_server.Meta{
			Status:  true,
			Code:    code,
			Message: config.MessageSuccess(messageSuccess),
		},
		Data: data,
	})
}

func successResponsePagination(ctx *gin.Context, code int, messageSuccess string, data interface{}, pagination models_server.MetaPagination) {
	response(ctx, code, models_server.ResponsePagination{
		Meta: models_server.Meta{
			Status:  true,
			Code:    code,
			Message: config.MessageSuccess(messageSuccess),
		},
		Pagination: pagination,
		Data:       data,
	})
}


func response(c *gin.Context, code int, res interface{}) {
	mars, _ := json.Marshal(res)
	getTrace := util.GetTraceIdFromContext(c.Request.Context())
	util.SetAttributeReqRes(c.Request, string(mars))
	logger.ApiLogger(c.Request, code, string(mars), c.GetString("X-Request-Id"), getTrace)
	c.JSON(code, res)
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body   *bytes.Buffer
	status int
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
