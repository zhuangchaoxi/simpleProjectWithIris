package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
)

var hmacJwtSecret = []byte("MjAxOeW5tDEy5pyIIDLml6Ug5pi")

func Authentication(ctx iris.Context) {

	//tokenString := ctx.GetHeader("authorization")
	tokenString := ctx.GetCookie("authorization")

	if tokenString == "" {
		ctx.StatusCode(iris.StatusForbidden)
		logrus.Error("token is empty")
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacJwtSecret, nil
	})

	if err != nil {
		ctx.StatusCode(iris.StatusForbidden)
		logrus.Error(err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		(*ctx.Values()).Set("company_id", claims["company_id"])
	} else {
		ctx.StatusCode(iris.StatusForbidden)
		logrus.Error(err)
		return
	}

	ctx.Next()
}
