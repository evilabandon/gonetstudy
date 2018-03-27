package connect

import (
	"github.com/silenceper/pool"
	"fmt"
	"time"
)

func InitThread(min,max int,factory func() (interface{},error),close func(v interface{}) error) (pool.Pool,error) {
	poolConfig := &pool.PoolConfig{
		InitialCap:min,
		MaxCap:max,
		Factory: factory,
		Close:close,
		IdleTimeout:15*time.Second,
	}
	p, err := pool.NewChannelPool(poolConfig)
	if err!=nil {
		fmt.Println("Init err=",err)
		return nil,err
	}
	return p,nil
}
