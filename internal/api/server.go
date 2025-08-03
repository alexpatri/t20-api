package api

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"dnd-api/internal/utils/config"
)

// Server encapsula o servidor HTTP Fiber
type Server struct {
	app    *fiber.App
	config *config.Config
}

// Start inicia o servidor HTTP
func (s *Server) Start() error {
	return s.app.Listen(fmt.Sprintf(":%s", s.config.Server.Port))
}

// GetApp retorna a instância do Fiber para testes
func (s *Server) GetApp() *fiber.App {
	return s.app
}

// NewServer cria uma nova instância do servidor
func NewServer(config *config.Config) *Server {
	app := fiber.New(fiber.Config{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	return &Server{
		app:    app,
		config: config,
	}
}
