package domain

type AcceptTOSRequest struct {
	Device
	TOSVersion int `json:"tos_version"`
}
type AcceptTOSResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

type TosStatus struct {
	HasAcceptedCurrentTos bool `json:"hasAcceptedCurrentTos"`
}

type UserTOSAgreement struct {
	ID             int64  `sql:"id,omitempty"`
	UserID         int64  `sql:"user_id,omitempty"`
	BundleID       string `sql:"bundle_id,omitempty"`
	TOSVersion     int64  `sql:"tos_version,omitempty"`
	CreatedAtStamp int64  `sql:"created_at_stamp,omitempty"`
}
