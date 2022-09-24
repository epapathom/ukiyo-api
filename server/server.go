package server

import (
	"context"
	"os"
	"ukiyo/configuration"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

type Server struct{}

var ginLambda *ginadapter.GinLambda

func (s *Server) Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func (s *Server) Init() {
	router := GetRouter()

	if os.Getenv("APPLICATION_MODE") == "Lambda" {
		ginLambda = ginadapter.New(router)
		lambda.Start(s.Handler)
	} else {
		router.Run(configuration.ConfigurationSingleton.Server.Url)
	}
}
