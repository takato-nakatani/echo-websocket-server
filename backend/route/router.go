package route

import (
	"github.com/labstack/echo"
	"log"
	"../authentication"
	"github.com/takato-nakatani/ChatEcho/backend/api/chat"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderAccessControlAllowCredentials, echo.HeaderAccessControlAllowOrigin, echo.HeaderOrigin,
		                       echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderAuthorization},
	}))

	ws := e.Group("/ws")
	{
		ws.GET("", chat.UserChat)
	}

	return e
}

//echo.Middleware型の関数を用意する(引数にはhandlerを与える。これは戻り値で使用)
func authMiddleware(handler echo.HandlerFunc) echo.MiddlewareFunc {
	//　次にecho.Middlewareのシグネチャである、
	//　func(next echo.HandlerFunc) echo.HandlerFunc
	//　を実装する
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		//　最後にecho.Handlerのシグネチャである、
		//　HandlerFunc func(Context) error
		//　を実装する。
		return	func(c echo.Context) error{
			token := authentication.Auth(c)
			log.Printf("Verifying ID token: %v\n", token)
			return handler(c)
		}
	}
}