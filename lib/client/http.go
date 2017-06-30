package client

import (
	"net"
	"net/http"
	"time"
)

type ConfigOpt func(*Config)

type Config struct {
	dialer    *net.Dialer
	transport *http.Transport
	client    *http.Client
}

func NewHttpClient(configOpt ...ConfigOpt) *http.Client {
	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{IP: net.IPv4zero},
		KeepAlive: time.Second * 30,
		Timeout:   time.Second * 5,
	}
	transport := &http.Transport{
		Dial: dialer.Dial,
		ResponseHeaderTimeout: time.Second * 5,
	}
	client := &http.Client{
		Transport: transport,
	}
	config := &Config{
		dialer:    dialer,
		transport: transport,
		client:    client,
	}

	for _, opt := range configOpt {
		opt(config)
	}
	return config.client
}

func Timeout(duration time.Duration) ConfigOpt {
	return func(config *Config) {
		config.dialer.Timeout = duration
		config.transport.ResponseHeaderTimeout = duration
	}
}

func KeepAlive(duration time.Duration) ConfigOpt {
	return func(config *Config) {
		config.dialer.KeepAlive = duration
	}
}
