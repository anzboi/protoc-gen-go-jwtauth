
package pkg

import (
	"github.com/anzx/pkg/jwtauth"
)


func ValidateHelloScopes(claims jwtauth.Claims) bool {
	return claims.HasScopes(["fabric.read","fabric.write"])
}

func ValidateScopes(claims jwtauth.Claims, methodName string) bool {
	switch methodName {
	case "test.Hello/Hello":
		return ValidateHelloScopes(claims)
	default:
		return false
	}
}
