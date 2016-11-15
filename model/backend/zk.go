package backend

import (
	"path"
	"strings"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var defaultACL = zk.WorldACL(zk.PermAll)

type Zk struct {
	conn     *zk.Conn
	basePath string
}

func NewZkBackend() *Zk {
	return &Zk{
		basePath: "/openvdc",
	}
}

func (z *Zk) Connect(servers []string) error {
	if z.conn != nil {
		return ErrConnectionExists
	}
	c, f, err := zk.Connect(servers, time.Second)
	if err != nil {
		return err
	}
	<-f
	z.conn = c
	return nil
}

func (z *Zk) Close() error {
	if z.conn == nil {
		return ErrConnectionNotReady
	}
	defer func() {
		z.conn = nil
	}()
	z.conn.Close()
	return nil
}

func (z *Zk) setupParents(key string) error {
	var parts []string
	for i, k := range strings.Split(key, "/") {
		parts = append(parts, k)
		if i < 1 {
			// first iteration is blank.
			continue
		}
		partial := strings.Join(parts, "/")
		exists, _, err := z.conn.Exists(partial)
		if err != nil {
			return err
		}
		if !exists {
			_, err = z.conn.Create(partial, []byte{}, int32(0), defaultACL)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (z *Zk) ifNotExist(key string, callee func() error) error {
	exists, _, err := z.conn.Exists(key)
	if err != nil {
		return err
	}
	if !exists {
		return callee()
	}
	return nil
}

func (z *Zk) Create(key string, value []byte) error {
	if z.conn == nil {
		return ErrConnectionNotReady
	}
	var err error
	absKey := path.Clean(path.Join(z.basePath, key))
	dir, _ := path.Split(absKey)
	dir = strings.TrimSuffix(dir, "/")

	err = z.ifNotExist(dir, func() error {
		return z.setupParents(dir)
	})
	if err != nil {
		return err
	}
	_, err = z.conn.Create(absKey, value, int32(0), defaultACL)
	return err
}

func (z *Zk) Update(key string, value []byte) error {
	if z.conn == nil {
		return ErrConnectionNotReady
	}
	var err error
	absKey := path.Join(z.basePath, path.Clean(key))
	dir, _ := path.Split(absKey)

	err = z.ifNotExist(dir, func() error {
		return z.setupParents(dir)
	})
	if err != nil {
		return err
	}
	_, err = z.conn.Set(absKey, value, 1)
	return err
}