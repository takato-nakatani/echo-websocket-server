package chat

import (
	"github.com/labstack/echo"
	"github.com/gorilla/websocket"
	"github.com/takato-nakatani/ChatEcho/backend/models"
	"fmt"
)

var (
	upgrader = websocket.Upgrader{}
)

func UserChat(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		models.SaveChat(string(msg))

		fmt.Printf("%s\n", msg)
	}
}
