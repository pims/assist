## Assist: Dribbble client written in Go.

To get started:

```go
import "github.com/pims/assist"
```

then 

```go
client := assist.NewClient("API-KEY")
user, err := client.Users.Get("simplebits") // handle err
fmt.Println(user.Name) // Dan Cederholm
```


or 

```go
config := assist.NewConfig(os.Getenv("DRIBBBLE_TOKEN"), assist.DefaultApiEndpoint)
client = assist.NewClient(config) 
user, err = client.Users.Get("simplebits") // handle err
fmt.Println(user.Name) // Dan Cederholm
```
