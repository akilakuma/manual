package snowflake

import (
	"fmt"

	"github.com/sony/sonyflake"
)

type node struct {
	node *sonyflake.Sonyflake
}

type id struct {
	id uint64
}

//Node 封裝結構
type Node interface {
	Generate() ID
}

//Setting 初始設定
type Setting struct {
	sonySet sonyflake.Settings
}

//ID 封裝結構
type ID interface {
	Int64() int64
	String() string
	Bytes() []byte
}

//NewNode 產生編號生產節點
func NewNode(opts ...Option) *node {
	var set Setting
	for _, o := range opts {
		o.apply(&set)
	}

	return &node{node: sonyflake.NewSonyflake(set.sonySet)}
}

//Generate 產生 UniqueID
func (s *node) Generate() ID {
	sony, err := s.node.NextID()
	if err != nil {
		fmt.Println("Generate NextID err:", err.Error())
		return &id{id: 0}
	}

	return &id{id: sony}
}

// Int64 returns an int64 of the snowflake ID
func (s *id) Int64() int64 {
	return int64(s.id)
}

// String returns an String of the snowflake ID
func (s *id) String() string {
	return fmt.Sprint(s.id)
}

// Bytes returns a byte slice of the snowflake ID
func (s *id) Bytes() []byte {
	return []byte(fmt.Sprintf("%d", s.id))
}
