package repository

import (
	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/entity"
)

type PromotionRepository interface {
	GetPromotions(isbns []string) ([]entity.Promotion, error)
}
