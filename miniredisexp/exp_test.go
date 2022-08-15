package miniredisexp

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
)

func TestSomething(t *testing.T) {
	s := miniredis.RunT(t)

	// Optionally set some keys your code expects:
	s.Set("foo", "bar")
	s.HSet("some", "other", "key")

	// Run your code and see if it behaves.
	// An example using the redigo library from "github.com/gomodule/redigo/redis":
	c, _ := redis.Dial("tcp", s.Addr())
	_, _ = c.Do("SET", "foo", "bar")

	// Optionally check values in redis...
	if got, err := s.Get("foo"); err != nil || got != "bar" {
		t.Error("'foo' has the wrong value")
	}
	// ... or use a helper for that:
	s.CheckGet(t, "foo", "bar")

	// TTL and expiration:
	s.Set("foo", "bar")
	s.SetTTL("foo", 10*time.Second)
	s.FastForward(11 * time.Second)
	if s.Exists("foo") {
		t.Fatal("'foo' should not have existed anymore")
	}
}
