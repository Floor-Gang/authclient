# Authentication Client Library
For Floor Gang bot. Every command should interface with this library to authorize each user that
uses an admin-related command. It does this by communicating to a process over RPC+HTTP. You should
use this library to verify each guild member trying to execute a command. This is easy to do.

Make sure an instance of [authserver](https://github.com/floor-gang/authserver) is running on
your local machine where the authclient can reach it.

## Example Usage
```go
package main

import auth "github.com/Floor-Gang/authclient"

func main() {
    client := auth.GetClient("http://localhost:6969")
    
    memberID := "1234"
    
    // result will be an AuthResponse which includes two properties:
    // IsAdmin (boolean): Whether or not they're allowed to us the command
    // Role (string): The ID of the role that's enabling them.
    result := client.Auth("member ID 1234")
   
    if result.IsAdmin {
    	// do stuff...
    } else {
        // tell them they're not allowed to use the command
    }
}
```

## Notes

### What if the auth server is down?
Don't allow execution of a given command, and certainly don't setup your own authentication
process unless it's specialized to particular roles rather than the shared ones stored in the
[authserver](https://github.com/floor-gang/authserver) 


### What if the auth server has an issue?
Report it here

## Getting
Since this library is under a private repository [knowledge yourself](https://medium.com/cloud-native-the-gathering/go-modules-with-private-git-repositories-dfe795068db4)
about getting packages in Go from private repositories.

Make sure your enviroment variable "GOPRIVATE" includes the following
```shell script
export GOPRIVATE='github.com/floor-gang'
```

Finally run this to get the library
```shell script
$ go get github.com/Floor-Gang/authclient@latest
```
