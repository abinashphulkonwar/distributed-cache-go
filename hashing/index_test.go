package hashing_test

import (
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/abinashphulkonwar/dist-cache/hashing"
)

var node = [4]string{"node1", "node2", "node3", "node4"}

func TestAdd(t *testing.T) {
	c := hashing.Init()
	c.AddNode(node[0])
}

func UUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return "", err
	}
	id := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return id, nil

}

func TestGet(t *testing.T) {
	c := hashing.Init()

	for i := range node {
		uuid, err := UUID()
		if err != nil {
			t.Error(err)
		}
		val := node[i]
		c.AddNode(val)
		c.GetNode(uuid)
	}

}
