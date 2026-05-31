package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health godoc
//
//	@Summary		Health Check
//	@Description	Health Check
//	@Tags			system
//	@Accept			x-www-form-urlencoded
//	@Produce		plain
//	@Success		200	"OK"
//	@Router			/api/health [get]
func (s *Service) Health(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
