package authclient

import "github.com/Floor-Gang/authserver/pkg"

func (ac AuthClient) Verify(memberID string) (bool, error) {
	response, err := ac.Auth(memberID)

	if err != nil {
		return false, err
	} else {
		return response.IsAdmin, nil
	}
}

func (ac AuthClient) Auth(memberID string) (pkg.AuthResponse, error) {
	response := pkg.AuthResponse{}
	args := pkg.AuthArgs{MemberID: memberID}

	err := ac.client.Call("AuthServer.Auth", args, &response)

	if err != nil {
		return response, err
	} else {
		return response, nil
	}
}
