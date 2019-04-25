package auth

import (
	"context"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"fmt"
	"net/http"
)


type ResolverFunc func(p graphql.ResolveParams) (interface{}, error)

type AuthLevel int

const (
	NoAuth AuthLevel = iota
	FundsAuth
)

func AuthWrapper(resolver ResolverFunc, authLevel AuthLevel, p graphql.ResolveParams) (interface{}, error) {
        //var auth jwt = p.Args["auth"]  //this is where we get jwt TODO
	var isAuth bool
	var err error
	fmt.Println(p.Context.Value("token").(string))
	switch authLevel {
	//this is where we verify jwt is valid and has priviledges TODO
        case NoAuth:
		isAuth = true
	case FundsAuth:
		isAuth = checkAuth(FundsAuth, p.Context.Value("token").(string))
        }

	if isAuth {
		return resolver(p)
	}else{
		return "", err
	}
}


func checkAuth(authLevel AuthLevel, rawToken string) bool {
	_, err := CheckTokenSignature(rawToken)
	return err == nil
}


func GetAuthHandler(graphqlHandler http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//authorizationHeader := r.Header.Get("token")
		//add in token sig auth here! TODO
		var rawToken string = r.Header.Get("token")
		newCtx := context.WithValue(context.Background(), "token", rawToken)
		newR := r.WithContext(newCtx)
	        graphqlHandler.ServeHTTP(w, newR)
	})

}

var jwtSecret []byte = []byte("thepolyglotdeveloper")

func CheckTokenSignature(t string) (*jwt.Token, error) {
	if t == "" {
		return nil, errors.New("Token is not present")
	}
        token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		        return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	        }

	        // jwtSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtSecret, nil
		})
	return token, err
}

