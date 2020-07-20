package authclient

// Feature to register.
type Feature struct {
	Name          string       // ie. ".admin"
	Description   string       // "For managing administrator"
	Commands      []SubCommand // Optionally empty
	CommandPrefix string       // Optionally empty
}

// Commands that come with this feature.
type SubCommand struct {
	Name        string   // ie. "add"
	Description string   // ie. "This is for adding administrator roles"
	Example     []string // ie. [".admin", "add", "...role ID's"]
}

// Response from the server while registering.
type RegisterResponse struct {
	Token   string // A Discord bot access token
	Serving string // The ID of the Guild being served
}

func (ac AuthClient) Register(feature Feature) (RegisterResponse, error) {
	response := RegisterResponse{}

	err := ac.client.Call("AuthServer.Register", feature, &response)

	if err != nil {
		return response, err
	} else {
		return response, nil
	}
}
