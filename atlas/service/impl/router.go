package impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var host = os.Getenv("ROUTER_HOST")

type RouterController struct {
}

func InitRouterController() *RouterController {
	return &RouterController{}
}

type evalResponse struct {
	Priority []string
}

func (r *RouterController) GetPreference(paymentMode string, metadata map[string]interface{}) (string, error) {
	url := host + "/eval"
	innerMap := map[string]string{
		"bank_name": "Yes Bank",
	}
	data := map[string]interface{}{
		"payment_mode": "netbanking",
		"eval_params":  innerMap,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		return "", err
	}
	req, _ := http.NewRequest("POST", url, &buf)

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(res)
	fmt.Println(string(body))

	evalResp := evalResponse{}
	if err := json.Unmarshal(body, &evalResp); err != nil {
		fmt.Println("error in unmarshalling megamind /eval --> ", err, string(body))
		return "", err
	}
	return evalResp.Priority[0], nil
}
