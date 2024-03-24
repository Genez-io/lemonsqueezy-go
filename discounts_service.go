package lemonsqueezy

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// DiscountsService is the API client for the `/v1/discounts` endpoint
type DiscountsService service

// Create a discount.
//
// https://docs.lemonsqueezy.com/api/discounts#create-a-discount
func (service *DiscountsService) Create(ctx context.Context, params *DiscountCreateParams) (*DiscountAPIResponse, *Response, error) {
	payload := map[string]any{
		"data": map[string]any{
			"type": "discounts",
			"attributes": map[string]any{
				"name":        params.Name,
				"code":        params.Code,
				"amount":      params.Amount,
				"amount_type": params.AmountType,
			},
			"relationships": map[string]any{
				"store": map[string]any{
					"data": map[string]any{
						"type": "stores",
						"id":   strconv.Itoa(params.StoreID),
					},
				},
			},
		},
	}

	response, err := service.client.do(ctx, http.MethodPost, "/v1/discounts/", payload)
	if err != nil {
		return nil, response, err
	}

	discount := new(DiscountAPIResponse)
	if err = json.Unmarshal(*response.Body, discount); err != nil {
		return nil, response, err
	}

	return discount, response, nil
}

// Get the discount with the given ID.
//
// https://docs.lemonsqueezy.com/api/discounts#retrieve-a-discount
func (service *DiscountsService) Get(ctx context.Context, discountID string) (*DiscountAPIResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/discounts/"+discountID)
	if err != nil {
		return nil, response, err
	}

	discount := new(DiscountAPIResponse)
	if err = json.Unmarshal(*response.Body, discount); err != nil {
		return nil, response, err
	}

	return discount, response, nil
}

// Delete a discount with the given ID.
//
// https://docs.lemonsqueezy.com/api/discounts#delete-a-discount
func (service *DiscountsService) Delete(ctx context.Context, discountID string) (*Response, error) {
	return service.client.do(ctx, http.MethodDelete, "/v1/discounts/"+discountID)
}

// List returns a paginated list of discounts.
//
// https://docs.lemonsqueezy.com/api/discounts#list-all-discounts
func (service *DiscountsService) List(ctx context.Context) (*DiscountsAPIResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/discounts")
	if err != nil {
		return nil, response, err
	}

	discounts := new(DiscountsAPIResponse)
	if err = json.Unmarshal(*response.Body, discounts); err != nil {
		return nil, response, err
	}

	return discounts, response, nil
}
