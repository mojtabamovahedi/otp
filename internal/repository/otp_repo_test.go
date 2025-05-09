package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mojtabamovahedi/otp/internal/repository/types"
	"github.com/mojtabamovahedi/otp/pkg/redis"
	"github.com/stretchr/testify/suite"

	"testing"

	rdis "github.com/redis/go-redis/v9"
)

type RedisConnTest struct {
	r rdis.Cmdable
}

func (r *RedisConnTest) Set(ctx context.Context, key string, value []byte) error {
	return r.r.Set(ctx, key, string(value), 5*time.Minute).Err()
}

func (r *RedisConnTest) Get(ctx context.Context, key string) ([]byte, error) {
	get, err := r.r.Get(ctx, key).Bytes()
	return get, err
}

func (r *RedisConnTest) Del(ctx context.Context, key string) error {
	return r.r.Del(ctx, key).Err()
}

// OtpRepoTestSuite defines the test suite for OtpRepo
type OtpRepoTestSuite struct {
	suite.Suite
	ctx  context.Context
	oc   *redis.ObjectCacher[types.OTP]
	r    rdis.Cmdable
	repo OtpRepo
}

// SetupTest initializes the test suite
func (s *OtpRepoTestSuite) SetupTest() {
	s.r = redisClient()
	s.ctx = context.Background()
	oc := redis.NewObjectCacher[types.OTP](&RedisConnTest{r: s.r})
	s.repo = NewOtpRepo(oc)
}

func redisClient() rdis.Cmdable {
	return rdis.NewClient(&rdis.Options{
		Addr:     "localhost:6379", // Update with your address
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
}

func (s *OtpRepoTestSuite) TearDownTest() {
	s.r.Shutdown(s.ctx)
}

// TODO: new more tests

func (s *OtpRepoTestSuite) TestAddGenerateOTP() {
	var tests = []struct {
		phone       string
		otp         string
		delay       time.Duration
		exceptedErr error
	}{
		{phone: "1234567890", delay: 1 * time.Second, exceptedErr: nil},
		{phone: "1234567890", delay: 1 * time.Second, exceptedErr: nil},
		{phone: "1234567891", delay: 1 * time.Second, exceptedErr: nil},
		{phone: "1234567892", delay: 6 * time.Minute, exceptedErr: ErrOtpNotFound},
	}

	for i, tt := range tests {
		log.Printf("start %d test", i+1)
		err := s.repo.Generate(s.ctx, tt.phone)
		s.Nil(err)

		get, err := s.r.Get(s.ctx, redis.CreateKey(tt.phone)).Bytes()
		s.Nil(err)

		var out types.OTP
		err = s.oc.Unmarshal(get, &out)
		s.Nil(err)

		log.Printf("waiting in test for %v", tt.delay)
		time.Sleep(tt.delay)

		_, err = s.repo.Verify(s.ctx, tt.phone, out.Code)
		s.Equal(tt.exceptedErr, err, fmt.Sprintf("error should be %v", err))
		log.Printf("ended %d test", i+1)
	}

}

func TestOtpRepo(t *testing.T) {
	suite.Run(t, new(OtpRepoTestSuite))
}
