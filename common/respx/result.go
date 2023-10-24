package respx

import (
	"context"
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

var statusCode = http.StatusOK

func SetStatusCode(code int) {
	statusCode = code
}

// HttpResult
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	ctx := r.Context()
	if err == nil {

		//成功返回
		rp := NewSuccessResponse(resp)
		//w.Header().Add(projectConst.OperateLogResultHeaderKey, "ok")
		httpx.WriteJson(w, statusCode, rp)
	} else {
		//错误返回
		dfe := errorx.DefaultErr
		errCode := dfe.Code
		errmsg := goi18nx.FormatText(ctx, dfe.MsgKey, dfe.DefaultMsg)
		errreason := err.Error()

		causeErr := errors.Cause(err)                  // err类型
		if e, ok := causeErr.(*errorx.CodeError); ok { //自定义错误类型
			//自定义CodeError
			if e.Code != dfe.Code {
				errCode = e.Code
				//errmsg = goi18nx.FormatText(ctx, errorx.MapErrMsgKey(errcode), e.Message)
				errmsg = e.Message
				if len(e.Reason) > 0 {
					errreason = e.Reason
				}
			} else {
				errmsg = e.Message
			}
		} else if e, ok := causeErr.(*errorx.I18nCodeError); ok {
			if isHasI18n(ctx) {
				errmsg = goi18nx.FormatText(ctx, e.MsgKey, e.DefaultMsg)
			} else {
				errmsg = dfe.DefaultMsg
			}
			errCode = e.Code
		} else if err == gorm.ErrRecordNotFound {
			dfe = errorx.NotFoundResourceErr
			errCode = dfe.Code
			if isHasI18n(ctx) {
				errmsg = goi18nx.FormatText(ctx, dfe.MsgKey, dfe.DefaultMsg)
			} else {
				errmsg = dfe.DefaultMsg
			}
			httpx.WriteJson(w, statusCode, NewErrorResponse(errCode, errmsg))
			logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
			logx.WithContext(r.Context()).Errorf("【API-ERR】 reason: %+v ", errreason)
			return
		} else if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
			grpcCode := uint32(gstatus.Code())
			if grpcCode != errorx.ErrCodeDefault {
				// grpc err
				// must add interceptors in grpc server, like this:
				// s.AddUnaryInterceptors(interceptor.LoggerInterceptor)
				errCode = grpcCode
				errmsg = gstatus.Message()
					
			}
		}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
		logx.WithContext(r.Context()).Errorf("【API-ERR】 reason: %+v ", errreason)

		httpx.WriteJson(w, statusCode, NewErrorResponse(errCode, errmsg))
	}
}

func isHasI18n(ctx context.Context) bool {
	v := ctx.Value(goi18nx.I18nKey)
	if v != nil {
		return true
	}
	return false
}

//func ErrHandle(err error) (int, interface{}) {
//	switch e := err.(type) {
//	case *errorx.CodeError:
//		return http.StatusOK, NewErrorResponse(e.Code, e.Message)
//	default:
//		st, ok := status.FromError(err)
//		if ok {
//			return http.StatusOK, NewErrorResponse(uint32(st.Code()), st.Message())
//		}
//		return http.StatusOK, NewErrorReasonResponse(errorx.DEFAULT_ERROR, errorx.MapErrMsgKey(errorx.DEFAULT_ERROR), err.Error())
//	}
//}

//http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	dfe := errorx.RequestParamsErr
	msg := goi18nx.FormatText(r.Context(), dfe.MsgKey, dfe.DefaultMsg)
	errMsg := fmt.Sprintf("%s, %s", msg, err.Error())
	logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
	logx.WithContext(r.Context()).Errorf("【API-ERR】 reason: %+v ", errMsg)
	httpx.WriteJson(w, statusCode, NewErrorResponse(errorx.ErrCodeRequestParams, errMsg))
}
