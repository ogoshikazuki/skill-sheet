package usecase

import (
	"context"

	"github.com/ogoshikazuki/skill-sheet/entity"
)

type (
	UpdateBasicInformationUsecase struct {
		repository            entity.BasicInformationRepository
		transactionController entity.TransactionController
	}
)

func NewUpdateBasicInformationUsecase(
	repository entity.BasicInformationRepository,
	transactionController entity.TransactionController,
) UpdateBasicInformationUsecase {
	return UpdateBasicInformationUsecase{
		repository:            repository,
		transactionController: transactionController,
	}
}

func (u UpdateBasicInformationUsecase) Handle(ctx context.Context, input entity.UpdateBasicInformationInput) (entity.BasicInformation, error) {
	err := u.transactionController.Transaction(ctx, func(tx entity.Tx) error {
		return u.repository.Update(ctx, tx, input)
	})
	if err != nil {
		return entity.BasicInformation{}, err
	}

	basicInformation, err := u.repository.Find(ctx)
	if err != nil {
		return entity.BasicInformation{}, err
	}

	return basicInformation, nil
}
