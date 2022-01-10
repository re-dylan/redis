package pool

import (
	"context"
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrColsed = errors.New("redis: client is closed")

	ErrorPoolTimeout = errors.New("redis: connection pool timeout")
)

type Pooler interface {
	NewConn(context.Context) (*Conn, error)
	CloseConn(*Conn) error

	Get(context.Context) (*Conn, error)
	Put(context.Context, *Conn)
	Remove(context.Context, *Conn, error)

	Len() int
	IdleLen() int
	Stats() *Stats

	Close() error
}

type Options struct {
	Dialer  func(context.Context) (net.Conn, error)
	OnClose func(*Conn) error

	PoolFIFO           bool
	PoolSize           int
	MinIdleConns       int
	MaxConnAge         time.Duration
	PoolTimeout        time.Duration
	IdleTimeout        time.Duration
	IdleCheckFrequency time.Duration
}

type ConnPool struct {
	opt *Options

	dialErrorNum  uint32
	lastDialError atomic.Value

	queue chan struct{}

	connMu    sync.Mutex
	conns     []*Conn
	idleConns []*Conn

	stats Stats

	_close  uint32
	closeCh chan struct{}
}

var _ Pooler = (*ConnPool)(nil)

func NewConnPool(opt *Options) *ConnPool {
	p := &ConnPool{
		opt: opt,

		queue:     make(chan struct{}, opt.PoolSize),
		conns:     make([]*Conn, 0, opt.PoolSize),
		idleConns: make([]*Conn, 0, opt.PoolSize),
		closeCh:   make(chan struct{}),
	}
	// TODO: check min idle conn

	return p
}

func (p *ConnPool) checkMinIdleConns() {
	if p.opt.MinIdleConns == 0 {
		return
	}

}

func (p *ConnPool) addIdleConn() error {
	//cn, err :=
}

func (p *ConnPool) dialConn(ctx context.Context, pooled bool) (*Conn, error) {
	if p.closed() {
		return nil, ErrColsed
	}

	if atomic.LoadUint32(&p.dialErrorNum) >= uint32(p.opt.PoolSize) {
		err, _ := p.lastDialError.Load().(*)
		return nil,
	}
}

func (p *ConnPool) closed() bool {
	return atomic.LoadUint32(&p._close) == 1
}
