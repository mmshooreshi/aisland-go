package domain

// Tone, Mood, Style, and Length could be better defined within AI.go if it's specific to AI.

type AIRequest struct {
	UserInput   string  `json:"user_input"`
	SessionID   string  `json:"session_id"`
	UserID      string  `json:"user_id"`
	Mood        Mood    `json:"mood"`
	Style       Style   `json:"style"`
	Length      Length  `json:"length"`
	Temperature float32 `json:"temperature,omitempty"`
}

type AIResponse struct {
	Value string `json:"value"`
}

type Tone string
type Mood string
type Style string
type Length string

const (
	CasualTone       Tone = "casual"
	FormalTone       Tone = "formal"
	FriendlyTone     Tone = "friendly"
	CalmTone         Tone = "calm"
	EnthusiasticTone Tone = "enthusiastic"
	SincereTone      Tone = "sincere"
	RespectfulTone   Tone = "respectful"
	ExcitingTone     Tone = "exciting"
	SeriousTone      Tone = "serious"
	OptimisticTone   Tone = "optimistic"
)

const (
	CasualMood       Mood = "casual"
	FormalMood       Mood = "formal"
	FriendlyMood     Mood = "friendly"
	CalmMood         Mood = "calm"
	EnthusiasticMood Mood = "enthusiastic"
	SincereMood      Mood = "sincere"
	RespectfulMood   Mood = "respectful"
	ExcitingMood     Mood = "exciting"
	SeriousMood      Mood = "serious"
	OptimisticMood   Mood = "optimistic"
)

const (
	PlainStyle        Style = "plain"
	CreativeStyle     Style = "creative"
	ProfessionalStyle Style = "professional"
	SocialStyle       Style = "social"
	AcademicStyle     Style = "academic"
)

const (
	ShortLength    Length = "short"
	MediumLength   Length = "medium"
	LongLength     Length = "long"
	LongLongLength Length = "long-long"
)
