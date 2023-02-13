package hashing

import (
	"fmt"

	"github.com/buraksezer/consistent"
	"github.com/cespare/xxhash"
)

// In your code, you probably have a custom data type
// for your cluster members. Just add a String function to implement
// consistent.Member interface.
type myMember string

func (m myMember) String() string {
	return string(m)
}

// consistent package doesn't provide a default hashing function.
// You should provide a proper one to distribute keys/members uniformly.
type hasher struct {
	c *consistent.Consistent
}

func (h hasher) Sum64(data []byte) uint64 {
	// you should use a proper hash function for uniformity.
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

func (c *hasher) GetNode(key string) {
	keyByts := []byte(key)
	owner := c.c.LocateKey(keyByts)
	fmt.Println(owner.String(), key)
}
