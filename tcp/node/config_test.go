package node_test

import (
	"os"
	"testing"

	"github.com/abinashphulkonwar/dist-cache/tcp/node"
)

func TestWriteConfig(t *testing.T) {
	os.Setenv("mode", "test")
	node.WriteConnectionConfig()
}
