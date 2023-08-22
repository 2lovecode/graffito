package redis

import (
	"context"
	"sync"
	"time"
)

type Server struct {
	data   sync.Map
	expire sync.Map
}

func NewServer() *Server {
	return &Server{
		data:   sync.Map{},
		expire: sync.Map{},
	}
}

func (s *Server) Start() {
	go func() {
		for {
			needDel := make([]string, 0)
			timeNow := time.Now()
			s.expire.Range(func(key, value any) bool {
				if v, ok := value.(int64); ok {
					vt := time.Unix(v/1000, 0)
					if vt.Before(timeNow) {
						needDel = append(needDel, key.(string))
					}
				}
				return true
			})
			if len(needDel) > 0 {
				for _, v := range needDel {
					s.expire.Delete(v)
					s.data.Delete(v)
				}
			}
		}
	}()
}

type Client struct {
	server *Server
}

func NewClient() *Client {
	s := NewServer()
	s.Start()
	return &Client{
		server: s,
	}
}

func (cli *Client) Set(ctx context.Context, key string, val string) {
	cli.server.data.Store(key, val)
}

func (cli *Client) Get(ctx context.Context, key string) (val string) {
	if v, ok := cli.server.data.Load(key); ok {
		val = v.(string)
	}
	return
}

func (cli *Client) Expire(ctx context.Context, key string, expire time.Duration) {
	cli.server.expire.Store(key, expire.Milliseconds())
}
