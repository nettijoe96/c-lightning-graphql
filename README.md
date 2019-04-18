graphql api for c-lightning. 


to run (must be already running bitcoind and c-lightning for backend):

1. Must have Go version >= 11

2. `git clone https://github.com/nettijoe96/c-lightning-api.git` (make sure it is outside gopath)

3. set LightningDir in config.go to full path.

4. `go build -o c-lightning-api *.go`

5. `./c-lightning-api`

