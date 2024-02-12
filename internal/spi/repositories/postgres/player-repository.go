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

const (
	tableName = "player"
)

type PostgresPlayerRepository struct {
	client *BunPostgresDatabaseClient
}

func NewPlayerRepository(dbClient *BunPostgresDatabaseClient) *PostgresPlayerRepository {
	log.Info().Msg("Player repository initialized.")
	return &PostgresPlayerRepository{
		client: dbClient,
	}
}

func (repo *PostgresPlayerRepository) CreateOne(ctx model.Context, createPlayerRequest spiports.CreatePlayerRequest) (*model.Player, error) {
	player := &model.Player{}

	_, err := repo.client.DB.NewInsert().Model(&createPlayerRequest).ModelTableExpr(tableName).Returning("*").Exec(ctx, player)

	if err != nil {
		return nil, NewUnkownDatabaseError(err)
	}
	return player, nil
}

func (repo *PostgresPlayerRepository) GetOne(ctx model.Context, gid string) (*model.Player, error) {
	player := &model.Player{}

	err := repo.client.DB.NewSelect().Model(player).ModelTableExpr(tableName).Where("gid = ?", gid).Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewBeeNotFoundError(fmt.Errorf("Player with gid %s could not be found.", gid))
		}
		return nil, NewUnkownDatabaseError(err)
	}

	return player, nil
}

func (repo *PostgresPlayerRepository) GetMany(ctx model.Context, query apiports.GetManyPlayersQuery) ([]model.Player, model.ResponseMeta, error) {
	players := make([]model.Player, 0)
	pagination := dbutils.GetDefaultPagination(query.Pagination)
	responseMeta := model.ResponseMeta{
		Pagination: pagination,
		ItemsTotal: 0,
		PagesTotal: 1,
	}

	dbQuery := repo.client.DB.NewSelect().Model(&players).ModelTableExpr(tableName)

	dbutils.UrlToDbQuery(dbQuery, query)

	dbutils.UrlToDbPagination(dbQuery, pagination)

	err := dbutils.DbQuerySort(dbQuery, new(model.Player), model.Sorting{SortBy: query.SortBy, Sort: query.Sort})

	if err != nil {
		return players, responseMeta, err
	}

	count, err := dbQuery.ScanAndCount(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return players, responseMeta, nil
		}
		return players, responseMeta, NewUnkownDatabaseError(err)
	}

	responseMeta.ItemsTotal = count
	responseMeta.PagesTotal = dbutils.GetPagesTotal(count, pagination.PageSize)

	return players, responseMeta, nil
}

func (repo *PostgresPlayerRepository) UpdateOne(ctx model.Context, gid string, updatePlayerRequest spiports.UpdatePlayerRequest) (*model.Player, error) {
	bee := &model.Player{}

	_, err := repo.client.DB.NewUpdate().OmitZero().Model(&updatePlayerRequest).Where("gid = ?", gid).ModelTableExpr(tableName).Returning("*").Exec(ctx, bee)

	if err != nil {
		return nil, NewUnkownDatabaseError(err)
	}
	return bee, nil
}
