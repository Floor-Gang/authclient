package authclient

type AuthArgs struct{ MemberID string }

type AuthResponse struct {
	Role    string
	IsAdmin bool
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
