## Setup and Run test
```
## Check go path
echo $GOPATH
## Default would be ~/go
## Copy folder to go path module
cp -R ../go_lang $GOPATH/src/github.com/xdeepakv

cd $GOPATH/src/github.com/xdeepakv/go_lang

## run test case
go test -v ./test/
```