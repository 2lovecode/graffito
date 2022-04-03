package cache

import (
	"context"
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
	"math/rand"
	"os"
	"time"

	"github.com/willf/bloom"
)

const mapLen = 50
const randLen = 1000

const executeTime = 2
const expireTime = 1

var totalCount = 0
var cacheGetCount = 0
var cacheSetCount = 0
var cacheSpace int64 = 0

var stdTotal = 0

type CacheDb struct {
	data  map[string]string
	count int
}

func (db *CacheDb) Select(key string) string {
	value, ok := db.data[key]

	db.count++

	if ok {
		return value
	}

	return ""
}

func (db *CacheDb) Insert(key string, value string) {
	db.data[key] = value
}

func (db *CacheDb) GetCount() int {
	return db.count
}

func (db *CacheDb) SetCount(count int) {
	db.count = count
}

func NewCacheDb(data map[string]string) *CacheDb {
	return &CacheDb{data, 0}
}

//总访问量，缓存访问量，数据库访问量，缓存空间占用量
func CrossRun() {
	ctx := context.TODO()

	n := uint(1000)
	filter := bloom.New(20*n, 5) // load of 20, 5 keys

	c := initRedis()
	if c == nil {
		os.Exit(2)
	}
	db := initDb(filter)

	//场景模拟

	//缓存穿透
	fmt.Println("缓存穿透:")
	penetrate(ctx, c, db)
	printFormat("缓存穿透示例", totalCount, cacheGetCount+cacheSetCount, db.GetCount(), cacheSpace)
	stdTotal = totalCount

	//缓存穿透 -- 解决方案【全部缓存】
	initData(ctx, c, db)
	penetrateSolutionOne(ctx, c, db)
	printFormat("缓存穿透解决方案【全部缓存】", totalCount, cacheGetCount+cacheSetCount, db.GetCount(), cacheSpace)
	//缓存穿透 -- 解决方案【布隆过滤器】
	initData(ctx, c, db)
	penetrateSolutionTwo(ctx, c, db, filter)
	printFormat("缓存穿透解决方案【布隆过滤器】", totalCount, cacheGetCount+cacheSetCount, db.GetCount(), cacheSpace)

}

func initRedis() *redis2.Client {
	//初始化缓存
	return redis2.NewClient(&redis2.Options{
		Addr: "127.0.0.1:6379",
	})
}

func initDb(filter *bloom.BloomFilter) *CacheDb {
	//数据库写入数据
	db := NewCacheDb(make(map[string]string, mapLen))

	for i := 0; i < mapLen; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		db.Insert(key, value)
		filter.Add([]byte(key))
	}
	return db
}

func initData(ctx context.Context, c *redis2.Client, db *CacheDb) {
	c.FlushDB(ctx)
	totalCount = 0
	cacheGetCount = 0
	cacheSetCount = 0
	db.SetCount(0)
	cacheSpace = 0
}

func setCache(ctx context.Context, key string, value string, c *redis2.Client) {

	res := c.Set(ctx, key, value, expireTime)

	if res.Err() != nil {
		fmt.Println("set redis error")
	}

	res1 := c.DBSize(ctx)

	if res1.Err() != nil {
		fmt.Println("set redis error")
	} else {
		if cacheSpace < res1.Val() {
			cacheSpace = res1.Val()
		}
	}

	cacheSetCount++
}

func printFormat(desc string, totalCount int, cacheCount int, dbCount int, cacheSpace int64) {
	fmt.Println("<<<<", desc)
	if stdTotal > 0 {

	}
	fmt.Printf("<<<<Total visit: %d(%d)| Cache visit: %d(%d) | Db visit: %d(%d) | Cache Space %d(%d)\n", totalCount, stdTotal, cacheCount, stdTotal*cacheCount/totalCount, dbCount, stdTotal*dbCount/totalCount, cacheSpace, stdTotal*int(cacheSpace)/totalCount)
}

func penetrate(ctx context.Context, c *redis2.Client, db *CacheDb) {
	startTime := time.Now().Second()

	for endTime := 0; endTime < executeTime; endTime = time.Now().Second() - startTime {
		totalCount++
		key := fmt.Sprintf("key%d", rand.Intn(randLen))
		value := c.Get(ctx, key)
		cacheGetCount++
		if value.Err() != nil {
			fmt.Println("Error", value.Err())
			break
		}
		if value.Val() != "" {
			continue
		}

		dbVal := db.Select(key)

		if dbVal != "" {
			setCache(ctx, key, dbVal, c)
		}
	}
}

//全部使用缓存
func penetrateSolutionOne(ctx context.Context, c *redis2.Client, db *CacheDb) {
	startTime := time.Now().Second()

	for endTime := 0; endTime < executeTime; endTime = time.Now().Second() - startTime {
		totalCount++
		key := fmt.Sprintf("key%d", rand.Intn(randLen))
		value := c.Get(ctx, key)
		cacheGetCount++
		if value.Err() != nil {
			fmt.Println("Error", value.Err())
			break
		}
		if value.Val() != "" {
			continue
		}

		dbVal := db.Select(key)

		if dbVal != "" {
			setCache(ctx, key, dbVal, c)
		}
	}
}

//使用布隆过滤器
func penetrateSolutionTwo(ctx context.Context, c *redis2.Client, db *CacheDb, filter *bloom.BloomFilter) {
	startTime := time.Now().Second()

	for endTime := 0; endTime < executeTime; endTime = time.Now().Second() - startTime {
		totalCount++

		key := fmt.Sprintf("key%d", rand.Intn(randLen))
		if !filter.Test([]byte(key)) {
			continue
		}
		value := c.Get(ctx, key)
		cacheGetCount++
		if value.Err() != nil {
			fmt.Println("Error", value.Err())
			break
		}
		if value.Val() != "" {
			continue
		}

		dbVal := db.Select(key)

		if dbVal != "" {
			setCache(ctx, key, dbVal, c)
		}
	}
}
