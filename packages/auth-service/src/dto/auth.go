package dto

type (
	RegisterRequest struct {
		Nik  string `json:"nik" binding:"required,min=16,max=16"`
		Role string `json:"role" binding:"required"`
	}

	RegisterResponse struct {
		Nik      string `json:"nik"`
		Role     string `json:"role"`
		Password string `json:"password"`
	}
)
