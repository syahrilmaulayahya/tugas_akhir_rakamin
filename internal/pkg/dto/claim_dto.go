package dto

import "github.com/google/uuid"

type ClaimJwt struct {
	JwtID          uuid.UUID `json:"jwt_id"`
	Subject        uint      `json:"subject"`
	Issuer         string    `json:"issuer"`
	Audience       string    `json:"audience"`
	Scope          string    `json:"scope"`
	Type           string    `json:"type"`
	IssuedAt       int64     `json:"issued_at"`
	NotValidBefore int64     `json:"not_valid_before"`
	ExpiredAT      int64     `json:"expired_at"`
}
