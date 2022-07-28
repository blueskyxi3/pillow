# pillow

pillow is a CouchDB client in [Go(Golang)](https://golang.org/).

## Resources

* [Installation](#installation)
* [Usage](#usage)
	* [Example](#example)

## Installation

Install pillow as you normally would for any Go package:

```bash
go get -u github.com/blueskyxi3/pillow/pkg/pillow
```

## Usage

Please consult the the [package documentation](https://godoc.org/github.com/blueskyxi3/pillow) for all available API methods, and a complete usage documentation.

### Example

For additional usage examples, [consult the wiki](https://github.com/blueskyxi3/pillow/wiki/Usage-Examples).

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/blueskyxi3/pillow/pkg/pillow"
)

func main() {
	client, err := pillow.New(dsn)
	if err != nil {
		panic(err)
	}

	db := client.Database(context.TODO(), "tenants")

	document := map[string]interface{}{
		"_id":            "tenants:john-doe",
		"first_name":     "John",
		"last_name":      "Doe",
	}

	_, err = db.CreateDocument(context.TODO(), document)
	if err != nil {
		panic(err)
	}

	fmt.Println("Tenant created")
}
```