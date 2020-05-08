package pkg

import (
	"github.com/anzx/pkg/jwtauth"
)

func ValidateHelloScopes(claims jwtauth.Claims) bool {
	if claims.HasAllScopes([]string{"your.scope.read", "your.scope.write"}) {
		return true
	}
	if claims.HasAllScopes([]string{"another.scope"}) {
		return true
	}
	return false
}

func ValidateScopes(claims jwtauth.Claims, methodName string) bool {
	switch methodName {
	case "test.Hello/Hello":
		return ValidateHelloScopes(claims)
	default:
		return false
	}
}
