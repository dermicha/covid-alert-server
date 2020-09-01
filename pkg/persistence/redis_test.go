package persistence

import (
	"fmt"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestGenerateNonce(t *testing.T) {
	s, _ := miniredis.Run()
	defer s.Close()
	opt, _ := redis.ParseURL("redis://" + s.Addr() + "/0")
	conn := DialRedis(opt)

	response, _ := conn.GenerateNonce()

	assert.Equal(t, 32, len(response), "Nonce should be 32 chars")
}

func TestGenerateNonceError(t *testing.T) {
	s, _ := miniredis.Run()
	defer s.Close()
	opt, _ := redis.ParseURL("redis://x.y.z.y/1")

	conn := DialRedis(opt)

	response, _ := conn.GenerateNonce()
	fmt.Println(response)

	assert.Equal(t, 32, len(response), "Nonce should be 32 chars")
}
