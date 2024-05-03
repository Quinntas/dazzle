package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/quinntas/go-rest-template/internal/api/redis"
)

type DatabaseClient struct {
	client *sqlx.DB
}

func GetDBClientFromContext(ctx context.Context) DatabaseClient {
	return ctx.Value(DATABASE_CTX_KEY).(DatabaseClient)
}

func NewDatabaseClient(databaseUrl string) DatabaseClient {
	db := sqlx.MustConnect("mysql", databaseUrl)
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return DatabaseClient{
		client: db,
	}
}

func (db *DatabaseClient) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.client.Exec(query, args...)
}

func (db *DatabaseClient) MustExec(query string, args ...interface{}) sql.Result {
	return db.client.MustExec(query, args...)
}

// TODO: refactor this monster
func QueryRowWithCache[T interface{}](
	db *DatabaseClient,
	redis *redis.RedisClient,
	key string,
	duration time.Duration,
	query string,
	args ...interface{},
) (*T, error) {
	redisResult, err := redis.Get(key)
	if err != nil {
		if err.Error() != "redis: nil" {
			return nil, err
		}
	}
	if redisResult != "" {
		var result *T
		err := json.Unmarshal([]byte(redisResult), &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	var result T
	err = db.client.QueryRowx(query, args...).StructScan(&result)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	err = redis.Set(key, string(bytes), duration)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func Query[T interface{}](db *DatabaseClient, query string, args ...interface{}) ([]T, error) {
	rows, err := db.client.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]T, 0)
	for rows.Next() {
		var row T
		err = rows.StructScan(&row)
		if err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	return results, nil
}

func QueryRow[T interface{}](db *DatabaseClient, query string, args ...interface{}) (*T, error) {
	var result T
	err := db.client.QueryRowx(query, args...).StructScan(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
