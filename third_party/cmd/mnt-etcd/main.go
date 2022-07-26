package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"k0s.io/pkg/tunnel/listener"
)

func main() {
	// port := ":8236"
	addr := "https://chassis.k0s.io/mnt/myetcd/"
	from := "etcd://user:pass@etcd.k0s.io:2379/"

	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	if len(os.Args) > 2 {
		from = os.Args[2]
	}

	log.Println("listening on", addr)
	http.Serve(listener.Listener(addr, "/"), handlers.LoggingHandler(os.Stderr, ldapMiddleware(etcdHandler(from))))
}

func parseEtcdUrl(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}

func NewClient(etcdurl string) (client *clientv3.Client, err error) {
	u := parseEtcdUrl(etcdurl)
	endpoint := u.Host               //"etcd.k0s.io:2379"
	username := u.User.Username()    //"user"
	password, _ := u.User.Password() //"pass"
	cfg := clientv3.Config{
		Endpoints:            []string{endpoint},
		DialTimeout:          10 * time.Second,
		DialKeepAliveTime:    10 * time.Second,
		DialKeepAliveTimeout: 10 * time.Second,
		Username:             username,
		Password:             password,
	}
	client, err = clientv3.New(cfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func etcdHandler(u string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// "etcd://user:pass@etcd.k0s.io:2379"
		switch r.Method {
		case http.MethodGet:
			handleGet(u, w, r)
		case http.MethodDelete:
			handleDelete(u, w, r)
		case http.MethodPut:
			handlePut(u, w, r)
		default:
		}
	})
}

func handlePut(etcdurl string, w http.ResponseWriter, r *http.Request) {
	client, err := NewClient(etcdurl)
	if err != nil {
		http.Error(w, "PUT: failed to initalize client", http.StatusInternalServerError)
		return
	}

	var key, val string = r.URL.Path, ""

	valBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("PUT: failed to parse request body"), http.StatusBadRequest)
		return
	}
	val = string(valBytes)

	resp, err := client.Put(context.Background(), key, val)
	if err != nil {
		http.Error(w, fmt.Sprintf("PUT: failed to put %s %s", key, val), http.StatusInternalServerError)
		return
	}
	_ = resp
}

func handleDelete(etcdurl string, w http.ResponseWriter, r *http.Request) {
	client, err := NewClient(etcdurl)
	if err != nil {
		http.Error(w, "DEL: failed to initalize client", http.StatusInternalServerError)
		return
	}

	var key string = r.URL.Path

	resp, err := client.Delete(context.Background(), key)
	if err != nil {
		http.Error(w, fmt.Sprintf("DEL: failed to del %s", key), http.StatusInternalServerError)
		return
	}
	_ = resp
}

func handleGet(etcdurl string, w http.ResponseWriter, r *http.Request) {
	client, err := NewClient(etcdurl)
	if err != nil {
		http.Error(w, "GET: failed to initalize client", http.StatusInternalServerError)
		return
	}

	var prefix string = r.URL.Path

	resp, err := client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		http.Error(w, fmt.Sprintf("GET: failed to get prefix %s", prefix), http.StatusInternalServerError)
		return
	}
	for _, kv := range resp.Kvs {
		fmt.Fprintln(w, string(kv.Key))
		fmt.Fprintln(w, string(kv.Value))
	}
}
