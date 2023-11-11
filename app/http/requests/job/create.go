package job

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type Create struct {
	Link    string `json:"link" form:"link"`
	Role    string `json:"role" form:"role"`
	Company string `json:"company" form:"company"`
}

func (r *Create) Authorize(ctx http.Context) error {
	return nil
}

func (r *Create) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"link":    "required|url",
		"role":    "required|min_length:8",
		"company": "required|min_length:8",
	}
}

func (r *Create) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *Create) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *Create) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
