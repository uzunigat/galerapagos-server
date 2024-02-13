package postgres

import (
	"database/sql"
	"fmt"

	apiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/api-ports"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	spiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/spi-ports"
	dbutils "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/spi/repositories/postgres/utils"
	"github.com/rs/zerolog/log"
)

type PostgresGameRepository struct {
	client *BunPostgresDatabaseClient
}

func NewGameRepository(dbClient *BunPostgresDatabaseClient) *PostgresGameRepository {
	log.Info().Msg("Game repository initialized.")
	return &PostgresGameRepository{
		client: dbClient,
	}
}

func (repo *PostgresGameRepository) CreateOne(ctx model.Context, createGameRequest spiports.CreateGameRequest) (*model.Game, error) {
	game := &model.Game{}

	_, err := repo.client.DB.NewInsert().Model(&createGameRequest).ModelTableExpr(tableGame).Returning("*").Exec(ctx, game)

	if err != nil {
		return nil, NewUnkownDatabaseError(err)
	}
	return game, nil
}

func (repo *PostgresGameRepository) GetOne(ctx model.Context, gid string) (*model.Game, error) {
	game := &model.Game{}

	err := repo.client.DB.NewSelect().Model(game).ModelTableExpr(tableGame).Where("gid = ?", gid).Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewGameNotFoundError(fmt.Errorf("Game with gid %s could not be found.", gid))
		}
		return nil, NewUnkownDatabaseError(err)
	}

	return game, nil
}

func (repo *PostgresGameRepository) GetMany(ctx model.Context, query apiports.GetManyGamesQuery) ([]model.Game, model.ResponseMeta, error) {
	games := make([]model.Game, 0)
	pagination := dbutils.GetDefaultPagination(query.Pagination)
	responseMeta := model.ResponseMeta{
		Pagination: pagination,
		ItemsTotal: 0,
		PagesTotal: 1,
	}

	dbQuery := repo.client.DB.NewSelect().Model(&games).ModelTableExpr(tableGame)

	dbutils.UrlToDbQuery(dbQuery, query)

	dbutils.UrlToDbPagination(dbQuery, pagination)

	err := dbutils.DbQuerySort(dbQuery, new(model.Game), model.Sorting{SortBy: query.SortBy, Sort: query.Sort})

	if err != nil {
		return games, responseMeta, err
	}

	count, err := dbQuery.ScanAndCount(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return games, responseMeta, nil
		}
		return games, responseMeta, NewUnkownDatabaseError(err)
	}

	responseMeta.ItemsTotal = count
	responseMeta.PagesTotal = dbutils.GetPagesTotal(count, pagination.PageSize)

	return games, responseMeta, nil
}

func (repo *PostgresGameRepository) UpdateOne(ctx model.Context, gid string, updateGameRequest spiports.UpdateGameRequest) (*model.Game, error) {
	game := &model.Game{}

	_, err := repo.client.DB.NewUpdate().OmitZero().Model(&updateGameRequest).Where("gid = ?", gid).ModelTableExpr(tableGame).Returning("*").Exec(ctx, game)

	if err != nil {
		return nil, NewUnkownDatabaseError(err)
	}
	return game, nil
}
