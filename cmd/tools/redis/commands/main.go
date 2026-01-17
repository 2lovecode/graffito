package commands

import (
	"context"
	"fmt"
	"strconv"
	"time"

	redis2 "github.com/go-redis/redis/v8"
)

type RedisCommand interface {
	Execute(ctx context.Context, client *redis2.Client, args []string)
}

type RedisCommandNil struct {
}

func (cmd *RedisCommandNil) Execute(ctx context.Context, client *redis2.Client, args []string) {
	PrintError("无效的命令！")
}

type RedisCommandSet struct {
}

func (cmd *RedisCommandSet) Execute(ctx context.Context, client *redis2.Client, args []string) {
	expiration := 0 * time.Second
	if len(args) > 1 {
		if len(args) > 2 {
			expirationInt, _ := strconv.Atoi(args[2])
			expiration = time.Duration(expirationInt) * time.Second
		}
		res := client.Set(ctx, args[0], args[1], expiration)
		str, er := res.Result()
		if er == nil {
			PrintInfo(str)
		} else {
			PrintError(er.Error())
		}
	} else {
		PrintError("至少需要两个参数！")
	}
}

type RedisCommandGet struct {
}

func (cmd *RedisCommandGet) Execute(ctx context.Context, client *redis2.Client, args []string) {
	if len(args) > 0 {
		res := client.Get(ctx, args[0])
		if res != nil {
			str, er := res.Result()
			if er == nil {
				PrintInfo(str)
			} else if er == redis2.Nil {
				PrintInfo("nil")
			} else {
				PrintError(er.Error())
			}
		}
	}
}

type RedisCommandHSet struct {
}

func (cmd *RedisCommandHSet) Execute(ctx context.Context, client *redis2.Client, args []string) {
	if len(args) > 2 {
		res := client.HSet(ctx, args[0], args[1:])
		if res != nil {
			i, er := res.Result()
			if er == nil {
				fmt.Println(i)
			} else if er == redis2.Nil {
				PrintInfo("nil")
			} else {
				PrintError(er.Error())
			}
		}
	}
}

type RedisCommandHGet struct {
}

func (cmd *RedisCommandHGet) Execute(ctx context.Context, client *redis2.Client, args []string) {
	if len(args) > 1 {
		res := client.HGet(ctx, args[0], args[1])
		if res != nil {
			str, er := res.Result()
			if er == nil {
				PrintInfo(str)
			} else if er == redis2.Nil {
				PrintInfo("nil")
			} else {
				PrintError(er.Error())
			}
		}
	}
}

type RedisCommandHGetAll struct {
}

func (cmd *RedisCommandHGetAll) Execute(ctx context.Context, client *redis2.Client, args []string) {
	if len(args) > 0 {
		res := client.HGetAll(ctx, args[0])
		if res != nil {
			strmap, er := res.Result()
			if er == nil {
				fmt.Println(strmap)
			} else if er == redis2.Nil {
				PrintInfo("nil")
			} else {
				PrintError(er.Error())
			}
		}
	}
}

func PrintInfo(msg string) {
	fmt.Println(msg)
}

func PrintError(msg string) {
	fmt.Println("错误: ", msg)
}

type RedisCommandSelect struct {
}

func (cmd *RedisCommandSelect) Execute(ctx context.Context, client *redis2.Client, args []string) {
	if len(args) > 0 {
		index, _ := strconv.Atoi(args[0])
		fmt.Println(index)
		pipe := client.Pipeline()
		pipe.Do(ctx, "select", index)
		_, err := pipe.Exec(ctx)

		if err == nil {
			PrintInfo("OK")
		} else {
			PrintError(err.Error())
		}
	}
}
