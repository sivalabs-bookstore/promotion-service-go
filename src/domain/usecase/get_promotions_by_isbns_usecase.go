package usecase

import (
	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/command"
	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/exception"
	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/repository"
	"go.uber.org/zap"
)

type GetPromotionsByIsbnsUseCase interface {
	GetPromotionsByIsbns(command command.GetPromotionsByIsbnsCommand) ([]command.GetPromotionsByIsbnsResult, error)
}

type defaultGetPromotionsByIsbnsUseCase struct {
	promotionRepository repository.PromotionRepository
	logger              *zap.Logger
}

func CreateGetPromotionsByIsbnsUseCase(repository repository.PromotionRepository, logger *zap.Logger) GetPromotionsByIsbnsUseCase {
	return defaultGetPromotionsByIsbnsUseCase{
		promotionRepository: repository,
		logger:              logger,
	}
}

func (useCase defaultGetPromotionsByIsbnsUseCase) GetPromotionsByIsbns(cmd command.GetPromotionsByIsbnsCommand) ([]command.GetPromotionsByIsbnsResult, error) {
	results, err := useCase.promotionRepository.GetPromotions(cmd.Isbns)
	if len(results) == 0 {
		err = exception.CreateNotFoundError("No promotions found")
	}
	return command.FromPromotions(results), err
}
