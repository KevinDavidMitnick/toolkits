package httpclient

import (
	"testing"
	"time"
)

func TestHttpClientPool(t *testing.T) {
	p := NewHttpClientPool()
	// AddHttpClient && GetHttpClient && Size
	client1 := p.AddHttpClient("c1", time.Duration(1)*time.Second, time.Duration(15)*time.Second)
	client2 := p.AddHttpClient("c1", time.Duration(1)*time.Second, time.Duration(15)*time.Second)
	client3, found := p.GetHttpClient("c1")
	if !(found && client1 == client2 && client1 == client3 && p.Size() == 1) {
		t.Error("error, AddHttpClient & GetHttpClient & Size")
	}
	// RemoveHttpClient
	p.RemoveHttpClient("c1")
	if !(p.Size() == 0) {
		t.Error("error, RemoveHttpClient")
	}
	// RemoveAllHttpClients
	p.AddHttpClient("c1", time.Duration(1)*time.Second, time.Duration(15)*time.Second)
	p.AddHttpClient("c2", time.Duration(1)*time.Second, time.Duration(15)*time.Second)
	size := p.Size()
	p.RemoveAllHttpClients()
	if !(size == 2 && p.Size() == 0) {
		t.Error("error, RemoveAllHttpClients")
	}
}
