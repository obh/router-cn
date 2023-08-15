package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dop251/goja"
	"github.com/labstack/echo/v4"
)

type handler struct {
	nbEvalScript goja.Callable
}

type rulerequest struct {
	PaymentMode string `json:"payment_mode"`
	EvalScript  string `json:"eval_script"`
}

type ruleevalrequest struct {
	Providers  []string      `json:"providers"`
	EvalScript nbevalrequest `json:"eval_params"`
}

type nbevalrequest struct {
	BankName string `json:"bank_name"`
}

func InitRoutes(e echo.Echo, nbProgram goja.Callable) {
	h := &handler{nbEvalScript: nbProgram}
	e.POST("/rule", h.AddRule)
	e.POST("/eval", h.EvalRule)
}

func (h *handler) AddRule(c echo.Context) error {
	req := new(rulerequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	//eval script
	fmt.Println("rule --> ", req)

	prog, err := goja.Compile("", req.EvalScript, true)
	if err != nil {
		fmt.Printf("Error compiling the script %v ", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	vm := goja.New()
	_, err = vm.RunProgram(prog)

	err = vm.ExportTo(vm.Get("nbEvalScript"), &h.nbEvalScript)
	if err != nil {
		fmt.Printf("Error exporting the function %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	resp := make(map[string]interface{})
	resp["status"] = "Accepted"
	resp["request"] = req
	return c.JSON(http.StatusOK, resp)
}

func (h *handler) EvalRule(c echo.Context) error {
	req := new(ruleevalrequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	vm := goja.New()
	res, err := h.nbEvalScript(goja.Undefined(), vm.ToValue("message from go"))
	if err != nil {
		fmt.Printf("Error calling function %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	result := res.ToString()
	providerList := strings.Split(result.String(), ",")
	resp := map[string]interface{}{
		"priority": providerList,
	}
	return c.JSON(http.StatusOK, resp)
}
