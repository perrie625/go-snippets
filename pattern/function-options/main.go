package main

import "fmt"

var defaultStuffClientOptions = StuffClientOptions{
	Retries: 3,
	Timeout: 2,
}
type StuffClientOption func(*StuffClientOptions)
type StuffClientOptions struct {
	Retries int //number of times to retry the request before giving up
	Timeout int //connection timeout in seconds
}
func WithRetries(r int) StuffClientOption {
	return func(o *StuffClientOptions) {
		o.Retries = r
	}
}
func WithTimeout(t int) StuffClientOption {
	return func(o *StuffClientOptions) {
		o.Timeout = t
	}
}
type StuffClient interface {
	Describe()
}
type stuffClient struct {
	conn    Connection
	timeout int
	retries int
}
type Connection struct {}
func NewStuffClient(conn Connection, opts ...StuffClientOption) StuffClient {
	options := defaultStuffClientOptions
	for _, o := range opts {
		o(&options)
	}
	return &stuffClient{
		conn:    conn,
		timeout: options.Timeout,
		retries: options.Retries,
	}
}
func (c stuffClient) Describe() {
	fmt.Printf("timeout: %d, retries: %d.\n", c.timeout, c.retries)
}

func main() {
	defaultClient := NewStuffClient(Connection{})
	defaultClient.Describe()

	client1 := NewStuffClient(Connection{}, WithRetries(5))
	client1.Describe()

	client2 := NewStuffClient(Connection{}, WithTimeout(10))
	client2.Describe()

	client3 := NewStuffClient(Connection{}, WithTimeout(8), WithRetries(10))
	client3.Describe()
}