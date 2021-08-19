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
- Update collection
- List collections
- Update collection content
- Delete collection content
- Complete collection content
- Review collection content
- Approve collection
- Unlock collection
- Publish collection

#### Teams
- Add team member
- Remove team member
- Create new team
- Delete a team
- List teams
- Get team (by name)

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

#### Editing collection content


```go
import (
    "encoding/json"
    "github.com/ONSdigital/dp-zebedee-sdk-go/zebedee"
)

...

// declare a new collection description type to hold collection data
collection := zebedee.NewCollection("test1")

// call the create collection endpoint on Zebedee
collection, err := zebCli.CreateCollection(sess, collection)
if err != nil {
    return err
}

// declare the page content to be added to the collection
jsonString := `{"description":{"summary":"","keywords":[],"metaDescription":"","title":"Aboutus"},"markdown":["We are"],"type":"static_page","uri":"/about/contactus","breadcrumb":[{"uri":"/"},{"uri":"/about"}],"fileName":"contactus"}`
var content interface{}
err = json.Unmarshal([]byte(jsonString), &content)
if err != nil {
    return err
}

// update content to save the collection content in the 'in progress' state
err = zebCli.UpdateCollectionContent(sess, collection.ID, "/test/data.json", content)
if err != nil {
    return err
}

// complete content to mark the content as complete and ready for review
err = zebCli.CompleteCollectionContent(sess, collection.ID, "/test/data.json")
if err != nil {
    return err
}

// review content as a second user 
// please note that this must be a different user, hence using another session 'sess2'
err = zebCli.ReviewCollectionContent(sess2, collection.ID, "/test/data.json")
if err != nil {
    return err
}

// approve the collection once all content has been reviewed
err = zebCli.ApproveCollection(sess, collection.ID)
if err != nil {
    return err
}

// publish if it's a manual collection
err = zebCli.PublishCollection(sess, collection.ID)
if err != nil {
    return err
}

```