## ## Setup and Run test
# Prerequisite:
* Nodejs
* npm/npx

```
$ mkdir src test
## Create util.js and add function in src diretcory

//util.js
exports.add = (a, b) => a+b


## Create test file in test folder[Name should contain test.js]
## Init npm package
$ npm init --y 
## install jest testing framework 
$ npm i --save-dev jest 
## Run test
$ npx jest
```