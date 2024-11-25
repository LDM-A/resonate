package natsio

import (
	"time"

	"github.com/resonatehq/resonate/internal/aio"
)

type Config struct {
	Size    int
	Workers int
	Timeout time.Duration
}

type Nats struct {
	sq      chan *aio.Message
	workers []*NatsWorker
}

type NatsWorker struct {
	//sq <-chan *aio.Message
}
