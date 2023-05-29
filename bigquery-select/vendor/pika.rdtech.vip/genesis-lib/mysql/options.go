package mysql

import (
	"time"

	"pika.rdtech.vip/eden-lib/gormlog"
	"github.com/jinzhu/gorm"
)

type Option interface {
	apply(*gorm.DB)
}

type funcOption struct {
	f func(*gorm.DB)
}

func (fo *funcOption) apply(c *gorm.DB) {
	fo.f(c)
}

func newFuncOption(f func(*gorm.DB)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func WithSetMaxIdleConns(conns int) Option {
	return newFuncOption(func(c *gorm.DB) {
		c.DB().SetMaxIdleConns(conns)
	})
}

func WithSetMaxOpenConns(conns int) Option {
	return newFuncOption(func(c *gorm.DB) {
		c.DB().SetMaxOpenConns(conns)
	})
}

func WithSetConnMaxLifetime(sec int) Option {
	return newFuncOption(func(c *gorm.DB) {
		c.DB().SetConnMaxLifetime(time.Duration(sec) * time.Second)
	})
}

func WithSetLogger(log *gormlog.Logger) Option {
	return newFuncOption(func(c *gorm.DB) {
		c.SetLogger(log)
	})
}

func WithSingularTable(status bool) Option {
	return newFuncOption(func(c *gorm.DB) {
		c.SingularTable(status)
	})
}

func WithLogMode(status bool) Option {
	return newFuncOption(func(c *gorm.DB) {
		c.LogMode(status)
	})
}
