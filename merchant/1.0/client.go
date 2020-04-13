package merchant

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

func (m *Client) Order() *OrderService {
	return &OrderService{
		apiKey: m.apiKey,
	}
}

func (m *Client) Webhook() *WebhookService {
	return &WebhookService{
		apiKey: m.apiKey,
	}
}
