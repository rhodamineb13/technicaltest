package model

type CreateRequest struct {
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	IsProject    string `json:"is_project"`
	SelfCapture  string `json:"self_capture"`
	ClientPrefix string `json:"client_prefix"`
	ClientLogo   string `json:"client_logo"`
	Address      string `json:"address,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	City         string `json:"city,omitempty"`
}

type UpdateRequest struct {
	Name         string `json:"name,omitempty"`
	Slug         string `json:"slug,omitempty"`
	IsProject    string `json:"is_project,omitempty"`
	SelfCapture  string `json:"self_capture,omitempty"`
	ClientPrefix string `json:"client_prefix,omitempty"`
	ClientLogo   string `json:"client_logo,omitempty"`
	Address      string `json:"address,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	City         string `json:"city,omitempty"`
}

type QueryResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	IsProject    string `json:"is_project"`
	SelfCapture  string `json:"self_capture"`
	ClientPrefix string `json:"client_prefix"`
	ClientLogo   string `json:"client_logo"`
	Address      string `json:"address,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	City         string `json:"city,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	DeletedAt    string `json:"deleted_at,omitempty"`
}

type ModelResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}
