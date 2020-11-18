package response_helper

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

type ResponseHelper struct {
	Status  int                  `json:"status"`
	Msg     string               `json:"message"`
	Data    interface{}          `json:"data,omitempty"`
	ctx     iris.Context         `json:"-"`
	logFunc func(...interface{}) `json:"-"`
}

func NewResponseHelper(ctx *iris.Context, logFunc func(...interface{})) *ResponseHelper {
	return &ResponseHelper{
		ctx:     *ctx,
		Status:  200,
		Msg:     "ok",
		logFunc: logFunc,
	}
}

func (h *ResponseHelper) ResponseJSON(errLabel string) {
	_, err := h.ctx.JSON(h)
	if err != nil {
		h.ErrorLog(errLabel, err)
	}
}

func (h *ResponseHelper) Response(contentType string, errLabel string) {
	if h.Data != nil {
		h.ctx.Header("Content-Type", contentType)
		_, err := h.ctx.Write(h.Data.([]byte))
		if err != nil {
			h.ErrorLog(errLabel, err)
		}
	}
}

func (h *ResponseHelper) IsFailed(err error, status int) (isFailed bool) {
	isFailed = false
	if err != nil {
		h.Status = status
		h.Msg = err.Error()
		isFailed = true
	}

	return
}

func (h *ResponseHelper) ErrorLog(errLabel string, err error) {
	h.logFunc(fmt.Sprintf("%s: %v", errLabel, err))
}
