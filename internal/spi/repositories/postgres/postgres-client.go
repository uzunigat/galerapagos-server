package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/oiime/logrusbun"
	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type BunPostgresDatabaseClient struct {
	DB     *bun.DB
	config *config.DbConfig
}

func (client *BunPostgresDatabaseClient) getPostgresURL() string {
	return "postgres://" + client.config.Username + ":" + url.QueryEscape(client.config.Password) + "@" + client.config.Host + ":" + client.config.Port + "/" + client.config.Name + "?sslmode=disable"
}

func NewBunPostgresDatabaseClient(config *config.DbConfig) *BunPostgresDatabaseClient {
	client := &BunPostgresDatabaseClient{}
	client.config = config
	client.Connect()
	log.Info().Msg("Database client initialized.")
	return client
}

func (client *BunPostgresDatabaseClient) Connect() error {
	connectionString := client.getPostgresURL()
	log.Info().Msgf("Connecting to database: %s", connectionString)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connectionString)))
	client.DB = bun.NewDB(sqldb, pgdialect.New())
	log := logrus.New()
	client.DB.AddQueryHook(logrusbun.NewQueryHook(logrusbun.QueryHookOptions{Logger: log}))
	err := client.DB.Ping()
	return err
}

func (client *BunPostgresDatabaseClient) getMigrations() (*migrate.Migrate, error) {
	log.Info().Msg("Running migrations.")
	var foundDir string
	dirs := []string{"./migrations", "../migrations", "../../migrations", "../../../migrations", "/migrations"}

	for _, dir := range dirs {
		_, err := os.Stat(dir)
		if err != nil {
			log.Info().Msg(fmt.Sprintf("Migration folder: %s doesn't exist.", dir))
		} else {
			foundDir = dir
			break
		}
	}

	m, err := migrate.New(fmt.Sprintf("file://%s", foundDir), client.getPostgresURL())

	return m, err
}

func (client *BunPostgresDatabaseClient) MigrateUp() {
	m, err := client.getMigrations()
	if err != nil {
		log.Panic().Err(err).Msg("Failed to initialize migrations")
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Panic().Err(err).Msg("Failed to apply up migrations")
	}
}

func (client *BunPostgresDatabaseClient) MigrateDown() {
	m, err := client.getMigrations()
	if err != nil {
		log.Panic().Err(err).Msg("Failed to initialize migrations")
	}

	err = m.Down()
	log.Info().Msg("MigrateDown: Applying migration")
	if err != nil && err != migrate.ErrNoChange {
		log.Panic().Err(err).Msg("Failed to apply down migrations")
	}
}

func (client *BunPostgresDatabaseClient) IsConnected() (bool, error) {
	err := client.DB.Ping()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (client *BunPostgresDatabaseClient) Close(ctx context.Context) {
	client.DB.Close()
}
