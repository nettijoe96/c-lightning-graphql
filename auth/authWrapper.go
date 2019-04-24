package auth

import (
	"github.com/graphql-go/graphql"
)


type ResolverFunc func(p graphql.ResolveParams) (interface{}, error)


func AuthWrapper(resolver ResolverFunc, p graphql.ResolveParams) (interface{}, error) {
        //var auth jwt = p.Args["auth"]  //this is where we get jwt TODO
	var isAuth bool = true
	//this is where we verify jwt is valid and has priviledges TODO
	var err error

	if isAuth {
		return resolver(p)
	}else{
		return "", err
	}
}

