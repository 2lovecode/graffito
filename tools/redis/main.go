package redis

import (
	"bufio"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"graffito/tools/redis/commands"
	"os"
	"strings"

	redis2 "github.com/go-redis/redis/v8"
)

type RedisInstance struct {
	host    string
	port    string
	cmdPool map[string]commands.RedisCommand
}

func NewRedisInstance(host string, port string) *RedisInstance {
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "6379"
	}

	instance := &RedisInstance{
		host: host,
		port: port,
	}
	instance.bootstrap()

	return instance
}

func (instance *RedisInstance) bootstrap() {
	instance.cmdPool = map[string]commands.RedisCommand{
		"set":     &commands.RedisCommandSet{},
		"get":     &commands.RedisCommandGet{},
		"hset":    &commands.RedisCommandHSet{},
		"hget":    &commands.RedisCommandHGet{},
		"hgetall": &commands.RedisCommandHGetAll{},
		"select":  &commands.RedisCommandSelect{},
	}
}

func (instance *RedisInstance) command(cmd string) commands.RedisCommand {
	if c, ok := instance.cmdPool[cmd]; ok {
		return c
	}
	return &commands.RedisCommandNil{}
}

func (instance *RedisInstance) Run() {
	rdb := redis2.NewClient(&redis2.Options{
		Addr: instance.host + ":" + instance.port,
	})

	ctx := context.Background()

	inputReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(":> ")
		input, err := inputReader.ReadString('\n')
		if err == nil {
			if input == "" {
				commands.PrintInfo("请输入命令！")
			} else {
				inputSlice := strings.Split(strings.Trim(input, "\n"), " ")
				cmd := strings.ToLower(inputSlice[0])
				if cmd == "exit" {
					commands.PrintInfo("退出！")
					break
				}
				args := inputSlice[1:]
				command := instance.command(cmd)
				command.Execute(ctx, rdb, args)
			}
		} else {
			commands.PrintError(err.Error())
			continue
		}
	}

}

func NewCommand() *cobra.Command {
	return &cobra.Command{Use: "redis", Run: func(cmd *cobra.Command, args []string) {
		host := ""
		port := ""

		if len(args) > 0 {
			host = args[0]
			if len(args) > 1 {
				port = args[1]
			}
		}
		r := NewRedisInstance(host, port)
		r.Run()
	}, Short: "命令列表", Example: "graffito tools helper"}
}
