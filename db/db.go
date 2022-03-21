package db

import "context"

type Db interface {
	Connect(uri string) (interface{}, context.Context, context.CancelFunc, error)
	Close() (interface{}, context.Context, context.CancelFunc)
	Ping(interface{}, context.Context) error
}
