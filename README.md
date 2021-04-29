# dp-zebedee-sdk-go
Golang SDK for Zebedee CMS.

### Current functionality
- Open session / login
- Create user, Delete user, Get user(s)
- Set user password
- Set/Get user permissions



### Getting started
Get the library:
```
go get github.com/ONSdigital/dp-zebedee-sdk-go/zebedee@latest
```
Opening a session:
```go
import (
    "github.com/ONSdigital/dp-zebedee-sdk-go/zebedee"
)

...

var host = "http://localhost:8082"

...

httpCli := zebedee.NewHttpClient(time.Second * 5)
zebCli := zebedee.NewClient(host, httpCli)

c := zebedee.Credentials{
    Email:    "test.email@ons.gov.uk",
    Password: "this is my password",
}

sess, err := zebCli.OpenSession(c)
if err != nil {
    return err
}
```
