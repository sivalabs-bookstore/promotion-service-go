package repository_impl

import (
	"strings"
	"time"

	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/entity"
	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/repository"
	"go.uber.org/zap"
)

type inMemoryPromotionRepository struct {
	promotions map[int]entity.Promotion
	logger     *zap.Logger
}

func CreateInMemoryPromotionRepository(logger *zap.Logger) repository.PromotionRepository {
	promotion1 := entity.Promotion{
		Id:          1,
		Isbn:        "ISBN-1",
		Discount:    1.5,
		Enabled:     true,
		CreatedDate: time.Now(),
	}
	promotion2 := entity.Promotion{
		Id:          2,
		Isbn:        "ISBN-3",
		Discount:    3.5,
		Enabled:     true,
		CreatedDate: time.Now(),
	}
	promotions := make(map[int]entity.Promotion)
	promotions[promotion1.Id] = promotion1
	promotions[promotion2.Id] = promotion2
	return inMemoryPromotionRepository{
		promotions: promotions,
		logger:     logger,
	}
}

func (inMemoryPromotionRepository inMemoryPromotionRepository) GetPromotions(isbns []string) ([]entity.Promotion, error) {
	var result []entity.Promotion
	logger := inMemoryPromotionRepository.logger
	logger.Debug("isbns to search", zap.Any("isbns", isbns))
	logger.Debug("Looking the isbns in memory")
	for _, promotion := range inMemoryPromotionRepository.promotions {
		for _, isbn := range isbns {
			if strings.EqualFold(promotion.Isbn, isbn) {
				result = append(result, promotion)
				break
			}
		}
	}
	return result, nil
}
