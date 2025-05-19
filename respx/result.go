package respx

import (
	"fmt"
	"github.com/SpectatorNan/go-zero-i18n/goi18nx"
	"github.com/SpectatorNan/goutils/errors"
	errorx2 "github.com/SpectatorNan/goutils/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"net/http"
)

var okStatusCode = http.StatusOK

func SetOkStatusCode(code int) {
	okStatusCode = code
}

var errRequestStatusCode = http.StatusOK

func SetErrStatusCode(code int) {
	errRequestStatusCode = code
}

var forbiddenStatusCode = http.StatusForbidden

var debugMode = false

func SetDebugMode(mode bool) {
	debugMode = mode
}

// HttpResult
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	ctx := r.Context()
	statusCode := okStatusCode
	if err == nil {
		//成功返回
		rp := NewSuccessResponse(resp)
		httpx.WriteJson(w, statusCode, rp)
	} else {
		//错误返回
		dfe := errorx2.DefaultErr
		errCode := dfe.Code
		errmsg := dfe.DefaultMsg
		if goi18nx.IsHasI18n(ctx) {
			errmsg = goi18nx.FormatText(ctx, dfe.MsgKey, dfe.DefaultMsg)
		}
		statusCode = errRequestStatusCode

		causeErr := errors.Cause(err)
		var codeE *errorx2.CodeError
		var i18nE *errorx2.I18nCodeError
		var forbiddenE *errorx2.ForbiddenError
		// err类型
		if errors.As(causeErr, &forbiddenE) {
			errCode = uint32(forbiddenStatusCode)
			errmsg = forbiddenE.Message
			statusCode = forbiddenStatusCode
			if len(forbiddenE.Reason) > 0 && debugMode {
				errmsg = fmt.Sprintf("%s, %s", forbiddenE.Message, forbiddenE.Reason)
			}
		} else if errors.As(causeErr, &codeE) { //自定义错误类型
			//自定义CodeError
			if codeE.Code != dfe.Code {
				errCode = codeE.Code
				errmsg = codeE.Message
				if len(codeE.Reason) > 0 && debugMode {
					errmsg = fmt.Sprintf("%s, %s", codeE.Message, codeE.Reason)
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
			dfe = errorx2.NotFoundResourceErr
			errCode = dfe.Code
			if goi18nx.IsHasI18n(ctx) {
				errmsg = goi18nx.FormatText(ctx, dfe.MsgKey, dfe.DefaultMsg)
			} else {
				errmsg = dfe.DefaultMsg
			}
		} else if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
			grpcCode := uint32(gstatus.Code())
			if grpcCode != errorx2.ErrCodeDefault {
				errCode = grpcCode
				errmsg = gstatus.Message()
			}
			//if grpcCode == uint32(forbiddenStatusCode) {
			//	statusCode = forbiddenStatusCode
			//}
			for _, detail := range gstatus.Details() {
				if info, ok := detail.(*errdetails.ErrorInfo); ok {
					domain := info.Domain
					//reason := info.Reason

					if domain == errorx2.GrpcErrorInfoDomain_Forbidden {
						//errCode = uint32(forbiddenStatusCode)
						statusCode = forbiddenStatusCode
					}
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 %+v ", err)
		if debugMode {
			errmsg = err.Error()
		}
		httpx.WriteJson(w, statusCode, NewErrorResponse(errCode, errmsg))
	}
}

// http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	dfe := errorx2.RequestParamsErr
	msg := goi18nx.FormatText(r.Context(), dfe.MsgKey, dfe.DefaultMsg)
	errMsg := fmt.Sprintf("%s, %s", msg, err.Error())
	logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
	logx.WithContext(r.Context()).Errorf("【API-ERR】 reason: %+v ", errMsg)
	httpx.WriteJson(w, errRequestStatusCode, NewErrorResponse(errorx2.ErrCodeRequestParams, errMsg))
}
