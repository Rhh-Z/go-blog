package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(secret string, username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	fmt.Println(tokenString, "tokenString")
	if err != nil {
		panic(err)
	}
	return tokenString
}

func JWTAuth(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenString := tr.RequestHeader().Get("Authorization")
				// 分割token字符串
				auths := strings.SplitN(tokenString, " ", 2)
				fmt.Println(auths, "auths")
				if len(auths) != 2 || !strings.EqualFold(auths[0], "Token") {
					return nil, errors.New("jwt token missing")
				}
				// 校验token
				token, err := jwt.Parse(auths[1], func(token *jwt.Token) (any, error) {
					return []byte("secret"), nil
				}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

				if err != nil {
					return nil, err
				}

				if claims, ok := token.Claims.(jwt.MapClaims); ok {
					fmt.Println(claims["username"])
				} else {
					return nil, errors.New("Token Invalid")
				}
			}

			return handler(ctx, req)
		}
	}
}
