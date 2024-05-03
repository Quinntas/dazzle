package main

import (
	"context"
	"fmt"

	sharedRouter "github.com/quinntas/go-rest-template/api/shared/infra/router"
	"github.com/quinntas/go-rest-template/internal/api/database"
	"github.com/quinntas/go-rest-template/internal/api/redis"
	"github.com/quinntas/go-rest-template/internal/api/utils/env"
	"github.com/quinntas/go-rest-template/internal/api/web"
)

func main() {
	envVariables := env.NewEnvVariables()
	ctx := context.WithValue(context.Background(), env.ENV_CTX_KEY, envVariables)

	dbClient := database.NewDatabaseClient(envVariables.DatabaseURL)
	ctx = context.WithValue(ctx, database.DATABASE_CTX_KEY, dbClient)

	redisClient := redis.NewRedisClient(envVariables.RedisURL)
	ctx = context.WithValue(ctx, redis.REDIS_CTX_KEY, redisClient)

	sharedRouter.InitRouters(ctx)

	fmt.Println("[Server] running on", fmt.Sprintf("http://localhost:%s", envVariables.Port))

	web.Serve(envVariables.Port)
}
