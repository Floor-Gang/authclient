package authclient

import (
	"net/rpc"
)

type AuthClient struct {
	client *rpc.Client
}

func GetClient(address string) (AuthClient, error) {
	client, err := rpc.DialHTTP("tcp", address)

	if err != nil {
		return AuthClient{}, err
	}

	return AuthClient{
		client: client,
	}, nil
}
