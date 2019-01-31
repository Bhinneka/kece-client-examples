package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"sync"
)

// Client SDK for Go
type Client interface {
	Set(key string, value interface{})
	Get(key string) []byte
	Del(key string) error
}

// Option for kece server
type Option struct {
	Network  string
	Host     string
	Port     int
	Username string
	Password string
}

type client struct {
	conn net.Conn
	lock sync.Mutex
}

// NewClient construct new client with option
func NewClient(op *Option) (Client, error) {
	if op.Network == "" {
		op.Network = "tcp"
	}
	conn, err := net.Dial(op.Network, fmt.Sprintf("%s:%d", op.Host, op.Port))
	if err != nil {
		return nil, err
	}
	return &client{conn: conn, lock: sync.Mutex{}}, nil
}

// Set data with key to kece server
func (c *client) Set(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	valByte, _ := json.Marshal(value)
	fmt.Fprintf(c.conn, "SET %s '%s'\n", key, string(valByte))
	bufio.NewReader(c.conn).ReadString('\n')
}

// Get data with key from kece server
func (c *client) Get(key string) []byte {
	c.lock.Lock()
	defer c.lock.Unlock()

	fmt.Fprintf(c.conn, "GET %s\n", key)
	message, _ := bufio.NewReader(c.conn).ReadString('\n')
	message = strings.TrimRight(message, "\n")
	return []byte(message)
}

// Del for delete data with key in kece server
func (c *client) Del(key string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	fmt.Fprintf(c.conn, "DEL %s\n", key)
	bufio.NewReader(c.conn).ReadString('\n')
	return nil
}
