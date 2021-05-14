package repositories

import (
	"github.com/VarthanV/pizza/users/models"
	"github.com/go-redis/redis/v8"
)

type redisrepository struct {
	rdb *redis.Client
}

func NewRedisRepository(rdb *redis.Client) models.TokenRepository {
	return &redisrepository{
		rdb: rdb,
	}
}

func (r redisrepository) CreateToken(user models.User) (tokenDetails models.TokenDetails, err error) {

	return models.TokenDetails{}, nil
}

func (r redisrepository) VerifyTokenValidity(token models.TokenDetails) (isValid bool) {

	return false
}
