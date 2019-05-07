graphql api for c-lightning. 


to run (must be already running bitcoind for backend):

1. Must have Go version 12

2. `git clone https://github.com/nettijoe96/c-lightning-graphql.git` (make sure it is outside gopath)

3. `go build -o c-lightning-graphql`

4. `ln -s c-lightning-graphql <path to c-lightning source>/plugins/c-lightning-graphql`

5. Create openssl rsa key and self signed cert for server

    `openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyin key.pem -out cert.pem`

6.  `./lightningd --keyfile=/path/to/server/keyfile --certfile=/path/to/server/certfile --graphql-port=<port> 

7. `./lightning-cli graphql`

Some api calls are protected under Json Web Token authentification. [jwt-factory](https://github.com/nettijoe96/jwt-factory) is needed. You need to use the same cert and keyfile that you created here for jwt-factory (because the cmdline options have the same name)


Supported c-lightning commands:
1. connect
2. delinvoice
3. feerates
4. getinfo
5. getroute
6. invoice
7. listchannels
8. listnodes
9. listpeers
10. pay
11. waitanyinvoice
