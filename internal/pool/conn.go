package pool

import (
	"net"
	"sync/atomic"
	"time"
)

type Conn struct {
	usedAt  int64
	netConn net.Conn

	Inited    bool
	pooled    bool
	createdAt time.Time
}

func NewConn(c net.Conn) *Conn {
	cn := &Conn{
		netConn:   c,
		createdAt: time.Now(),
	}
	return cn
}

func (cn *Conn) UsedAt() time.Time {
	t := atomic.LoadInt64(&cn.usedAt)
	return time.Unix(t, 0)
}

func (cn *Conn) SetUseAt(t time.Time) {
	atomic.StoreInt64(&cn.usedAt, t.Unix())
}

func (cn *Conn) Write(b []byte) (int, error) {
	return cn.netConn.Write(b)
}
