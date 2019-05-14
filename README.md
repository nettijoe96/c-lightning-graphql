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

All graphql mutations calls are protected under Json Web Token authentification. [jwt-factory](https://github.com/nettijoe96/jwt-factory) is needed. You need to use the same cert and keyfile that you created here for jwt-factory (because the cmdline options have the same name)

Supported c-lightning commands: the argument structure follows exactly as c-lightning/lightning-cli. 

Protected by token authentification: 
the auth for a command requires a token with "graphql-<command-name>" or "graphql-admin"
1. close
2. connect
3. delinvoice
4. disconnect
5. fundchannel
6. invoice
7. pay
8. sendpay

No authentification required:
1. decodepay
2. feerates
3. getinfo
4. getroute
5. listchannels
6. listforwards
7. listfunds
8. listinvoices
9. listnodes
10. listpayments
11. listpeers
12. waitanyinvoice
13. waitinvoice
