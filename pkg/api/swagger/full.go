package swagger

import (
	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/perchance-image-lab/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// fullSwagger godoc
//
//	@title						Perchance Image Lab API
//	@version					1.0
//	@description				API documentation for the Perchance Image Lab server.
//	@termsOfService				http://swagger.io/terms/
//	@BasePath					/
//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func fullSwagger() gin.HandlerFunc {
	return ginSwagger.WrapHandler(
		swaggerfiles.NewHandler(),
		ginSwagger.PersistAuthorization(true),
		ginSwagger.InstanceName(docs.SwaggerInfoFull.InstanceName()),
	)
}
