package cpool

import (
	"errors"
	"github.com/toolkits/consistent"
	"github.com/toolkits/file"
	"github.com/toolkits/logger"
	"github.com/toolkits/rpool/conn_pool"
	"strings"
	"sync"
)

type RingBackend struct {
	sync.RWMutex
	Addrs map[string][]string
	Ring  *consistent.Consistent
	Pools map[string]*conn_pool.ConnPool
}

func (t *RingBackend) LocateRing(pkey string) (string, error) {
	t.RLock()
	defer t.RUnlock()

	if t.Ring == nil {
		return "", errors.New("nil ring")
	}

	name, err := t.Ring.Get(pkey)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (t *RingBackend) GetConnPoolsByName(name string) ([]*conn_pool.ConnPool, error) {
	conns := []*conn_pool.ConnPool{}

	t.RLock()
	defer t.RUnlock()

	addr_list, ok := t.Addrs[name]
	if !ok {
		return conns, errors.New("no such name")
	}

	if len(addr_list) == 0 {
		return conns, errors.New("empty addrs")
	}

	for _, addr := range addr_list {
		c, ok := t.Pools[addr]
		if !ok {
			continue
		}
		conns = append(conns, c)
	}
	if len(conns) == 0 {
		return conns, errors.New("no conn pool")
	}

	return conns, nil
}

func (t *RingBackend) InitRing(replicas int) {
	var tmp_ring *consistent.Consistent = consistent.New()
	tmp_ring.NumberOfReplicas = replicas

	t.RLock()
	for name, _ := range t.Addrs {
		tmp_ring.Add(name)
	}
	t.RUnlock()

	t.Lock()
	defer t.Unlock()

	t.Ring = tmp_ring
}

func (t *RingBackend) LoadAddrs(f string) error {
	if !file.IsExist(f) {
		return errors.New("backends file is not exist")
	}

	file_content, err := file.ToString(f)
	if err != nil {
		return err
	}
	file_content = strings.Trim(file_content, " \n\t")

	lines := strings.Split(file_content, "\n")
	if len(lines) == 0 {
		return errors.New("empty backends")
	}

	tmp_addrs := make(map[string][]string)
	for _, line := range lines {
		fields := strings.Fields(line)
		size := len(fields)
		if size < 2 {
			logger.Warn("invalid backend %s", line)
			continue
		}
		name := fields[0]
		addr := fields[1:size]
		tmp_addrs[name] = addr
	}

	t.Lock()
	defer t.Unlock()
	t.Addrs = tmp_addrs

	return nil
}

func (t *RingBackend) DestroyConnPools() {
	t.Lock()
	for _, pool := range t.Pools {
		pool.Destroy()
	}
	t.Unlock()
}
