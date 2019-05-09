package auth

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/graphql-go/graphql"
	"github.com/nettijoe96/c-lightning-graphql/crypto"
	"github.com/nettijoe96/c-lightning-graphql/global"
	"log"
	"net/http"
)


type ResolverFunc func(p graphql.ResolveParams) (interface{}, error)

type AuthLevel string

const (
	NoAuth AuthLevel = "noauth" //the noauth string is not expected to be in a token. The rest of the levels are expected to be in a token
	Close = "graphql-close"
	Connect = "graphql-connect"
	Delinvoice = "graphql-delinvoice"
	Fundchannel = "graphql-fundchannel"
	Invoice = "graphql-invoice"
	Pay = "graphql-pay"
	Admin = "graphql-admin" //access to all commands
)

/* if token has at least 1 of the authlevels, it is authentificated to execute the command */
func AuthWrapper(resolver ResolverFunc, authLevels []AuthLevel, p graphql.ResolveParams) (interface{}, error) {
	var isAuth, hasPrivilege, isNotExpired bool
	var rawToken, certfile string
	var token *jwt.Token
	var err, e error
	var pub *rsa.PublicKey
	var containsToken bool = p.Context.Value("token") != nil
        if len(authLevels) == 1 && authLevels[0] == NoAuth {
		isAuth = true
	}else if containsToken {
		plugin := global.GetGlobalPlugin()
		certfile = plugin.GetOptionValue("certfile")
	        pub, err = crypto.LoadPubRSA(certfile)
		if err != nil {
			isAuth = false
			log.Print("server cannot load keyfile")
		}else{
	                rawToken = p.Context.Value("token").(string)
	                token, err = checkTokenSignature(rawToken, pub)
		        isNotExpired = checkTimestamp(token)
		        if err != nil {
			        isAuth = false
			        err = errors.Wrap(err, "token signature failed")
		        }else if !isNotExpired {
			        isAuth = false
			        err = errors.New("token is expired")
		        }else{
		                for _, authLevel := range authLevels {
		                        hasPrivilege, e = checkPrivilege(authLevel, token)
				        if e != nil {
					        err = errors.Wrap(err, e.Error())
				        }
					if hasPrivilege {
						isAuth = true
						break
					}
			        }
			}
                }
	}else{
		isAuth = false
	}

	if isAuth {
		return resolver(p)
	}else{
		return nil, errors.Wrap(err, "Not Authentificated")
	}
}


func checkTimestamp(token *jwt.Token) bool {
	//TODO
	return true
}

func checkPrivilege(authLevel AuthLevel, token *jwt.Token) (bool, error) {
	var hasPrivilege bool
	var err error
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	        if privileges, ok := claims["privileges"]; ok {
			for _, p := range privileges.([]interface{}) {
				if string(authLevel) == p.(string) {
					hasPrivilege = true
					break
				}
			}
			hasPrivilege = false
		}else{
		        err = errors.New("token does not contain privilege claim")
	        }
	}else{
		err = errors.New("token is invalid")
	}

	return hasPrivilege, err
}


func GetAuthHandler(graphqlHandler http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rawToken string = r.Header.Get("token")
		newCtx := context.WithValue(context.Background(), "token", rawToken)
		newR := r.WithContext(newCtx)
	        graphqlHandler.ServeHTTP(w, newR)
	})

}

func checkTokenSignature(t string, priv *rsa.PublicKey) (*jwt.Token, error) {
	if t == "" {
		return nil, errors.New("Token is not present")
	}
        token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
                if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		        return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	        }

		return priv, nil
		})
	return token, err
}

