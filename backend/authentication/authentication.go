package authentication

import (
	"firebase.google.com/go/auth"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"os"
	"firebase.google.com/go"
	"fmt"
	"strings"
	"github.com/labstack/echo"
	"log"
	"golang.org/x/net/context"
)

func Auth(c echo.Context) *auth.Token{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
	opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	authentication, err := app.Auth(context.Background())
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	//クライアント側から送られてきたJWTをheaderから取得
	authHeader := c.Request().Header.Get("Authorization")
	//文字列を変換
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)

	//JWTの認証
	token, err := authentication.VerifyIDToken(context.Background(), idToken)
	//JWTが無効ならHandlerに進まず別処理
	if err != nil {
		fmt.Printf("error verifying ID token: %v\n", err)
		//error：401
		os.Exit(1)
	}
	return token
}
