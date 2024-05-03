package userRepo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	userDomain "github.com/quinntas/go-rest-template/api/user/domain"
	userDatabase "github.com/quinntas/go-rest-template/api/user/infra/database"
	userMapper "github.com/quinntas/go-rest-template/api/user/mapper"
	userModel "github.com/quinntas/go-rest-template/api/user/model"
	"github.com/quinntas/go-rest-template/internal/api/database"
	"github.com/quinntas/go-rest-template/internal/api/redis"
)

func Create(user userDomain.UserDomain, ctx context.Context) (sql.Result, error) {
	db := database.GetDBClientFromContext(ctx)

	return db.Exec(fmt.Sprintf("INSERT INTO %s (pid, email, password, roleId) VALUES (?, ?, ?, ?)", userDatabase.TABLE_NAME),
		user.PID.Value,
		user.Email.Value,
		user.Password.Value,
		user.RoleId.Value,
	)
}

func FindWithEmail(email string, ctx context.Context) (userDomain.UserDomain, error) {
	db := database.GetDBClientFromContext(ctx)
	redis := redis.GetRedisClientFromContext(ctx)

	key := fmt.Sprintf("user:FindWithEmail:%s", email)
	duration := time.Minute * 30

	userModel, err := database.QueryRowWithCache[userModel.UserModel](
		&db,
		&redis,
		key,
		duration,
		fmt.Sprintf("SELECT * FROM %s WHERE email=?",
			userDatabase.TABLE_NAME),
		email,
	)
	if err != nil {
		return userDomain.UserDomain{}, err
	}

	return userMapper.ToDomain(*userModel), nil
}

func FindWithId(id int, ctx context.Context) (userDomain.UserDomain, error) {
	db := database.GetDBClientFromContext(ctx)

	userModel, err := database.QueryRow[userModel.UserModel](&db, fmt.Sprintf("SELECT * FROM %s WHERE id=?", userDatabase.TABLE_NAME), id)
	if err != nil {
		return userDomain.UserDomain{}, err
	}

	return userMapper.ToDomain(*userModel), nil
}
