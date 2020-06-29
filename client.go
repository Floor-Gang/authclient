package authclient

import (
	"net/rpc"
)

type AuthArgs struct{ MemberID string }

type AuthResponse struct {
	Role    string
	IsAdmin bool
}

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

func (ac AuthClient) Auth(memberID string) (AuthResponse, error) {
	response := AuthResponse{}
	args := AuthArgs{MemberID: memberID}

	err := ac.client.Call("AuthServer.Auth", args, &response)

	if err != nil {
		return response, err
	} else {
		return response, nil
	}
}
