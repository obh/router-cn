package cashfree

import (
	"atlas/models"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type Cashfree struct {
	AppId  string
	Secret string
}

var baseUrl = "https://api.cashfree.com/pg"
var version = "2022-09-01"

func (c *Cashfree) GetName() string {
	return "cashfree"
}

func (c *Cashfree) CreatePayment(ctx context.Context, pi *models.PaymentIntent) (string, error) {
	url := baseUrl + "/orders"
	r, _ := http.NewRequest("POST", url, reqToPayload(pi))
	c.addHeaders(r)
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("error in /orders --> ", err)
		return "", err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	order := &Order{}
	if err := json.Unmarshal(body, &order); err != nil {
		fmt.Println("error in unmarshalling --> ", err, string(body))
		return "", err
	}
	fmt.Println("order created --> ", order, string(body))

	url = baseUrl + "/orders/sessions"
	r, err = http.NewRequest("POST", url, nbPayload(order.PaymentSessionID))
	c.addHeaders(r)
	res, err = http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("error in /orders/pay --> ", err)
		return "", err
	}
	defer res.Body.Close()

	payment := &Payment{}
	body, err = ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(body, &payment); err != nil {
		fmt.Println("error in unmarshalling /order/pay --> ", err, string(body))
		return "", err
	}
	fmt.Println("Found here --> ", payment, string(body))
	if payment.Action == "link" && payment.Channel == "link" {
		return getHTML(payment.Data.URL), nil
	}
	return "", errors.New("failed")
}

func (c *Cashfree) addHeaders(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-client-id", c.AppId)
	req.Header.Add("x-client-secret", c.Secret)
	req.Header.Add("x-api-version", version)
}

func reqToPayload(pi *models.PaymentIntent) *bytes.Buffer {
	m := make(map[string]interface{})
	m["order_id"] = strconv.Itoa(pi.Id) + "-" + strconv.FormatInt(time.Now().Unix(), 10)
	m["order_amount"] = pi.Amount
	m["order_currency"] = pi.Currency
	c := make(map[string]string)
	c["customer_id"] = "1234"
	c["customer_email"] = pi.CustomerEmail
	c["customer_phone"] = pi.CustomerPhone
	m["customer_details"] = c

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)
	return b
}

func nbPayload(sessionId string) *bytes.Buffer {
	m := make(map[string]interface{})
	n := make(map[string]interface{})
	n["channel"] = "link"
	n["netbanking_bank_code"] = 3058
	m["netbanking"] = n

	r := make(map[string]interface{})
	r["payment_session_id"] = sessionId
	r["payment_method"] = m

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)
	fmt.Println("/order/pay request --> ", r)
	return b
}

func getHTML(link string) string {
	text := `
<body>
  <script>
    setTimeout(function() {
      window.location.href = "{{.URL}}";
    }, 5000); // Delay in milliseconds (5 seconds)
  </script>
</head>
<body>
  <p>This page will automatically redirect in 5 seconds.</p>
  <a href="{{.URL}}">If not redirected, click here</a>
</body>
`
	type d struct {
		URL string
	}
	data := &d{URL: link}
	tmpl, err := template.New("test").Parse(text)
	if err != nil {
		panic(err)
	}
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
