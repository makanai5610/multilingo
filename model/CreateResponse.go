package model

type CreateResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (cr *CreateResponse) Log() {}
