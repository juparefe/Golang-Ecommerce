package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Context context.Context
var Config aws.Config
var err error

func StartAWS() {
	Context = context.TODO()
	Config, err = config.LoadDefaultConfig(Context, config.WithDefaultRegion("us-east-1"))

	if err != nil {
		panic("Error al cargar la configuracion .aws/config" + err.Error())
	}
}
