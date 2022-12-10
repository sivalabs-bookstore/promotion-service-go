package command

import "github.com/sivalabs-bookstore/promotion_service_go/src/domain/entity"

type GetPromotionsByIsbnsResult struct {
	Isbn     string  `json:"isbn"`
	Discount float64 `json:"discount"`
}

func fromPromotion(promotion entity.Promotion) GetPromotionsByIsbnsResult {
	return GetPromotionsByIsbnsResult{
		Isbn:     promotion.Isbn,
		Discount: promotion.Discount,
	}
}

func FromPromotions(promotions []entity.Promotion) []GetPromotionsByIsbnsResult {
	var results []GetPromotionsByIsbnsResult
	for _, promotion := range promotions {
		results = append(results, fromPromotion(promotion))
	}
	return results
}
