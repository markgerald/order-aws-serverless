# GO / Gin framework Example of best pratices
## Description
- 2 Golang Aplications / DynamoDB Table / S3 Bucket / API Gateway / Lambda Functions / CloudWatch Logs / CloudWatch Events / IAM Roles / IAM Policies

## Infrastructure and Software Deployment
- Serverless Framework(to local deploy, from your machine)
- Setting ACCESS_KEY_ID and AWS_SECRET_ID in Github Actions Secrets

### Run local
serverless command:
```
serverless deploy --stage dev
```

## Testing api
To test this api, send a post request to the api endpoint, with the following model body:
```json
{
  "UserId": 123,
  "IsPayed": false,
  "Items": [
    {
      "Name": "Producy X",
      "Price": 25.30,
      "Amount": 2
    }
  ]
}

```
### Other http methods
From same endpoint, you can use 2 get methods:
- /{id} - Get a specific order
- / - Get all orders
