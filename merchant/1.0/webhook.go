package merchant

import (
	"encoding/json"
	"errors"
	"github.com/rysavyvladan/go-revolut/merchant/1.0/request"
	"net/http"
)

type WebhookService struct {
	apiKey string
}

type WebhookUrl struct {
	// call back endpoint of the client system, https is the supported protocol
	Url string `json:"url"`
}

type WebhookResp struct {
	// Order ID of a completed order
	OrderId string `json:"order_id"`
}

// Set:
// doc: https://revolut-engineering.github.io/api-docs/merchant-api/#backend-api-backend-api-webhooks-set-or-revoke-webhook-url
func (w *WebhookService) Set(webhookReq *WebhookUrl) error {

	resp, statusCode, err := request.New(request.Config{
		Method:      http.MethodPost,
		Url:         "https://merchant.revolut.com/api/1.0/webhooks",
		ApiKey:      w.apiKey,
		Body:        webhookReq,
		ContentType: request.ContentType_APPLICATION_JSON,
	})
	if err != nil {
		return err
	}
	if statusCode != http.StatusNoContent {
		return errors.New(string(resp))
	}

	return nil
}

// List:
// doc: https://revolut-engineering.github.io/api-docs/merchant-api/#backend-api-backend-api-webhooks-retrieve-webhooks
func (w *WebhookService) List() ([]*WebhookUrl, error) {

	resp, statusCode, err := request.New(request.Config{
		Method: http.MethodGet,
		Url:    "https://merchant.revolut.com/api/1.0/webhooks",
		ApiKey: w.apiKey,
	})
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, errors.New(string(resp))
	}

	r := []*WebhookUrl{}
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return r, nil
}

//// Delete: Use this API request to delete a web-hook
//// doc: https://revolut-engineering.github.io/api-docs/business-api/#web-hooks-setting-up-a-web-hook
//func (w *WebhookService) Delete() error {
//
//	resp, statusCode, err := request.New(request.Config{
//		Method:      http.MethodDelete,
//		Url:         "https://merchant.revolut.com/api/1.0/webhook",
//		AccessToken: p.accessToken,
//		Sandbox:     p.sandbox,
//	})
//	if err != nil {
//		return err
//	}
//	if statusCode != http.StatusNoContent {
//		return errors.New(string(resp))
//	}
//
//	return nil
//}
