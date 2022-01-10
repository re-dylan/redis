package pool

type Stats struct {
	Hits     uint32
	Misses   uint32
	Timeouts uint32

	TotalConns uint32
	IdleConns  uint32
	StaleConns uint32
}
