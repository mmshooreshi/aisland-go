package domain

import "time"

// Subscription represents a user's subscription.
type Subscription struct {
	ProductIdentifier string            `json:"product_identifier"`
	PurchaseDate      time.Time         `json:"purchase_date"`
	ExpiresDate       time.Time         `json:"expires_date"`
	IsActive          bool              `json:"is_active"`
	EntitlementId     string            `json:"entitlement_id"`
	FreeTrialDuration FreeTrialDuration `json:"free_trial_duration"`
}

// Entitlement represents provisioned features or products.
type Entitlement struct {
	ProductIdentifier string    `json:"product_identifier"`
	PurchaseDate      time.Time `json:"purchase_date"`
	ExpiresDate       time.Time `json:"expires_date"`
}

// FreeTrialDuration specifies the duration of a free trial.
type FreeTrialDuration struct {
	Seconds  int    `json:"seconds"`
	Readable string `json:"readable"`
}
