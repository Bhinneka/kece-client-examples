# Simple Go Client for KECE

### Usage
- Run kece server
```shell
$ kece
```

- Build client binary from source and run it
```shell
$ go build && ./go
```

- Check data in kece server
```shell
$ nc localhost 9000
> GET test
{"a":"this is value of field 'A'","b":67}
```