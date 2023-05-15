package respx

import (
	"fmt"
	"github.com/SpectatorNan/go-zero-i18n/goi18nx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"goutils/common/errorx"
	"goutils/common/requestx"
	"net/http"
)

// HttpResult
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	ctx := r.Context()
	if err == nil {
		claim, err := requestx.FetchUserByJwtClaims(r.Context())
		if err == nil {
			requestx.SetHeaderUserId(r, &claim.ID)
		}

		//成功返回
		r := NewSuccessResponse(resp)
		//w.Header().Add(projectConst.OperateLogResultHeaderKey, "ok")
		httpx.WriteJson(w, http.StatusOK, r)
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
			errmsg = goi18nx.FormatText(ctx, e.MsgKey, e.DefaultMsg)
			errCode = e.Code
		} else if err == gorm.ErrRecordNotFound {
			dfe = errorx.NotFoundResourceErr
			errCode = dfe.Code
			errmsg = goi18nx.FormatText(ctx, dfe.MsgKey, dfe.DefaultMsg)
			httpx.WriteJson(w, http.StatusOK, NewErrorResponse(errCode, errmsg))
			logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
			logx.WithContext(r.Context()).Errorf("【API-ERR】 reason: %+v ", errreason)
			return
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if grpcCode != errorx.ErrCodeDefault {
					errCode = grpcCode
					errmsg = gstatus.Message()

					var ice errorx.I18nCodeError
					if err := jsonx.Unmarshal([]byte(errmsg), &ice); err == nil {
						errmsg = goi18nx.FormatText(ctx, ice.MsgKey, ice.DefaultMsg)
						errCode = ice.Code
					}
					var ce errorx.CodeError
					if err := jsonx.Unmarshal([]byte(errmsg), &ce); err == nil {
						if ce.Code != dfe.Code {
							errmsg = ce.Message
							if len(ce.Reason) > 0 {
								errreason = ce.Reason
							}
						} else {
							errmsg = ce.Message
						}
					}
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
		logx.WithContext(r.Context()).Errorf("【API-ERR】 reason: %+v ", errreason)

		httpx.WriteJson(w, http.StatusOK, NewErrorResponse(errCode, errmsg))
	}
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
	httpx.WriteJson(w, http.StatusBadRequest, NewErrorResponse(errorx.ErrCodeRequestParams, errMsg))
}
