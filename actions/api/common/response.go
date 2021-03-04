package common

// TokenRequest is a request for a token exchange
// swagger:parameters exchangeFacebook exchangeGoogle exchangeAdmin
type TokenRequest struct {
	// in:body
	// required:true
	Body TokenPayload
}

// TokenResponse is a response after a token exchange
// swagger:response tokenResponse
type TokenResponse struct {
	// in: body
	Body TokenPayload
}

// TokenPayload contain the body parameters for a token exchange
type TokenPayload struct {
	Token string `json:"token"`
}
