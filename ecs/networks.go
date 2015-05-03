// API on Network

package ecs

import ()

// Allocate Public Ip Address
type AllocatePublicIpAddressArgs struct {
	InstanceId string
}

type AllocatePublicIpAddressResponse struct {
	CommonResponse

	IpAddress string
}

func (client *Client) AllocatePublicIpAddress(instanceId string) (ipAddress string, err error) {
	args := AllocatePublicIpAddressArgs{
		InstanceId: instanceId,
	}
	response := AllocatePublicIpAddressResponse{}
	err = client.Invoke("AllocatePublicIpAddress", &args, &response)
	if err != nil {
		return "", err
	}
	return response.IpAddress, nil
}
