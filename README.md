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

#### Company Profile

```go
...
resp, err := client.CompanyProfile("2414183", []string{"id", "name", "ticker", "description"})
...
```
