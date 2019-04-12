graphql api for c-lightning. 


to run (must be already running bitcoind and c-lightning for backend):

1. `go get github.com/nettijoe96/c-lightning-api`

2. set LightningDir in config.go to full path.

3. `go get github.com/graphql-go/graphql`

4. `go get github.com/niftynei/glightning`

5. `go build -o c-lightning-api *.go`

6. `./c-lightning-api`

