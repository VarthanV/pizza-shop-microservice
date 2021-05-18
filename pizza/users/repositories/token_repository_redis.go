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
	rdb   *redis.Client
	utils utils.UtilityService
}

func NewRedisRepository(rdb *redis.Client, utilityservice utils.UtilityService) models.TokenRepository {
	return &redisrepository{
		rdb:   rdb,
		utils: utilityservice,
	}
}

func (r redisrepository) CreateToken(ctx context.Context, user models.User) (tokenDetails models.TokenDetails, err error) {
	token := models.TokenDetails{}
	accessToken, refreshToken, atExpiresAt, rtExpiresAt, err := r.utils.CreateToken(user)
	if err != nil {
		glog.Error("Unable to create the jwt token...", err)
		return models.TokenDetails{}, err
	}
	atExpirationInUTC := time.Unix(atExpiresAt, 0)
	rtExpirationInUTC := time.Unix(rtExpiresAt, 0)

	token.AccessToken = accessToken
	token.RefreshToken = refreshToken

	// Persist the accessToken and refresh in redis
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

func (r redisrepository) VerifyTokenValidity(ctx context.Context, acToken string) (isValid bool) {
	// If the token exists in the redis store then it is valid type of token needed to be
	// checked will be passed in context
	token := r.rdb.Get(ctx, acToken)
	result, err := token.Result()
	if err != nil {
		glog.Errorf("Unable to get the token %f", err)
		return false
	}
	return result != ""
}
