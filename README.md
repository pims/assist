## Assist: Dribbble client written in Go.

Assist is a Dribbble client written in Go. It is an attempt at learning how to build http wrappers in Go.
It is not feature complete at the moment, see [TODO.md](TODO.md) for feature implementation status.


To get started:

```go
import "github.com/pims/assist"
```

then 

```go
client := service.NewClient("API-KEY")
user, err := client.Users.Get("simplebits") // handle err
fmt.Println(user.Name) // Dan Cederholm
```

or 

```go
config := service.NewConfig(os.Getenv("DRIBBBLE_TOKEN"), assist.DefaultApiEndpoint)
client = service.NewClient(config) 
user, err = client.Users.Get("simplebits") // handle err
fmt.Println(user.Name) // Dan Cederholm
```