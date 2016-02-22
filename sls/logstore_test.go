package sls

import (
	"fmt"
	"testing"
)

func TestShards(t *testing.T) {
	p := DefaultProject()
	logstore, err := p.Logstore("wp-accesslog")
	if err != nil {
		t.Fatalf("error find logstore %v", err)
	}

	shards, err := logstore.Shards()
	if err != nil {
		t.Fatalf("error find logstore %v", err)
	}

	fmt.Println(shards)
}

func TestLogstores(t *testing.T) {
	p := DefaultProject()
	list, err := p.Logstores()
	if err != nil {
		t.Fatalf("TestLogstores error: %v", err)
	}
	fmt.Println(list)
}

func TestCreateLogstore(t *testing.T) {
	p := DefaultProject()
	logstore := &Logstore{
		TTL:   2,
		Shard: 3,
		Name:  "test-hello",
	}
	if err := p.CreateLogstore(logstore); err != nil {
		t.Fatalf("error create logstore %v", err)
	}
	logstore, err := p.Logstore("test-hello")
	if err != nil {
		t.Fatalf("error find logstore %v", err)
	}
	if err := logstore.Delete(); err != nil {
		t.Fatalf("error delete logstore %v", err)
	}
}
