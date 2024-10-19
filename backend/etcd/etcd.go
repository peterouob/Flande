package etcd

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	eclient "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log"
	"time"
)

const (
	EtcdAddress = "http://localhost:2379"
	Prefix      = "service"
)

var client *eclient.Client

func init() {
	var err error
	client, err = eclient.New(eclient.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
}

func RegisterEtcd(ctx context.Context, serverName, serverAddr string) error {
	log.Println("Try to register etcd ...")
	lease := eclient.NewLease(client)
	cancelCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	leaseResp, err := lease.Grant(cancelCtx, 3)
	if err != nil {
		return err
	}
	leaseCancel, err := lease.KeepAlive(ctx, leaseResp.ID)
	if err != nil {
		return err
	}
	em, err := endpoints.NewManager(client, Prefix)
	if err != nil {
		return err
	}

	cancelCtx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := em.AddEndpoint(cancelCtx, fmt.Sprintf("%s/%s/%s", Prefix, serverName, uuid.New().String()),
		endpoints.Endpoint{Addr: serverAddr},
		eclient.WithLease(leaseResp.ID)); err != nil {
		return err
	}

	log.Println("Register etcd success")
	del := func() {
		log.Println("Register close")
		cancelCtx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		em.DeleteEndpoint(cancelCtx, serverName)
		lease.Close()
	}

	go func() {
		failedCount := 0
		for {
			select {
			case resp := <-leaseCancel:
				if resp != nil {
					//log.Println("keep alive success")
				} else {
					log.Println(resp)
					log.Println("keep alive failed")
					failedCount++
					for failedCount > 3 {
						del()
						if err := RegisterEtcd(ctx, serverName, serverAddr); err != nil {
							time.Sleep(time.Second)
							continue
						}
						return
					}
					continue
				}
			case <-ctx.Done():
				log.Println("context done, cleaning up...")
				del()
				client.Close()
				return
			}
		}
	}()
	return nil
}
