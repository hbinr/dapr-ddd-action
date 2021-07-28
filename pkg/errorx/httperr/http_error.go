package httperr

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/dapr-ddd-action/pkg/errorx"
	"github.com/go-chi/render"
)

func InternalError(slug string, msg string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, msg, w, r, "Internal server error", http.StatusInternalServerError)
}

func Unauthorised(slug string, msg string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, msg, w, r, "Unauthorised", http.StatusUnauthorized)
}

func BadRequest(slug string, msg string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, msg, w, r, "Bad request", http.StatusBadRequest)
}

func RespondWithError(err error, msg string, w http.ResponseWriter, r *http.Request) {
	respErr, ok := err.(errorx.SlugErr)
	if !ok {
		InternalError("internal-server-error", msg, err, w, r)
		return
	}

	switch respErr.ErrorType() {
	case errorx.ErrorTypeAuthorization:
		Unauthorised(respErr.Slug(), respErr.Error(), respErr, w, r)
	case errorx.ErrorTypeIncorrectInput:
		BadRequest(respErr.Slug(), respErr.Error(), respErr, w, r)
	default:
		InternalError(respErr.Slug(), respErr.Error(), respErr, w, r)
	}
}

func httpRespondWithError(err error, slug string, msg string,
	w http.ResponseWriter, r *http.Request, logMsg string, status int) {
	resp := ErrorResponse{slug, msg, status}

	zap.L().Error(logMsg, zap.Error(err))
	if err := render.Render(w, r, resp); err != nil {
		panic(err)
	}
}

type ErrorResponse struct {
	Slug       string `json:"slug"`
	Msg        string `json:"msg"`
	httpStatus int
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(e.httpStatus)
	return nil
}
