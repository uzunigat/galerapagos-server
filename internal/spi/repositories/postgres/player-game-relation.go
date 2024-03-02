package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	spiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/spi-ports"
)

type PlayerGameRelationRepository struct {
	client *BunPostgresDatabaseClient
}

func NewPlayerGameRelationRepository(dbClient *BunPostgresDatabaseClient) *PlayerGameRelationRepository {
	return &PlayerGameRelationRepository{
		client: dbClient,
	}
}

func (repo *PlayerGameRelationRepository) GetOne(ctx model.Context, playerGid string, gameGid string) (*model.PlayerGameRelation, error) {
	playerGameRelation := &model.PlayerGameRelation{}

	err := repo.client.DB.NewSelect().Model(playerGameRelation).ModelTableExpr(tablePlayerGameRelation).Where("player_gid = ? and game_gid = ?", playerGid, gameGid).Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewPlayerGameRelationNotFoundError(fmt.Errorf("Player with gid %s and game with gid %s could not be found.", playerGid, gameGid))
		}
		return nil, NewUnkownDatabaseError(err)
	}

	return playerGameRelation, nil
}

func (repo *PlayerGameRelationRepository) CreateOne(ctx model.Context, request spiports.CreateOnePlayerGameRelationRequest) (*model.PlayerGameRelation, error) {
	relation := &model.PlayerGameRelation{}

	_, err := repo.client.DB.NewInsert().Model(&request).ModelTableExpr(tablePlayerGameRelation).Returning("*").Exec(ctx, relation)

	if err != nil {
		return nil, err
	}

	return relation, nil
}

func (repo *PlayerGameRelationRepository) UpdateOne(ctx model.Context, gid string, request spiports.UpdateOnePlayerGameRelationRequest) (*model.PlayerGameRelation, error) {
	relation := &model.PlayerGameRelation{}

	_, err := repo.client.DB.NewUpdate().Model(relation).ModelTableExpr(tablePlayerGameRelation).Set("status = ?", request.Status).Where("gid = ?", gid).Returning("*").Exec(ctx)

	if err != nil {
		return nil, err
	}

	return relation, nil
}

func (repo *PlayerGameRelationRepository) GetByGameGid(ctx model.Context, gameGid string) ([]*model.PlayerGameRelation, error) {
	var relations []*model.PlayerGameRelation

	err := repo.client.DB.NewSelect().Model(&relations).ModelTableExpr(tablePlayerGameRelation).Where("game_gid = ?", gameGid).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return relations, nil
}
