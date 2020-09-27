# Test Project
This app's API provides endpoint for concurrent processing of requested URLs using only standard Go packages.

## Required environment variables

| Variable              | Description                                | Default    | Optional |
|-----------------------|--------------------------------------------|:----------:|:--------:|
| `AFS_PORT_HTTP`       | TCP port on which to start HTTP server     |    8080    |   true   |

## How to build and run the app

Is required [Go](https://golang.org/doc/install). Use next command to lunch it:

```
$ go build 
# export env variables described above ...
./afs
```

