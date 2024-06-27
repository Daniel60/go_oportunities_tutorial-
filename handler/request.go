package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)

}

type CreateOpeningRequest struct {
	Role     string `json: "role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"`
	Link     string `json: "link"`
	Salary   string `json: "salaty"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Link == "" && r.Location == "" && r.Remote == nil && r.Salary == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Salary == "" {
		return errParamIsRequired("salary", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	return nil
}

type ModifyOpeningRequest struct {
	Role     string `json: "role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"`
	Link     string `json: "link"`
	Salary   string `json: "salaty"`
}

func (r *ModifyOpeningRequest) Validate() error {
	//If any field is provided, validation is truthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary != "" {
		return nil
	}
	// If none of the field were provided, return falsy
	return fmt.Errorf("at least one valid field  must be provided")
}
