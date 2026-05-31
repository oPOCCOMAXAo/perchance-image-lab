package middleware

import "github.com/gin-gonic/gin"

// Init adds to context required information for other services.
func (s *Service) Init(ctx *gin.Context) {
	rawCtx := ctx.Request.Context()

	rawCtx = s.tasks.ContextWithSource(rawCtx, ctx.Request.Method+" "+ctx.Request.URL.Path)

	ctx.Request = ctx.Request.WithContext(rawCtx)
}
