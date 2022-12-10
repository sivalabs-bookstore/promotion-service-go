package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/exception"
	"go.uber.org/zap"
)

func encode(w io.Writer) func(e any) error {
	return func(e any) error {
		return json.NewEncoder(w).Encode(e)
	}
}

type ErrorResult struct {
	Message string `json:"message"`
}

func errorHandler(err error, w http.ResponseWriter, logger *zap.Logger) {
	serr, ok := err.(exception.NotFoundError)
	if ok {
		logger.Error(fmt.Sprintf("%s", serr.Msg()))
		er := ErrorResult{
			Message: serr.Msg(),
		}
		w.WriteHeader(http.StatusNotFound)
		encode(w)(er)
		return
	}
	serr2, ok := err.(exception.InvalidArgumentError)
	if ok {
		logger.Error(serr2.Msg())
		er := ErrorResult{
			Message: serr2.Msg(),
		}
		w.WriteHeader(http.StatusNotFound)
		encode(w)(er)
		return
	}
	logger.Error("Unrecognized error")
	logger.Error(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	er := ErrorResult{
		Message: err.Error(),
	}
	encode(w)(er)
}
