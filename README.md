```
$ go version

go version go1.18beta1 linux/amd64

$ go run cmd/test.go

<stack trace>
```

In `internal/foo/foo.go` you can try commenting out `UnmarshalBroken` vs `UnmarshalWorking` to observe expected vs actual behaviour.