package impl

import (
	"atlas/ent"
	"atlas/ent/paymentattempt"
	"atlas/models"
	"context"
	"fmt"
)

type PaymentController struct {
	/* for now no repo
	paymentIntentRepo
	paymentAttemptRepo
	*/
	entClient *ent.Client
}

func InitPaymentController(c *ent.Client) *PaymentController {
	p := &PaymentController{entClient: c}
	return p
}

func (p *PaymentController) Create(c context.Context, req *models.PaymentIntentReq) (*models.PaymentIntent, error) {
	pi, err := p.entClient.PaymentIntent.Create().
		SetCustomerEmail(req.CustomerEmail).
		SetCustomerPhone(req.CustomerPhone).
		SetCustomerAddress("test address").
		SetStatus("CREATED").
		SetAmount(float64(req.Amount)).
		SetCurrency(req.Currency).
		Save(c)

	if err != nil {
		fmt.Println("something went wrong while persistence!", err)
		return nil, err
	}
	return models.ConvertEntToModel(*pi), nil
}

func (p *PaymentController) Get(c context.Context, id int) (*models.PaymentIntent, error) {
	pi, err := p.entClient.PaymentIntent.Get(c, id)
	return models.ConvertEntToModel(*pi), err
}

func (p *PaymentController) CreateAttempt(c context.Context, pi *models.PaymentIntent, req *models.PaymentAttemptReq) (*models.PaymentAttempt, error) {
	pa, err := p.entClient.PaymentAttempt.Create().
		SetPaymentIntentID(pi.Id).
		SetAmount(float64(pi.Amount)).
		SetCurrency(pi.Currency).
		SetStatus("CREATED").
		SetPaymentMethod(req.PaymentMethod).
		Save(c)
	if err != nil {
		fmt.Println("something went wrong while persistence!", err)
		return nil, err
	}
	return models.ConvertEntToPA(*pa), nil
}

func (p *PaymentController) UpdateProvider(c context.Context, pa *models.PaymentAttempt, metadata string) (*models.PaymentAttempt, error) {
	updated, err := p.entClient.PaymentAttempt.
		UpdateOneID(pa.Id).
		SetStatus("INITIATED").
		SetProcessor(pa.PaymentProcessor).
		SetMetadata(metadata).
		SetPaymentHash(pa.PaymentHash).
		Save(c)
	if err != nil {
		fmt.Println("failure in updating payment attempt", err)
		return nil, err
	}
	return models.ConvertEntToPA(*updated), nil
}

func (p *PaymentController) GetPaymentAttemptForHash(c context.Context, hash string) (*models.PaymentAttempt, error) {
	pa, err := p.entClient.PaymentAttempt.Query().Where(paymentattempt.PaymentHashEQ(hash)).First(c)
	return models.ConvertEntToPA(*pa), err
}

func (p *PaymentController) PaymentAttemptRedirected(c context.Context, id int) (*models.PaymentAttempt, error) {
	updated, err := p.entClient.PaymentAttempt.UpdateOneID(id).SetStatus("REDIRECTED").Save(c)
	if err != nil {
		fmt.Println("failure in updating payment attempt after redirection", err)
		return nil, err
	}
	return models.ConvertEntToPA(*updated), nil
}
