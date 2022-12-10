package application

import (
	"fmt"
	"net/http"

	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/usecase"
	"github.com/sivalabs-bookstore/promotion_service_go/src/infrastructure/http/handler"
	"go.uber.org/zap"
)

func setupPromotionsHandler(useCase usecase.GetPromotionsByIsbnsUseCase, logger *zap.Logger) {
	ph := handler.CreatePromotionsHandler(useCase, logger)
	ph.HandleRequests()
}

func setupMux(port int, configureHandlers func()) {
	host := fmt.Sprintf(":%d", port)
	configureHandlers()
	http.ListenAndServe(host, nil)
}
