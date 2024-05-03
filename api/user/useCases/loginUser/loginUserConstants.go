package loginUser

import "time"

const (
	TokenExpirationTime = time.Second * 3600 // 1 hr
	TokenRedisKey       = "login:user:"
)
