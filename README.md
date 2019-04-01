# go-module-bug-example

A bug about go module vendor

## Project module dependency

```
$ cat go.mod 
module go.mod/example

require (
	go.mod/a v0.0.0
	go.mod/b v0.0.0

)

replace (
	go.mod/a => ../a
	go.mod/b => ../b
)
```

```
$ go mod graph
go.mod/example go.mod/a@v0.0.0
go.mod/example go.mod/b@v0.0.0
go.mod/b@v0.0.0 go.mod/a@v0.0.0
```

### Reproduce bug

<img width="664" alt="Screen Shot 2019-04-01 at 10 17 23 AM" src="https://user-images.githubusercontent.com/1459834/55299858-756db700-5467-11e9-9a24-0d5d6051d321.png">

### Test Versions:

Version | Bug Reprocuce
------------ | -------------
Go 1.11.1 ~ Go 1.11.6 | Yes
Go 1.12.0 ~ Go 1.12.1 | No

### Others
- [Bug issue](https://github.com/golang/go/issues/30857)



