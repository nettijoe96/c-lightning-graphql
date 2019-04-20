graphql api for c-lightning. 


to run (must be already running bitcoind for backend):

1. Must have Go version >= 11

2. `git clone https://github.com/nettijoe96/c-lightning-api.git` (make sure it is outside gopath)

3. set LightningDir in config.go to full path.

4. `go build -o c-lightning-api`

5. `ln -s c-lightning-api <path to c-lightning source>/plugins/c-lightning-api`

6.  `./lightningd --graphql-port= <port> --graphql-page=<page>`

7. `./lightning-cli graphql`

