package postgres

import "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"

type WreckRepository struct {
	client *BunPostgresDatabaseClient
}

func NewWreckRepository(client *BunPostgresDatabaseClient) *WreckRepository {
	return &WreckRepository{client: client}
}

func (repository *WreckRepository) GetAll(ctx model.Context) ([]model.WreckCard, error) {
	wreckCards := make([]model.WreckCard, 0)

	err := repository.client.DB.NewSelect().Model(&wreckCards).ModelTableExpr(tableWreckCard).OrderExpr("RANDOM()").Scan(ctx)

	if err != nil {
		return nil, NewUnkownDatabaseError(err)
	}

	return wreckCards, nil
}
