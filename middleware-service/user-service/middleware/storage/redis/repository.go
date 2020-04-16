package redis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"user-service/core/model"
	"user-service/core/repo"
	"user-service/middleware/business"
)

type repository struct {
	client *redis.Client
}

func newRedisClient(redisURL string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewRedisRepository(redisURL string) (repo.UserRepository, error) {
	rep := &repository{}
	client, err := newRedisClient(redisURL)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRedisRepository")
	}
	rep.client = client
	return rep, nil
}

func (r *repository) generateKey(code  string) string {
	return fmt.Sprintf("user[%s]", code)
}

func (r *repository) FetchOne(code string) (*model.User, error) {
	key := r.generateKey(code)
	data, err := r.client.Get(key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "repository.User.Fetch")
	}
	if data == nil {
		return nil, errors.Wrap(business.ErrUserNotFound, "repository.User.Fetch")
	}
	user := &model.User{}
	if err = json.Unmarshal(data, user); err != nil{
		return nil, err
	}
	return user, nil
}

func (r *repository) Fetch(codes []string) ([]*model.User, error) {
	keys := funk.Map(codes, r.generateKey).([]string)
	data, err := r.client.MGet(keys...).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.User.Fetch")
	}
	if len(data) == 0 {
		return nil, errors.Wrap(business.ErrUserNotFound, "repository.User.Fetch")
	}
	users := funk.Map(data, func(val interface{}) (s *model.User) {
		json.Unmarshal([]byte(val.(string)), s)
		return
	} ).([]*model.User)
	return users, nil
}

func (r *repository) FetchAll() ([]*model.User, error) {
	keys := r.client.Keys("user*").Val()
	if len(keys) == 0 {
		return nil, errors.Wrap(business.ErrUserNotFound, "repository.User.FetchAll")
	}
	data, err := r.client.MGet(keys...).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchAll")
	}
	if len(data) == 0 {
		return nil, errors.Wrap(business.ErrUserNotFound, "repository.User.FetchAll")
	}
	users := funk.Map(data, func(val interface{}) (s *model.User) {
		json.Unmarshal([]byte(val.(string)), s)
		return
	} ).([]*model.User)
	return users, nil
}

func (r *repository) FetchQuery(query interface{}) ([]*model.User, error) {
	return r.FetchAll() //todo querying
}

func (r *repository) Store(user *model.User) error {
	key := r.generateKey(user.UniqueId)
	data, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "repository.User.Store")
	}
	if err = r.client.MSet(key, data).Err(); err != nil {
		return errors.Wrap(err, "repository.User.Store")
	}
	return nil
}

func (r *repository) Modify(user *model.User) error {
	if err := r.Store(user); err != nil {
		return errors.Wrap(err, "repository.User.Modify")
	}
	return nil
}

func (r *repository) Replace(user *model.User) error {
	if err := r.Store(user); err != nil {
		return errors.Wrap(err, "repository.User.Replace")
	}
	return nil
}

func (r *repository) Remove(code string) error {
	key := r.generateKey(code)
	if err := r.client.Del(key).Err(); err != nil {
		return errors.Wrap(err, "repository.User.Remove")
	}
	return nil
}

func (r *repository) RemoveObj(user *model.User) error {
	if err := r.Remove(user.UniqueId); err != nil {
		return errors.Wrap(err, "repository.User.RemoveObj")
	}
	return nil
}

func (r *repository) Count(query interface{}) (int64, error) {
	keys := r.client.Keys("user*").Val()
	return int64(len(keys)), nil
}

