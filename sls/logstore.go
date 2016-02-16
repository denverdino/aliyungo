package sls
import "encoding/json"

type Logstore struct {
	client *Client
	TTL    int `json:"ttl,omitempty"`
	Shard  int `json:"shardCount,omitempty"`
	Name   string `json:"logstoreName,omitempty"`
}


func (proj *Project) CreateLogstore(logstore *Logstore) error {

	data, err := json.Marshal(logstore)
	if err != nil {
		return err
	}

	req := &request{
		method:METHOD_POST,
		path: "/logstores",
		contentType:"application/json",
		payload: data,
	}

	return proj.client.requestWithClose(req)
}

func (proj *Project) Logstore(name string) (*Logstore, error) {
	req := &request{
		method: METHOD_GET,
		path: "/logstores/" + name,
	}

	ls := &Logstore{}
	if err := proj.client.requestWithJsonResponse(req, ls); err != nil {
		return nil, err
	}
	ls.client = proj.client
	return ls, nil
}

func (ls *Logstore ) Delete() error {
	req := &request{
		method:METHOD_DELETE,
		path: "/logstores/" + ls.Name,
	}
	return ls.client.requestWithClose(req)
}

func (ls *Logstore) Update() error {

	data, err := json.Marshal(ls)
	if err != nil {
		return err
	}
	req := &request{
		method: METHOD_PUT,
		path :"/logstores/" + ls.Name,
		contentType:"application/json",
		payload: data,
	}

	return ls.client.requestWithClose(req)
}

type LogstoreList struct {
	count     int `json:"count,omitempty"`
	total     int `json:"total,omitempty"`
	logstores []string `json:"logstores,omitempty"`
}

func (proj *Project) Logstores() (*LogstoreList, error) {
	req := &request{
		method: METHOD_GET,
		path: "/logstores",
	}
	list := &LogstoreList{}
	if err := proj.client.requestWithJsonResponse(req, list); err != nil {
		return nil, err
	}
	return list, nil
}

type shard struct {
	Id int `json:"shardID,omitempty"`
}

func (ls *Logstore) Shards() ([]int, error) {
	req := &request{
		method: METHOD_GET,
		path: "/logstores/" + ls.Name + "/shards",
	}

	var resp []*shard
	if err := ls.client.requestWithJsonResponse(req, &resp); err != nil {
		return nil, err
	}

	var shards []int
	for _, shard := range resp {
		shards = append(shards, shard.Id)
	}
	return shards, nil
}

