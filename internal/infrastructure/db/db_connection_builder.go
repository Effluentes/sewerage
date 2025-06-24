package db

import (
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
	"your_project/internal/domain/repositories"
)

type (
	DBBuilder struct {
		postgresConnStr   string
		timescaleConnStr   string
		redisAddr         string
		redisPassword     string
		redisDB           int
	}

	builtConnections struct {
		Postgres repositories.PostgresDatabase
		Timescale repositories.TimescaleDatabase
		Redis    repositories.RedisDatabase
	}
)

func NewDBBuilder() *DBBuilder {
	return &DBBuilder{}
}

func (b *DBBuilder) WithPostgres(connStr string) *DBBuilder {
	b.postgresConnStr = connStr
	return b  // Fluent interface
}

func (b *DBBuilder) WithTimescale(connStr string) *DBBuilder {
	b.timescaleConnStr = connStr
	return b
}

func (b *DBBuilder) WithRedis(addr, password string, db int) *DBBuilder {
	b.redisAddr = addr
	b.redisPassword = password
	b.redisDB = db
	return b
}

func (b *DBBuilder) Build() (*builtConnections, error) {
	var err error
	conns := &builtConnections{}

	// Inicjalizacja PostgreSQL
	if b.postgresConnStr != "" {
		pgDB, pgErr := initPostgres(b.postgresConnStr)
		if pgErr != nil {
			err = fmt.Errorf("postgres init failed: %w", pgErr)
		}
		conns.Postgres = pgDB
	}

	// Inicjalizacja Timescale
	if b.timescaleConnStr != "" {
		tsDB, tsErr := initTimescale(b.timescaleConnStr)
		if tsErr != nil {
			err = fmt.Errorf("timescale init failed: %w", tsErr)
		}
		conns.Timescale = tsDB
	}

	// Inicjalizacja Redis
	if b.redisAddr != "" {
		rdb, redisErr := initRedis(b.redisAddr, b.redisPassword, b.redisDB)
		if redisErr != nil {
			err = fmt.Errorf("redis init failed: %w", redisErr)
		}
		conns.Redis = rdb
	}

	return conns, err
}

// Prywatne funkcje inicjalizujące (implementacje ukryte)
// ----------------------------------------------------------------------
func initPostgres(connStr string) (*postgresImpl, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &postgresImpl{db: db}, nil
}

func initTimescale(connStr string) (*timescaleImpl, error) {
	// Timescale używa drivera PostgreSQL, ale może mieć inną konfigurację
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &timescaleImpl{db: db}, nil
}

func initRedis(addr, password string, db int) (*redisImpl, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return &redisImpl{client: client}, nil
}

// Prywatne implementacje (ukryte)
// ----------------------------------------------------------------------
type postgresImpl struct {
	db *sql.DB
}

func (p *postgresImpl) Ping() error {
	return p.db.Ping()
}

func (p *postgresImpl) GetDB() interface{} {
	return p.db
}

type timescaleImpl struct {
	db *sql.DB
}

func (t *timescaleImpl) Ping() error {
	return t.db.Ping()
}

func (t *timescaleImpl) GetDB() interface{} {
	return t.db
}

type redisImpl struct {
	client *redis.Client
}

func (r *redisImpl) Ping() error {
	_, err := r.client.Ping(context.Background()).Result()
	return err
}

func (r *redisImpl) GetDB() interface{} {
	return r.client
}