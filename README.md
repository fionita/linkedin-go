# linkedin-go

## Installation

```bash
go get github.com/fionita/linkedin-go
```

## Usage

```go

package main

import (
	"fmt"

	linkedin "github.com/fionita/linkedin-go"
)

func main() {
	client, err := linkedin.Init(
		&linkedin.Config{
			AccessToken: "<ACCESS_TOKEN>",
		},
	)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// parameters id, fields
	resp, err := client.PeopleProfile("<ID>", []string{"id", "firstName", "lastName"})

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%v", resp)
}

```

### Share on linkedin

```go
...
content := map[string]interface{}{
  "comment": "test go!",
  "content": map[string]string{
    "title":         "LinkedIn Developers Resources",
    "description":   "Leverage LinkedIn's APIs to maximize engagement",
    "submitted-url": "https://developer.linkedin.com",
  },
  "visibility": map[string]string{
    "code": "anyone",
  },
}
resp, err := client.PeopleShare(content)
...

```

### Manage Company Pages
Required permission: rw_company_admin

#### Company Profile

```go
...
// parametes id, fields
// required id
resp, err := client.CompanyProfile("2414183", []string{"id", "name", "ticker", "description"})
...
```

### Get a company's updates

```go
...
params := map[string]string{
	"event-type": "status-update",
	"count":      "10",
	"start":      "0",
}
resp, err := client.CompanyUpdates("2414183", params)
...
```

### Get a specific company update

```go
...
resp, err := client.CompanyUpdate("<ID>", "<UPDATE-KEY>")
...
```
