Three modules which work together
- atlas
- ui
- megamind

## Requirements
 -  MySQL: To connect and create tables see schema.sql

## Setup of service
### atlas
`go run server.go`

### ui
`npm run build` to  build nextjs project
`npm run start` to start the server

### megamind
`go run main.go` to start the go server

## Curl requests

### Create Payment Intent
```
curl --request POST \
  --url http://localhost:9811/process_payment \
  --header 'Content-Type: application/json' \
  --cookie PHPSESSID=0aqngbliomb1ra5f5k292qkkb7 \
  --data '{
	"customer_email": "rohit@cashfree.com",
	"customer_phone": "9908734801",
	"amount": 100,
	"currency": "INR"
}'
```

### Attempt payment
```
curl --request POST \
  --url http://localhost:9811/attempt_payment \
  --header 'Content-Type: application/json' \
  --cookie PHPSESSID=0aqngbliomb1ra5f5k292qkkb7 \
  --data '{
	"payment_intent_id": 1,
	"payment_method": "netbanking",
	"payment_method_details": {
		"bank": "Yes Bank"
	}
}'
```

### Evaluate rule for megamind
```
curl --request POST \
  --url http://localhost:9812/eval \
  --header 'Content-Type: application/json' \
  --cookie PHPSESSID=0aqngbliomb1ra5f5k292qkkb7 \
  --data '{
	"payment_mode": "netbanking",
	"eval_params": {
		"bank_name": "Yes Bank"
	}
}'
```

### Update rule for megamind
```
curl --request POST \
  --url http://localhost:9812/rule \
  --header 'Content-Type: application/json' \
  --cookie PHPSESSID=0aqngbliomb1ra5f5k292qkkb7 \
  --data '{
	"payment_mode": "netbanking",
	"eval_script": "function nbEvalScript(param) {    return \"cashfree\" }"
}'
```


