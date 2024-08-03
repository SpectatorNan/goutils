package respx

import (
	"fmt"
	"github.com/SpectatorNan/go-zero-i18n/goi18nx"
	"github.com/SpectatorNan/goutils/common/errorx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"net/http"
)

var errRequestStatusCode = http.StatusOK

func SetErrStatusCode(code int) {
	errRequestStatusCode = code
}

var okStatusCode = http.StatusOK

func SetOkStatusCode(code int) {
	okStatusCode = code
}

var debugMode = false

func SetDebugMode(mode bool) {
	debugMode = mode
}

// HttpResult
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	ctx := r.Context()
	if err == nil {

		//成功返回
		rp := NewSuccessResponse(resp)
		httpx.WriteJson(w, okStatusCode, rp)
	} else {
		//错误返回
		dfe := errorx.DefaultErr
		errCode := dfe.Code
		errmsg := goi18nx.FormatText(ctx, dfe.MsgKey, dfe.DefaultMsg)
		//errreason := err.Error()

		causeErr := errors.Cause(err)
		var codeE *errorx.CodeError
		var i18nE *errorx.I18nCodeError
		// err类型

		if errors.As(causeErr, &codeE) { //自定义错误类型
			//自定义CodeError
			if codeE.Code != dfe.Code {
				errCode = codeE.Code
				errmsg = codeE.Message
				if len(codeE.Reason) > 0 {
					//errreason = codeE.Reason
				}
			} else {
				errmsg = codeE.Message
			}
		} else if errors.As(causeErr, &i18nE) { //自定义国际化错误类型
			if goi18nx.IsHasI18n(ctx) {
				errmsg = goi18nx.FormatText(ctx, i18nE.MsgKey, i18nE.DefaultMsg)
			} else {
				errmsg = dfe.DefaultMsg
			}
			errCode = i18nE.Code
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			dfe = errorx.NotFoundResourceErr
			errCode = dfe.Code
			if goi18nx.IsHasI18n(ctx) {
				errmsg = goi18nx.FormatText(ctx, dfe.MsgKey, dfe.DefaultMsg)
			} else {
				errmsg = dfe.DefaultMsg
			}
			httpx.WriteJson(w, errRequestStatusCode, NewErrorResponse(errCode, errmsg))
			return
		} else if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
			grpcCode := uint32(gstatus.Code())
			if grpcCode != errorx.ErrCodeDefault {
				errCode = grpcCode
				errmsg = gstatus.Message()
			}
		}

		//logx.WithContext(r.Context()).Errorf("【API-ERR】: %+v ", errreason)
		logx.WithContext(r.Context()).Errorf("【API-ERR】 %+v ", err)
		if debugMode {
			//errmsg = errreason
			errmsg = err.Error()
		}
		httpx.WriteJson(w, errRequestStatusCode, NewErrorResponse(errCode, errmsg))
	}
}

// http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	dfe := errorx.RequestParamsErr
	msg := goi18nx.FormatText(r.Context(), dfe.MsgKey, dfe.DefaultMsg)
	errMsg := fmt.Sprintf("%s, %s", msg, err.Error())
	logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
	logx.WithContext(r.Context()).Errorf("【API-ERR】 reason: %+v ", errMsg)
	httpx.WriteJson(w, errRequestStatusCode, NewErrorResponse(errorx.ErrCodeRequestParams, errMsg))
}
