package domain

// InternationalIndicator includes country, language, and timezone information.
type InternationalIndicator struct {
	CountryCode string `json:"country_code"`
	LangCode    string `json:"lang_code"`
	TimeZone    string `json:"time_zone"`
}
