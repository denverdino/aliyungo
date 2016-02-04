package sls
import (
	"testing"
	"fmt"
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
		TTL: 3,
		Shard: 3,
		Name: "test-jjz",
	}
	if err := p.CreateLogstore(logstore); err != nil {
		t.Fatalf("error create logstore %v", err)
	}
	logstore, err := p.Logstore("test-jjz")
	if err != nil {
		t.Fatalf("error find logstore %v", err)
	}
	if err := logstore.Delete(); err != nil {
		t.Fatalf("error delete logstore %v", err)
	}
}

