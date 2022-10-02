package producer

import (
	_ "encoding/json"
	_ "fmt"
	_ "log"

	_ "github.com/Shopify/sarama/mocks"
	"github.com/gofiber/fiber/v2"
)
type Commment struct {
	Text string `form : "text" json : "text" binding : "required"`
}

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")



	}
	app.Listen(":3000")
}
