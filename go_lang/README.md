## Setup and Run test
```
#Try1

cd unit-test--setup-in-all-languages/go_lang
go test -v ./test/

## OR try

## Check go path
echo $GOPATH
## Default would be ~/go
## Copy folder to go path module
cp -R ../go_lang $GOPATH/src/github.com/xdeepakv

cd $GOPATH/src/github.com/xdeepakv/go_lang

## run test case
go test -v ./test/
```