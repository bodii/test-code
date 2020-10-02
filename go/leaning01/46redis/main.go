package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func init() {
	initClient()
}

// init connect
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}

func redisGetAndSetDeom() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	scoreVal, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score:", scoreVal)

	notExistsVal, err := rdb.Get("name").Result()

	if err == redis.Nil {
		fmt.Println("name does not exists")
		return
	} else if err != nil {
		fmt.Printf("get name failed, err: %v\n", err)
		return
	}
	fmt.Println("name:", notExistsVal)

}

func redisZsetDemo() {
	zsetKey := "language_rank"
	language := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "JavaScript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
	}

	// zadd
	num, err := rdb.ZAdd(zsetKey, language...).Result()
	if err != nil {
		fmt.Printf("zadd failed. err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d success.\n", num)

	// setting Golang memeber increment 10 score
	goLangScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zIncrBy golang member score failed. err:%v\n", err)
		return
	}
	fmt.Printf("golang's socre is %f\n", goLangScore)

	// get top3
	ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrangewithscore failed, err:%v\n", err)
		return
	}
	for _, revScore := range ret {
		fmt.Println(revScore.Member, revScore.Score)
	}

	// get top 95 - 100
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}

	zRangeByScoreRet, err := rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}

	for _, z := range zRangeByScoreRet {
		fmt.Println(z.Member, z.Score)
	}
}

func redisPiplineDemo() {
	pipe := rdb.Pipeline()

	incr := pipe.Incr("pipeline_counter")
	pipe.Expire("pipeline_counter", time.Hour)

	_, err := pipe.Exec()
	fmt.Println(incr.Val(), err)
}

func redisPipelinedDemo() {
	var incr *redis.IntCmd
	_, err := rdb.Pipelined(func(pipe redis.Pipeliner) error {
		incr = pipe.Incr("pipelined_counter")
		pipe.Expire("pipelined_counter", time.Hour)
		return nil
	})
	fmt.Println(incr.Val(), err)
}

func redisTransactionDemo() {
	pipe := rdb.TxPipeline()
	incr := pipe.Incr("tx_pipeline_counter")
	pipe.Expire("tx_pipeline_counter", time.Hour)

	_, err := pipe.Exec()
	fmt.Println(incr.Val(), err)
}

func redisTransactionPipelinedDemo() {
	var incr *redis.IntCmd
	_, err := rdb.TxPipelined(func(pipe redis.Pipeliner) error {
		incr = pipe.Incr("tx_pipelined_counter")
		pipe.Expire("tx_pipelined_counter", time.Hour)
		return nil
	})

	fmt.Println(incr.Val(), err)
}

func redisWatchTransactionDemo() {
	key := "watch_count"
	watchErr := rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			pipe.Set(key, n+1, 0)
			return nil
		})
		return err
	}, key)

	fmt.Println(watchErr)
}

func redisTransactionGetAndSetDemo() {
	routineCount := 100
	increment := func(key string) error {
		txf := func(tx *redis.Tx) error {
			n, err := tx.Get(key).Int()
			if err != nil && err != redis.Nil {
				return err
			}

			n++

			_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
				pipe.Set(key, n, 0)
				return nil
			})
			return err
		}

		for retries := routineCount; retries > 0; retries-- {
			err := rdb.Watch(txf, key)
			if err != redis.TxFailedErr {
				return err
			}
		}

		return errors.New("increment reached maximum number of retries")
	}

	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		go func() {
			defer wg.Done()

			if err := increment("counter3"); err != nil {
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()

	n, err := rdb.Get("counter3").Int()
	fmt.Println("ended with", n, err)

}

func main() {
	// redisGetAndSetDeom()
	// redisZsetDemo()
	// redisPiplineDemo()
	// redisPipelinedDemo()
	// redisTransactionDemo()
	// redisTransactionPipelinedDemo()
	// redisWatchTransactionDemo()
	redisTransactionGetAndSetDemo()

}
