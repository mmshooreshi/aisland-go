package domain

type GetConfigRequest struct {
	Device
}

type GetConfigResponse struct {
	Config Config `json:"config"`
}

type Config struct {
	Links                Links         `yaml:"links" json:"links"`
	Notifications        Notifications `yaml:"notifications" json:"notifications"`
	Updates              Updates       `yaml:"updates" json:"updates"`
	Auth                 Auth          `yaml:"authentication" json:"authentication"`
	Suggestions          []Suggestion  `yaml:"askSuggestions" json:"askSuggestions"`
	ImproveCategories    []Category    `yaml:"improveCategories"  json:"improveCategories"`
	ToneChangeCategories []Category    `yaml:"toneChangeCategories"  json:"toneChangeCategories"`
}

// Auth represents authentication configuration.
type Auth struct {
	OTPEnabled bool `yaml:"otpAuthenticationEnabled" json:"otpAuthenticationEnabled"`
}

type Links struct {
	FAQ           string `yaml:"faqUrl" json:"faqUrl"`
	PrivacyPolicy struct {
		Default string `yaml:"privacyPolicyUrl" json:"privacyPolicyUrl"`
		InApp   string `yaml:"privacyPolicyInAppUrl"  json:"privacyPolicyInAppUrl"`
	} `yaml:"privacyPolicy"  json:"privacyPolicy"`
	TermsOfService struct {
		Version int    `yaml:"termsOfServiceVersion"  json:"termsOfServiceVersion"`
		Default string `yaml:"termsOfServiceUrl"  json:"termsOfServiceUrl"`
		InApp   string `yaml:"termsOfServiceInAppUrl"  json:"termsOfServiceInAppUrl"`
	} `yaml:"termsOfService"  json:"termsOfService"`
	Update string `yaml:"updateUrl"  json:"updateUrl"`
}

// Updates includes fields related to application updates.
type Updates struct {
	SoftUpdateAvailable  bool `yaml:"softUpdateAvailable" json:"softUpdateAvailable"`
	ForcedUpdateRequired bool `yaml:"forcedUpdateRequired" json:"forcedUpdateRequired"`
}

// Suggestion represents user interface suggestions.
type Suggestion struct {
	Text string `json:"text"`
}

// Category represents various action categories.
type Category struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Notifications struct {
	FCMTopics []string `yaml:"fcmTopics"`
}
