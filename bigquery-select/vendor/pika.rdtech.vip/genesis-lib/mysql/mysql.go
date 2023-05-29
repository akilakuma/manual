package mysql

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Orm struct {
	db *gorm.DB
}

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Scope struct {
	Scope *gorm.Scope
}

type Association struct {
	Association *gorm.Association
}

func Open(addr string, opt ...Option) (*Orm, error) {
	db, err := gorm.Open("mysql", addr)
	if err != nil {
		return nil, err
	}

	for _, o := range opt {
		o.apply(db)
	}

	return &Orm{db: db}, nil
}

func (o *Orm) Close() (err error) {
	err = o.db.Close()
	return
}

func (o *Orm) Ping() (err error) {
	err = o.db.DB().Ping()
	return
}

func (o *Orm) AutoMigrate(value ...interface{}) *Orm {
	return &Orm{db: o.db.AutoMigrate(value...)}
}

func (o *Orm) AddError(err error) error {
	return o.db.AddError(err)
}

func (o *Orm) Error() error {
	return o.db.Error
}

func (o *Orm) RowsAffected() int64 {
	return o.db.RowsAffected
}
