package authclient

type AuthArgs struct{ MemberID string }

type AuthResponse struct {
	Role    string
	IsAdmin bool
}

func (ac AuthClient) CheckMember(memberID string) (bool, error) {
	response, err := ac.Auth(memberID)

	if err != nil {
		return false, err
	} else {
		return response.IsAdmin, nil
	}
}

func (ac AuthClient) GetRole(memberID string) (string, error) {
	response, err := ac.Auth(memberID)

	if err != nil {
		return "", err
	} else {
		return response.Role, nil
	}
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
