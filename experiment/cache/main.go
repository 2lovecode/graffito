package cache

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
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
	penetrate(c, db)
	printFormat("缓存穿透示例", totalCount, cacheGetCount+cacheSetCount, db.GetCount(), cacheSpace)
	stdTotal = totalCount

	//缓存穿透 -- 解决方案【全部缓存】
	initData(c, db)
	penetrateSolutionOne(c, db)
	printFormat("缓存穿透解决方案【全部缓存】", totalCount, cacheGetCount+cacheSetCount, db.GetCount(), cacheSpace)
	//缓存穿透 -- 解决方案【布隆过滤器】
	initData(c, db)
	penetrateSolutionTwo(c, db, filter)
	printFormat("缓存穿透解决方案【布隆过滤器】", totalCount, cacheGetCount+cacheSetCount, db.GetCount(), cacheSpace)

}

func initRedis() redis.Conn {
	//初始化缓存
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("无法连接redis 127.0.0.1:6379 ")
		return nil
	}
	_, err0 := c.Do("FLUSHDB")
	if err0 != nil {
		fmt.Println("flush error")
		return nil
	}
	return c
}

func initDb(filter *bloom.BloomFilter) *CacheDb {
	//数据库写入数据
	db := NewCacheDb(make(map[string]string, mapLen))

	for i := 0; i < mapLen; i++ {
		key := "key" + string(i)
		value := "value" + string(i)
		db.Insert(key, value)
		filter.Add([]byte(key))
	}
	return db
}

func initData(c redis.Conn, db *CacheDb) {
	_, err0 := c.Do("FLUSHDB")
	if err0 != nil {
		fmt.Println("flush error")
	}
	totalCount = 0
	cacheGetCount = 0
	cacheSetCount = 0
	db.SetCount(0)
	cacheSpace = 0
}

func setCache(key string, value string, c redis.Conn) {
	_, err1 := c.Do("SET", key, value)
	if err1 != nil {
		fmt.Println("set redis error")
	}
	sizeDb, err2 := c.Do("DBSIZE")
	if err2 != nil {
		fmt.Println("set redis error")
	} else {
		if cacheSpace < sizeDb.(int64) {
			cacheSpace = sizeDb.(int64)
		}
	}
	_, err3 := c.Do("EXPIRE", key, expireTime)
	if err3 != nil {
		fmt.Println("set redis error")
	}

	cacheSetCount++
}

func printFormat(desc string, totalCount int, cacheCount int, dbCount int, cacheSpace int64) {
	fmt.Println("<<<<", desc)
	if stdTotal > 0 {

	}
	fmt.Printf("<<<<Total visit: %d(%d)| Cache visit: %d(%d) | Db visit: %d(%d) | Cache Space %d(%d)\n", totalCount, stdTotal, cacheCount, stdTotal*cacheCount/totalCount, dbCount, stdTotal*dbCount/totalCount, cacheSpace, stdTotal*int(cacheSpace)/totalCount)
}

func penetrate(c redis.Conn, db *CacheDb) {
	startTime := time.Now().Second()

	for endTime := 0; endTime < executeTime; endTime = time.Now().Second() - startTime {
		totalCount++
		key := "key" + string(rand.Intn(randLen))
		value, err := c.Do("GET", key)
		cacheGetCount++
		if err != nil {
			fmt.Println("Error", err)
			break
		}
		if value != nil {
			continue
		}

		value = db.Select(key)

		if value != "" {
			setCache(key, value.(string), c)
		}
	}
}

//全部使用缓存
func penetrateSolutionOne(c redis.Conn, db *CacheDb) {
	startTime := time.Now().Second()

	for endTime := 0; endTime < executeTime; endTime = time.Now().Second() - startTime {
		totalCount++
		key := "key" + string(rand.Intn(randLen))
		value, err := c.Do("GET", key)
		cacheGetCount++
		if err != nil {
			fmt.Println("Error", err)
			break
		}
		if value != nil {
			continue
		}

		value = db.Select(key)

		setCache(key, value.(string), c)
	}
}

//使用布隆过滤器
func penetrateSolutionTwo(c redis.Conn, db *CacheDb, filter *bloom.BloomFilter) {
	startTime := time.Now().Second()

	for endTime := 0; endTime < executeTime; endTime = time.Now().Second() - startTime {
		totalCount++
		key := "key" + string(rand.Intn(randLen))
		if !filter.Test([]byte(key)) {
			continue
		}
		value, err := c.Do("GET", key)
		cacheGetCount++
		if err != nil {
			fmt.Println("Error", err)
			break
		}
		if value != nil {
			continue
		}

		value = db.Select(key)

		if value != "" {
			setCache(key, value.(string), c)
		}
	}
}
