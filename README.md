# dp-zebedee-sdk-go

Golang SDK for Zebedee CMS.

### Current functionality

#### Auth:
- Open session (sign in)
- Set Password
- Get user permissions
- Set user permissions

#### Users:
- Create
- Delete
- Get user(s)

#### Collections:
- Create collection
- Get collection (by ID)
- Delete collection
- List collections

#### Teams
- Add team member
- Remove team member
- Create new team
- Delete a team
- List teams

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
