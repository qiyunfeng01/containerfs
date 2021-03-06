package cfs

import (
	"fmt"
	"github.com/ipdcode/containerfs/logger"
	"google.golang.org/grpc"
	"time"
)

// DialMeta
func DialMeta() (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error

	conn, err = grpc.Dial(MetaNodeAddr, grpc.WithInsecure())

	if err != nil {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("DialMeta 1: addr:%v\n", MetaNodeAddr)
		logger.Error("DialMeta 1: addr:%v", MetaNodeAddr)

		conn, err = grpc.Dial(MetaNodeAddr, grpc.WithInsecure())
		if err != nil {
			time.Sleep(300 * time.Millisecond)
			fmt.Printf("DialMeta 2: addr:%v\n", MetaNodeAddr)
			logger.Error("DialMeta 3: addr:%v", MetaNodeAddr)
			conn, err = grpc.Dial(MetaNodeAddr, grpc.WithInsecure())
		}
	}
	return conn, err
}

// DialData
func DialData(host string) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error
	conn, err = grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		time.Sleep(1000 * time.Millisecond)
		conn, err = grpc.Dial(host, grpc.WithInsecure())
		if err != nil {
			time.Sleep(1000 * time.Millisecond)
			conn, err = grpc.Dial(host, grpc.WithInsecure())
		}
	}
	return conn, err
}

// DialVolmgr
func DialVolmgr(host string) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error
	conn, err = grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		time.Sleep(300 * time.Millisecond)
		conn, err = grpc.Dial(host, grpc.WithInsecure())
		if err != nil {
			time.Sleep(300 * time.Millisecond)
			conn, err = grpc.Dial(host, grpc.WithInsecure())
		}
	}
	return conn, err
}
