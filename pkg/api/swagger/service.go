package swagger

import (
	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/gopkg/pkg/utils/ginutils"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Register(router gin.IRouter) error {
	router.GET("/api/swagger/full/*any", fullSwagger())
	router.GET("/swagger", ginutils.StaticRedirect("/api/swagger/full/index.html"))
	router.GET("/api/swagger", ginutils.StaticRedirect("/api/swagger/full/index.html"))
	router.GET("/api/swagger/full", ginutils.StaticRedirect("/api/swagger/full/index.html"))

	return nil
}
