package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/command"
	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/exception"
	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/usecase"
	"go.uber.org/zap"
)

type promotionHandler struct {
	getPromotionsByIsbnsUseCase usecase.GetPromotionsByIsbnsUseCase
	logger                      *zap.Logger
}

func CreatePromotionsHandler(useCase usecase.GetPromotionsByIsbnsUseCase, logger *zap.Logger) promotionHandler {
	return promotionHandler{
		getPromotionsByIsbnsUseCase: useCase,
		logger:                      logger,
	}
}

func (ph promotionHandler) HandleRequests() {
	ph.logger.Info("Configuring promotion handler")
	http.HandleFunc("/promotions", ph.handle)
}

func (ph promotionHandler) handle(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	ph.logger.Info(fmt.Sprintf("Path: %s, Method: %s", r.URL.Path, r.Method))
	switch r.Method {
	case "GET":
		if r.URL.Path == "/promotions" {
			err = ph.getPromotionsForIsbns(w, r)
		} else {
			findBySingleIsbnRegexp := regexp.MustCompile(`/promotions/.+`)
			if findBySingleIsbnRegexp.MatchString(r.URL.Path) {
				/*  TODO */
			}
		}
	default:
		err = exception.CreateNotFoundError("URI not found")
	}
	if err != nil {
		errorHandler(err, w, ph.logger)
	}
}

func (ph promotionHandler) getPromotionsForIsbns(w http.ResponseWriter, r *http.Request) error {
	rawIsbns := strings.TrimSpace(r.URL.Query().Get("isbn"))
	if rawIsbns == "" {
		return exception.CreateInvalidArgumentError("isbn argument cannot be empty")
	}
	cmd := command.GetPromotionsByIsbnsCommand{
		Isbns: strings.Split(rawIsbns, ","),
	}
	result, err := ph.getPromotionsByIsbnsUseCase.GetPromotionsByIsbns(cmd)
	if err != nil {
		return err
	}
	encode(w)(result)
	return nil
}
