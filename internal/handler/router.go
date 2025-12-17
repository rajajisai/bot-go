package handler

import (
	"net/http"
	"runtime/debug"

	"bot-go/internal/controller"
	"bot-go/pkg/mcp"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRouter(repoController *controller.RepoController, mcpServer *mcp.CodeGraphServer, codeAPIController *controller.CodeAPIController, logger *zap.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(CustomRecoveryMiddleware(logger))
	router.Use(LoggerMiddleware(logger))

	v1 := router.Group("/api/v1")
	{
		v1.POST("/buildIndex", repoController.BuildIndex)
		//v1.POST("/getFunctionsInFile", repoController.GetFunctionsInFile)
		//v1.POST("/getFunctionDetails", repoController.GetFunctionDetails)
		v1.POST("/functionDependencies", repoController.GetFunctionDependencies)
		v1.POST("/processDirectory", repoController.ProcessDirectory)
		v1.POST("/searchSimilarCode", repoController.SearchSimilarCode)

		// Index building endpoints
		v1.POST("/indexFile", repoController.IndexFile)

		// N-gram endpoints
		v1.POST("/processNGram", repoController.ProcessNGram)
		v1.POST("/getNGramStats", repoController.GetNGramStats)
		v1.POST("/getFileEntropy", repoController.GetFileEntropy)
		v1.POST("/analyzeCode", repoController.AnalyzeCode)
		v1.POST("/calculateZScore", repoController.CalculateZScore)

		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "healthy",
			})
		})
	}

	// CodeAPI routes
	if codeAPIController != nil {
		codeAPI := router.Group("/codeapi/v1")
		{
			// Reader endpoints
			codeAPI.GET("/repos", codeAPIController.ListRepos)
			codeAPI.POST("/files", codeAPIController.ListFiles)
			codeAPI.POST("/classes", codeAPIController.ListClasses)
			codeAPI.POST("/methods", codeAPIController.ListMethods)
			codeAPI.POST("/functions", codeAPIController.ListFunctions)
			codeAPI.POST("/classes/find", codeAPIController.FindClasses)
			codeAPI.POST("/methods/find", codeAPIController.FindMethods)
			codeAPI.POST("/class", codeAPIController.GetClass)
			codeAPI.POST("/method", codeAPIController.GetMethod)
			codeAPI.POST("/class/methods", codeAPIController.GetClassMethods)
			codeAPI.POST("/class/fields", codeAPIController.GetClassFields)

			// Analyzer endpoints
			codeAPI.POST("/callgraph", codeAPIController.GetCallGraph)
			codeAPI.POST("/callers", codeAPIController.GetCallers)
			codeAPI.POST("/callees", codeAPIController.GetCallees)
			codeAPI.POST("/data/dependents", codeAPIController.GetDataDependents)
			codeAPI.POST("/data/sources", codeAPIController.GetDataSources)
			codeAPI.POST("/impact", codeAPIController.GetImpact)
			codeAPI.POST("/inheritance", codeAPIController.GetInheritanceTree)
			codeAPI.POST("/field/accessors", codeAPIController.GetFieldAccessors)

			// Raw Cypher endpoints
			codeAPI.POST("/cypher", codeAPIController.ExecuteCypher)
			codeAPI.POST("/cypher/write", codeAPIController.ExecuteCypherWrite)

			// Health check
			codeAPI.GET("/health", func(c *gin.Context) {
				c.JSON(200, gin.H{"status": "healthy"})
			})
		}
	}

	// Setup MCP routes
	mcpServer.SetupHTTPRoutes(router)

	return router
}

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("HTTP Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("client_ip", c.ClientIP()),
		)
		c.Next()
	}
}

func CustomRecoveryMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Panic recovered",
					zap.Any("error", err),
					zap.String("stack", string(debug.Stack())),
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
				)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal server error",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
