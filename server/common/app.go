package common

import (
	"context"
	"sync"
)

type App struct {
	Backend       IBackend
	Body          map[string]interface{}
	Session       map[string]string
	Share         Share
	Context       context.Context
	Authorization string
	Languages     []string
	BodyMutex     *sync.Mutex
}
