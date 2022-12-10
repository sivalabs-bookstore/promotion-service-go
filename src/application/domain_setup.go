package application

import (
	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/repository"
	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/usecase"
	repository_impl "github.com/sivalabs-bookstore/promotion_service_go/src/infrastructure/db/memory"
	"go.uber.org/zap"
)

type domainSetup struct {
	UseCase           usecase.GetPromotionsByIsbnsUseCase
	useCaseRepository repository.PromotionRepository
}

func createDomainSetup(logger *zap.Logger) domainSetup {
	useCaseRepository := repository_impl.CreateInMemoryPromotionRepository(logger)
	useCase := usecase.CreateGetPromotionsByIsbnsUseCase(useCaseRepository, logger)
	return domainSetup{
		UseCase:           useCase,
		useCaseRepository: useCaseRepository,
	}
}
