# Get Started

This guide thoroughly describes the deployment steps to deploy Ukiyo API to a local environment and to AWS. It also provides a way to debug the API locally while utilizing AWS resources.

## Deployment

The sections below describe the deployment steps locally and on AWS.

### Local

To deploy the service locally, we have an option to run the *main.go* file and an option to build the project and execute the binary file, as suggested below.

1. To run the *main.go* file, run the following commands.

```
go run main.go
```

2. To build and execute the project, run the following commands.

```
go build
./ukiyo
```

The service will be accessible in the **localhost** of the device and to port **3000**. The default stage that will be used to connect to the AWS resources is the **dev** stage. Therefore, in order to connect the API to the AWS resources, we also need to deploy the **dev** stage on AWS.

### AWS

To deploy the service on AWS, first of all, we need build the service for the metal architecture utilized by AWS Lambda using the following command.

```
GOARCH=amd64 GOOS=linux go build
```

Then, we need to create an **env.yml** file and insert the required environment variables in the **env.dist.yml** file. Finally, we can deploy the service to a stage using the following commands.

```
sls deploy -s <STAGE> # e.g. sls deploy -s v1
```

If we omit the **-s** or **--stage** flag, the default stage that will be used by the serverless framework is the **dev** stage. The service will be accessible in the URL found in API Gateway > Stage.

## Debugging

To debug the service, first of all, we need to have deployed the service on AWS and the **dev** stage. Then, we need to have the following *launch.json* configurations in the *.vscode* directory in the root of the project.

```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "main.go"
        }
    ]
}
```

After setting the above configuration, we can set breakpoints in the code and click the F5 button to start debugging.
