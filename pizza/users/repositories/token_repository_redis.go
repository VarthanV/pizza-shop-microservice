package repositories

import (
	"context"
	"time"

	"github.com/VarthanV/pizza/users/models"
	"github.com/VarthanV/pizza/users/utils"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
)

type redisrepository struct {
	rdb *redis.Client
}

func NewRedisRepository(rdb *redis.Client) models.TokenRepository {
	return &redisrepository{
		rdb: rdb,
	}
}

func (r redisrepository) CreateToken(ctx context.Context, user models.User) (tokenDetails models.TokenDetails, err error) {
	token := models.TokenDetails{}
	accessToken, refreshToken, atExpiresAt, rtExpiresAt, err := utils.CreateToken(user)
	if err != nil {
		glog.Error("Unable to create the jwt token...", err)
		return models.TokenDetails{}, err
	}
	atExpirationInUTC := time.Unix(atExpiresAt, 0)
	rtExpirationInUTC := time.Unix(rtExpiresAt, 0)

	token.AccessToken = accessToken
	token.RefreshToken = refreshToken
	// Persist the accessToken in redis
	go func() {
		now := time.Now()
		err = r.rdb.Set(ctx, token.AccessToken, user.ID, atExpirationInUTC.Sub(now)).Err()
		if err != nil {
			glog.Errorf("Unable to set the accessToken in Redis  key store .. %s ", err)
		}
		err = r.rdb.Set(ctx, token.RefreshToken, user.ID, rtExpirationInUTC.Sub(now)).Err()
		if err != nil {
			glog.Errorf("Unable to set the refreshToken in the Redis key store..%s", err)
		}
	}()

	return token, nil
}

func (r redisrepository) VerifyTokenValidity(token models.TokenDetails) (isValid bool) {

	return false
}
