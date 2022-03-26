package client

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/k3s-io/kine/pkg/endpoint"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Value struct {
	Key      []byte
	Data     []byte
	Modified int64
}

var (
	ErrNotFound = errors.New("key not found")
)

type Client interface {
	List(ctx context.Context, key string, rev int) ([]Value, error)
	Get(ctx context.Context, key string) (Value, error)
	Put(ctx context.Context, key string, value []byte) error
	Create(ctx context.Context, key string, value []byte) error
	Update(ctx context.Context, key string, revision int64, value []byte) error
	Delete(ctx context.Context, key string, revision int64) error
	Version() (string, error)
	Close() error
}

type ClientK struct {
	c      *clientv3.Client
	config *endpoint.ETCDConfig
}

func New(config endpoint.ETCDConfig) (Client, error) {
	tlsConfig, err := config.TLSConfig.ClientConfig()
	if err != nil {
		return nil, err
	}

	c, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Endpoints,
		DialTimeout: 5 * time.Second,
		TLS:         tlsConfig,
	})
	if err != nil {
		return nil, err
	}

	return &ClientK{
		c:      c,
		config: &config,
	}, nil
}

func (c *ClientK) List(ctx context.Context, key string, rev int) ([]Value, error) {
	resp, err := c.c.Get(ctx, key, clientv3.WithPrefix(), clientv3.WithRev(int64(rev)))
	if err != nil {
		return nil, err
	}

	var vals []Value
	for _, kv := range resp.Kvs {
		vals = append(vals, Value{
			Key:      kv.Key,
			Data:     kv.Value,
			Modified: kv.ModRevision,
		})
	}

	return vals, nil
}

func (c *ClientK) Get(ctx context.Context, key string) (Value, error) {
	resp, err := c.c.Get(ctx, key)
	if err != nil {
		return Value{}, err
	}

	if len(resp.Kvs) == 1 {
		return Value{
			Key:      resp.Kvs[0].Key,
			Data:     resp.Kvs[0].Value,
			Modified: resp.Kvs[0].ModRevision,
		}, nil
	}

	return Value{}, ErrNotFound
}

func (c *ClientK) Put(ctx context.Context, key string, value []byte) error {
	val, err := c.Get(ctx, key)
	if err != nil {
		if err == rpctypes.ErrKeyNotFound {
			fmt.Println("---------")
		}
		return err
	}
	if val.Modified == 0 {
		return c.Create(ctx, key, value)
	}
	return c.Update(ctx, key, val.Modified, value)
}

func (c *ClientK) Create(ctx context.Context, key string, value []byte) error {
	resp, err := c.c.Txn(ctx).
		If(clientv3.Compare(clientv3.ModRevision(key), "=", 0)).
		Then(clientv3.OpPut(key, string(value))).
		Commit()
	if err != nil {
		return err
	}
	if !resp.Succeeded {
		return fmt.Errorf("key exists")
	}
	return nil
}

func (c *ClientK) Update(ctx context.Context, key string, revision int64, value []byte) error {
	resp, err := c.c.Txn(ctx).
		If(clientv3.Compare(clientv3.ModRevision(key), "=", revision)).
		Then(clientv3.OpPut(key, string(value))).
		Else(clientv3.OpGet(key)).
		Commit()
	if err != nil {
		return err
	}
	if !resp.Succeeded {
		return fmt.Errorf("revision %d doesnt match", revision)
	}
	return nil
}

func (c *ClientK) Delete(ctx context.Context, key string, revision int64) error {
	resp, err := c.c.Txn(ctx).
		If(clientv3.Compare(clientv3.ModRevision(key), "=", revision)).
		Then(clientv3.OpDelete(key)).
		Else(clientv3.OpGet(key)).
		Commit()
	if err != nil {
		return err
	}
	if !resp.Succeeded {
		return fmt.Errorf("revision %d doesnt match", revision)
	}
	return nil
}

func (c *ClientK) Close() error {
	return c.c.Close()
}

func (c *ClientK) Version() (string, error) {
	// pool := x509.NewCertPool()
	// caCrt, err := ioutil.ReadFile(c.config.TLSConfig.CAFile)
	// if err != nil {
	// 	return "", err
	// }
	// pool.AppendCertsFromPEM(caCrt)

	clientCrt, err := tls.LoadX509KeyPair(c.config.TLSConfig.CertFile, c.config.TLSConfig.KeyFile)
	if err != nil {
		return "", err
	}
	tr := &http.Transport{
		//
		TLSClientConfig: &tls.Config{
			//RootCAs:            pool,		// disable verify server cert
			InsecureSkipVerify: true,
			Certificates: []tls.Certificate{
				clientCrt,
			},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(getVersionUrl(c.config.Endpoints))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func getVersionUrl(endpoints []string) string {
	parts := strings.SplitN(endpoints[0], "://", 2)
	if len(parts) > 1 {
		return "https://" + parts[1] + "/version"
	}
	return "https://" + parts[0] + "/version"
}
