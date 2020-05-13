package main

import (
	"fmt"

	"github.com/ceph/go-ceph/rados"
)

var (
	// keep reference for destroy later
	c *rados.Conn
)

type cephTransport struct {
	conn *rados.Conn
}

func newCephTransport(path string) cephTransport {
	conn, err := rados.NewConn()
	if err != nil {
		panic(err)
	}

	if path == "" {
		err = conn.ReadDefaultConfigFile()
	} else {
		err = conn.ReadConfigFile(path)
	}
	if err != nil {
		conn.Shutdown()
		panic(err)
	}

	if err = conn.Connect(); err != nil {
		conn.Shutdown()
		panic(err)
	}

	// keep reference for destroy later
	c = conn

	return cephTransport{
		conn: conn,
	}
}

func (c cephTransport) RoundTrip(oid, poolName string) error {
	ioctx, err := c.conn.OpenIOContext(poolName)
	if err != nil {
		return err
	}
	defer ioctx.Destroy()

	stat, err := ioctx.Stat(oid)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", stat)

	data := make([]byte, stat.Size)
	if _, err = ioctx.Read(oid, data, 0); err != nil {
		return err
	}

	return nil
}

func shutdownCeph() {
	if c != nil {
		c.Shutdown()
	}
}

func main() {
	ceph := newCephTransport("/etc/ceph/ceph.conf")
	if err := ceph.RoundTrip("origin/62086805733799507", "ads-2019-03"); err != nil {
		fmt.Println(err)
	}
}
