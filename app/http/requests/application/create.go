package application

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"mgufrone.dev/job-tracking/app/http/requests/job"
)

type Create struct {
	Job        job.Create `json:"job"`
	ResumeName string     `json:"resume_name"`
}

func (r *Create) Authorize(ctx http.Context) error {
	return nil
}

func (r *Create) Rules(ctx http.Context) map[string]string {
	jobRules := r.Job.Rules(ctx)
	rules := map[string]string{
		"resume_name": "required",
	}
	for k, v := range jobRules {
		rules[fmt.Sprint("job.", k)] = v
	}
	return rules
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
