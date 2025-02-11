package secretmngr

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/juparefe/Golang-Ecommerce/awsgo"
	models "github.com/juparefe/Golang-Ecommerce/models"
)

// De esta forma se crea una lambda proxy por fuera de la VPC que accede a secrets manager
func GetSecretLambdaProxy(secretName string) (models.SecretRDSJson, error) {
	const LambdaProxyName = "secretsmanager-proxy"
	secretData := models.SecretRDSJson{}
	svc := lambda.NewFromConfig(awsgo.Config)
	// Crear la solicitud para la Lambda Proxy
	payload, err := json.Marshal(models.Request{SecretName: secretName})
	if err != nil {
		fmt.Println("Error getting secret value ", secretName, ": ", err.Error())
		return secretData, err
	}
	// Invocar la Lambda Proxy
	secretValue, err := svc.Invoke(awsgo.Context, &lambda.InvokeInput{
		FunctionName: aws.String(LambdaProxyName),
		Payload:      payload,
	})
	if err != nil {
		return secretData, fmt.Errorf("error invoking lambda proxy: %v", err)
	}
	// Procesar el secretValue y guardarlo en secretData
	err = json.Unmarshal(secretValue.Payload, &secretData)
	if err != nil {
		fmt.Println("Error unmarshalling secret: ", secretName, err.Error())
		return secretData, err
	}
	fmt.Println("Secret data OK: ")
	return secretData, nil
}

// De esta forma se accede a secrets manager dentro de la VPC a traves de un VPC enpoint y genera costos
func GetSecretVPCEndpoint(secretName string) (models.SecretRDSJson, error) {
	secretData := models.SecretRDSJson{}
	svc := secretsmanager.NewFromConfig(awsgo.Config)
	secretValue, err := svc.GetSecretValue(awsgo.Context, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println("Error getting secret value ", secretName, ": ", err.Error())
		return secretData, err
	}
	// Procesar el secretValue y guardarlo en secretData
	err = json.Unmarshal([]byte(*secretValue.SecretString), &secretData)
	if err != nil {
		fmt.Println("Error unmarshalling secret: ", secretName, err.Error())
		return secretData, err
	}
	fmt.Println("Secret data OK: ")
	return secretData, nil
}
