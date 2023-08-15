package razorpay

import (
	"atlas/models"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"

	"golang.org/x/net/html"
)

type Razorpay struct {
	AppId  string
	Secret string
}

var baseUrl = "https://api.razorpay.com/v1"

func (c *Razorpay) GetName() string {
	return "razorpay"
}

func (c *Razorpay) CreatePayment(ctx context.Context, pi *models.PaymentIntent) (string, error) {
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
	fmt.Println("rzp order created --> ", order, string(body))

	url = baseUrl + "/payments/create/checkout"
	data := c.nbPayload(order.ID, pi)
	r, err = http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	c.addFormEncodedHeaders(r)

	//debug
	fmt.Println("Request Method:", r.Method)
	fmt.Println("Request URL:", r.URL.String())
	fmt.Println("Request Headers:")
	for key, value := range r.Header {
		fmt.Println(key, ":", value)
	}
	fmt.Println("Request Body:", data.Encode())
	//end debug

	res, err = http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("error in /orders/pay --> ", err)
		return "", err
	}
	defer res.Body.Close()

	// body, err = ioutil.ReadAll(res.Body)
	// fmt.Println("body ---> ", string(body))
	doc, err := html.Parse(res.Body)
	if err != nil {
		fmt.Println("error in parsing html from rzp", err)
		log.Fatal(err)
	}
	targetID := "form1"
	formNode := findFormWithID(doc, targetID)
	htmlString := getHTML(targetID, nodeToString(formNode))
	fmt.Println("Found here --> ", htmlString)
	return htmlString, nil
}

func (c *Razorpay) addHeaders(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	auth := c.AppId + ":" + c.Secret
	base64Auth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", "Basic "+base64Auth)
}

func (c *Razorpay) addFormEncodedHeaders(req *http.Request) {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Referer", "https://api.razorpay.com/v1/checkout/embedded")
}

func reqToPayload(pi *models.PaymentIntent) *bytes.Buffer {
	m := make(map[string]interface{})
	m["amount"] = pi.Amount * 1000
	m["currency"] = pi.Currency

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)
	return b
}

func (c *Razorpay) nbPayload(orderId string, pi *models.PaymentIntent) url.Values {
	data := url.Values{}
	data.Add("method", "netbanking")
	data.Add("bank", "YESB")
	data.Add("email", "gaurav.kumar@example.com")
	data.Add("contact", "9123456780")
	data.Add("order_id", orderId)
	data.Add("amount", fmt.Sprintf("%.0f", 1000*pi.Amount))
	data.Add("currency", pi.Currency)
	data.Add("key_id", c.AppId)

	return data
}

func findFormWithID(node *html.Node, targetID string) *html.Node {
	if node.Type == html.ElementNode && node.Data == "form" {
		for _, attr := range node.Attr {
			if attr.Key == "id" && attr.Val == targetID {
				return node
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if form := findFormWithID(c, targetID); form != nil {
			return form
		}
	}
	return nil
}

func nodeToString(node *html.Node) string {
	var buf bytes.Buffer
	html.Render(&buf, node)
	return buf.String()
}

func getHTML(formId string, form string) string {
	text := `
	<body>
	  <script>
		document.addEventListener("DOMContentLoaded", function() {
			setTimeout(function() {
			document.getElementById("{{.FormId}}").submit();
			}, 5); // Delay in milliseconds (5 seconds)
		});
	</script>
	  </script>
	</head>
	<body>
	  {{.Form}}
	</body>
`
	type d struct {
		Form   string
		FormId string
	}
	data := &d{FormId: formId, Form: form}
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
