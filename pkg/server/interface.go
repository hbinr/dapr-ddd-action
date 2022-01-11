package server

import (
	"github.com/gofiber/fiber/v2"
)

// Server http or grpc server
type Server interface {
	// RegisterHTTPRouter register http router
	RegisterHTTPRouter(r *fiber.App)
}
