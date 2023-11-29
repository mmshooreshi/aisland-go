package domain

import (
	"time"
)

// User represents the end user.
type User struct {
	ID                int64         `json:"id,omitempty"`
	Username          string        `json:"username"`
	Email             string        `json:"email"`
	FirstSeen         time.Time     `json:"first_seen"`
	LastSeen          time.Time     `json:"last_seen"`
	OriginalAppUserID string        `json:"original_app_user_id"`
	Devices           []Device      `json:"devices"`
	Entitlements      []Entitlement `json:"entitlements"`
	Subscription      Subscription  `json:"subscription"`
}

type UserInternalRevenueCatID struct {
	ID             int64  `sql:"id,omitempty"`
	UserID         int64  `sql:"user_id,omitempty"`
	RevenueCatID   string `sql:"revenuecat_id,omitempty"`
	CreatedAtStamp int64  `sql:"created_at_stamp,omitempty"`
	UpdatedAtStamp int64  `sql:"updated_at_stamp,omitempty"`
}
