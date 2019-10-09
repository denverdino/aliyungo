package cs

import "net/http"

//modify cluster,include DeletionProtection and so on
type ModifyClusterArgs struct {
	DeletionProtection bool `json:"deletion_protection"`
}

//modify cluster
func (client *Client) ModifyCluster(clusterId string, args *ModifyClusterArgs) error {
	return client.Invoke("", http.MethodPut, "/api/v2/clusters/"+clusterId, nil, args, nil)
}
