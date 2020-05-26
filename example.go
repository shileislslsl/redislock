package redislock

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/shileislslsl/redislock"
)

func main() {
	cli = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	rlock := redislock.NewClient(cli) //will panic if the cli is nil
	//lock take three params,lock key,default expiration time,auto refresh key expiration time before unlock the key
	err := rlock.Lock("test_key2", 10*time.Second, true)
	if err == redislock.LOCKFailed { //get lock while other holding the lock
		//Retry lock or return

	}
	if err == redislock.NILClient { //redis client init error.please check the ipadd and auth
		//Renew the redisclient
	}
	// you can add a context
	backCtx := context.Background()
	ctx, cancel := context.WithCancel(backCtx)
	rlock.SetContext(ctx)
	ticker = time.After(5 * time.Second)
	select {
	case <-ticker.C:
		cancel()
	case <-handfuncreturn:
		err = rlock.Unlock()
		if err != nil {
			log.Error(err)
		}
	}
	//check the lock remain time
	leftSecond := rlock.TTL()
}
