// package domain

// type Device struct {
// 	ID                      int64                  `sql:"id,omitempty"`
// 	InternationalIndicators InternationalIndicator `sql:",inline"`
// 	DevicePublicFields      DevicePublicFields     `sql:",inline"`
// 	FCMToken                string                 `sql:"fcm_token"`
// 	CreatedAtStamp          int64                  `sql:"created_at_stamp,omitempty"`
// 	UpdatedAtStamp          int64                  `sql:"updated_at_stamp,omitempty"`
// }

// // Device Public fields are used in many places and that's the reason why we separated them for the sake of refactoring
// type DevicePublicFields struct {
// 	UUID           string `sql:"uuid,omitempty" json:"uuid" form:"uuid" header:"UUID"`
// 	Platform       string `sql:"platform,omitempty" json:"platform" form:"platform" header:"Platform"`
// 	Model          string `sql:"model,omitempty" json:"device_model" form:"device_model" header:"Device-Model"`
// 	BundleID       string `sql:"bundle_id,omitempty" json:"bundle_id" form:"bundle_id" header:"Bundle-ID"`
// 	OSVersion      string `sql:"os_version,omitempty" json:"os_version" form:"os_version" header:"OS-Version"`
// 	AppVersion     string `sql:"app_version,omitempty" json:"app_version" form:"app_version" header:"App-Version"`
// 	AppBuildNumber int16  `sql:"app_build_number,omitempty" json:"app_build_number" form:"app_build_number" header:"App-Build-Number"`
// 	AssetVersion   int16  `sql:"asset_version,omitempty" json:"asset_version" form:"asset_version"  header:"Asset-Version"`
// 	IsCRMDevice    bool   `sql:"-" json:"is_crm_device" form:"is_crm_device" header:"Is-CRM-Device"`
// }

package domain

import "time"

type Device struct {
	ID             int64     `json:"id"`
	UserID         int64     `json:"user_id"`
	UUID           string    `json:"uuid"`
	Platform       string    `json:"platform"`
	Model          string    `json:"device_model"`
	BundleID       string    `json:"bundle_id"`
	OSVersion      string    `json:"os_version"`
	AppVersion     string    `json:"app_version"`
	AppBuildNumber int16     `json:"app_build_number"`
	AssetVersion   int16     `json:"asset_version"`
	FCMToken       string    `json:"fcm_token,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	IsCRMDevice    bool      `json:"is_crm_device,omitempty"`
}

// InternationalIndicators includes country, language, and timezone information.
// Consider renaming it more meaningfully, depending on its usage.
type InternationalIndicators struct {
	CountryCode string `json:"country_code"`
	LangCode    string `json:"lang_code"`
	TimeZone    string `json:"time_zone"`
}
