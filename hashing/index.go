package hashing

import (
	"fmt"

	"github.com/buraksezer/consistent"
	"github.com/cespare/xxhash"
)

type myMember string

func (m myMember) String() string {
	return string(m)
}

type hasher struct {
	c *consistent.Consistent
}

func (h hasher) Sum64(data []byte) uint64 {

	return xxhash.Sum64(data)
}

func Init() *hasher {

	cfg := consistent.Config{
		PartitionCount:    7,
		ReplicationFactor: 20,
		Load:              1.25,
		Hasher:            hasher{},
	}

	c := consistent.New(nil, cfg)

	return &hasher{c: c}
}

func (c *hasher) AddNode(name string) {

	node := myMember(name)

	c.c.Add(node)

}

func (c *hasher) GetNode(key string) string {

	keyByts := []byte(key)
	owner := c.c.LocateKey(keyByts)

	fmt.Println(owner.String(), key)
	return owner.String()
}
