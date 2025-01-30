# alma-go

CLI and Go client for the Ex Libris Alma api

## Configuration

Configuration is through environment variables:

```sh
export ALMA_URL="https://api-eu.hosted.exlibrisgroup.com/almaws/v1"
export ALMA_API_KEY="****"
```

## Using the CLI

```sh
alma users get 1234 > /tmp/user-record.json
cat /tmp/user-record.json | alma user update 1234

alma requested-resources get --circ-desk DEFAULT_CIRC_DESK --library MAIN

alma printouts get
```

## Using the Go library

```go
import 	"github.com/ugentlib/alma-go"

almaClient := alma.New(alma.Config{
	URL:    "",
	ApiKey: "",
})

user, err := almaClient.GetUser(ctx, "1234")
```

## Running via Docker

```sh
docker pull ugentlib/alma-client:latest
docker run ugentlib/alma-client users get 1234
```

## Building from source

```sh
git clone git@github.com:ugent-library/alma-go.git
cd alma-go
go build -o ~/bin/alma cmd/alma/*
```
