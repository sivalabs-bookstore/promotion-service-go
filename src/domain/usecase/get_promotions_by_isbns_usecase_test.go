package usecase

import (
	"testing"

	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/command"
	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/entity"
	"github.com/sivalabs-bookstore/promotion_service_go/src/domain/exception"
	"go.uber.org/zap"
)

type mockRepository struct {
}

func (mr mockRepository) GetPromotions(isbns []string) ([]entity.Promotion, error) {
	if len(isbns) == 0 {
		return []entity.Promotion{}, nil
	}
	if isbns[0] == "isbn-1" {
		return []entity.Promotion{{
			Isbn:     "isbn-1",
			Discount: 1,
		}}, nil
	}
	return []entity.Promotion{}, nil
}

func createLogger() *zap.Logger {
	return zap.NewNop()
}

func TestGetPromotionsByIsbns(t *testing.T) {
	useCase := CreateGetPromotionsByIsbnsUseCase(mockRepository{}, createLogger())
	cmd := command.GetPromotionsByIsbnsCommand{
		Isbns: []string{"isbn-1"},
	}
	results, _ := useCase.GetPromotionsByIsbns(cmd)
	want := []command.GetPromotionsByIsbnsResult{
		{
			Isbn:     "isbn-1",
			Discount: 1,
		},
	}
	for i, result := range results {
		if result != want[i] {
			t.Errorf("got %v wanted %v", result, want[i])
		}
	}
}
func TestGetPromotionsByIsbnsNotFoundError(t *testing.T) {
	useCase := CreateGetPromotionsByIsbnsUseCase(mockRepository{}, createLogger())
	cmd := command.GetPromotionsByIsbnsCommand{
		Isbns: []string{"isbn-2"},
	}
	results, err := useCase.GetPromotionsByIsbns(cmd)
	_, errWant := []command.GetPromotionsByIsbnsResult{}, exception.CreateNotFoundError("No promotions found")
	if len(results) > 0 {
		t.Errorf("got %d wanted %d", len(results), 0)
	}
	if err.Error() != errWant.Error() {
		t.Errorf("got %v wanted %v", err.Error(), errWant.Error())
	}
}
