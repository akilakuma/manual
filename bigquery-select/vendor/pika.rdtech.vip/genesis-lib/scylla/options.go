package scylla

import (
	"time"

	"github.com/syhlion/gocql"
)

type Option interface {
	apply(*gocql.ClusterConfig)
}

type funcOption struct {
	f func(*gocql.ClusterConfig)
}

func (fo *funcOption) apply(c *gocql.ClusterConfig) {
	fo.f(c)
}

func newFuncOption(f func(*gocql.ClusterConfig)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func WithTimeout(timeout time.Duration) Option {
	return newFuncOption(func(c *gocql.ClusterConfig) {
		c.Timeout = timeout
	})
}

func WithConnectTimeout(connecttimeout time.Duration) Option {
	return newFuncOption(func(c *gocql.ClusterConfig) {
		c.ConnectTimeout = connecttimeout
	})
}

func WithPort(port int) Option {
	return newFuncOption(func(c *gocql.ClusterConfig) {
		c.Port = port
	})
}

func WithNumConns(numconns int) Option {
	return newFuncOption(func(c *gocql.ClusterConfig) {
		c.NumConns = numconns
	})
}

func WithConsistency(consistency Consistency) Option {
	return newFuncOption(func(c *gocql.ClusterConfig) {
		c.Consistency = gocql.Consistency(consistency)
	})
}

func WithRetryPolicy(retrypolicy gocql.RetryPolicy) Option {
	return newFuncOption(func(c *gocql.ClusterConfig) {
		c.RetryPolicy = retrypolicy
	})
}

func WithMaxPreparedStmts(maxPreparedStmts int) Option {
	return newFuncOption(func(c *gocql.ClusterConfig) {
		c.MaxPreparedStmts = maxPreparedStmts
	})
}

type Consistency uint16

const (
	Any         Consistency = 0x00
	One         Consistency = 0x01
	Two         Consistency = 0x02
	Three       Consistency = 0x03
	Quorum      Consistency = 0x04
	All         Consistency = 0x05
	LocalQuorum Consistency = 0x06
	EachQuorum  Consistency = 0x07
	LocalOne    Consistency = 0x0A
)
