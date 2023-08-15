package routes

import (
	"atlas/models"
	provider "atlas/providers"
	"atlas/providers/cashfree"
	"atlas/providers/razorpay"
	"atlas/service/impl"
	"atlas/utils"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, pi *impl.PaymentController, r *impl.RouterController) {
	h := &paymentHandler{
		controller: pi,
		routerCtrl: r,
		cfProvider: &cashfree.Cashfree{
			AppId:  os.Getenv("cashfree_appId"),
			Secret: os.Getenv("cashfree_secret"),
		},
		rzpProvider: &razorpay.Razorpay{
			AppId:  os.Getenv("rzp_appId"),
			Secret: os.Getenv("rzp_secret"),
		},
	}
	e.POST("/process_payment", h.CreatePayment)
	e.POST("/attempt_payment", h.AttemptPayment)
	e.GET("/redirect/:redirectHash", h.RedirectPayment)
}

type paymentHandler struct {
	controller  *impl.PaymentController
	routerCtrl  *impl.RouterController
	cfProvider  provider.Provider
	rzpProvider provider.Provider
}

func (h *paymentHandler) CreatePayment(c echo.Context) error {
	req := new(models.PaymentIntentReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// For example, sending back a simple response
	resp, err := h.controller.Create(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *paymentHandler) AttemptPayment(c echo.Context) error {
	req := new(models.PaymentAttemptReq)
	if err := c.Bind(req); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	// For example, sending back a simple response
	pi, err := h.controller.Get(c.Request().Context(), req.PaymentIntentId)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	// create payment attempt
	pa, err := h.controller.CreateAttempt(c.Request().Context(), pi, req)
	if err != nil {
		fmt.Println("failed in creating payment attempt", pa)
		return c.JSON(http.StatusInternalServerError, err)
	}

	//TODO: inject router module
	provider, err := h.getRoute(req.PaymentMethod)
	fmt.Println("what happened to getRoute? ", provider, err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	html, err := provider.CreatePayment(c.Request().Context(), pi)
	fmt.Println("what happened to createPayment? ", html, err)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	pa.PaymentHash = utils.NewSHA1Hash()
	pa.PaymentProcessor = provider.GetName()
	pa, err = h.controller.UpdateProvider(c.Request().Context(), pa, html)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	fmt.Println("url --> ", html)
	return c.JSON(http.StatusOK, pa)
}

func (h *paymentHandler) getRoute(paymentMethod string) (provider.Provider, error) {
	//TODO: inject router module
	route, err := h.routerCtrl.GetPreference(paymentMethod, make(map[string]interface{}))
	if err != nil {
		return nil, err
	}
	fmt.Println("got back from GetPreference: ", route)
	if route == "cashfree" {
		return h.cfProvider, nil
	} else if route == "razorpay" {
		return h.rzpProvider, nil
	}
	return nil, errors.New("could not find a valid provider from megamind!")
}

func (h *paymentHandler) RedirectPayment(c echo.Context) error {
	redirectHash := c.Param("redirectHash")
	pa, _ := h.controller.GetPaymentAttemptForHash(c.Request().Context(), redirectHash)
	h.controller.PaymentAttemptRedirected(c.Request().Context(), pa.Id)
	return c.Render(http.StatusOK, "redirect.html", map[string]interface{}{
		"HTML_BODY": pa.Metadata,
	})
}
