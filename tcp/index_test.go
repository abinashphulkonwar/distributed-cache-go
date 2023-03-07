package tcp_test

import (
	"net"
	"os"
	"testing"

	"github.com/abinashphulkonwar/dist-cache/storage"
	"github.com/abinashphulkonwar/dist-cache/tcp"
	"github.com/dgraph-io/badger/v3"
)

func Init() {
	connection, err := badger.Open(badger.DefaultOptions("").WithInMemory(true).WithLogger(nil))

	if err != nil {
		panic(err)
	}
	db := storage.NewBadgerStorage(connection)

	app := tcp.App(db)
	if app != nil {
		panic("Error")
	}
}

func Client() *net.TCPConn {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:3001")
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	return conn
}

func TestApp(t *testing.T) {
	Init()

}

func TestClient(t *testing.T) {
	conn := Client()
	println("Client connected", conn.RemoteAddr().String())
	conn.Write([]byte("data"))

	for {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			println("Read from server failed:", err.Error())
			os.Exit(1)
		}
		println(string(buf[0:n]))
	}
}
