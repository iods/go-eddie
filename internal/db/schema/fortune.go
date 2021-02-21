package schema

type Fortunes struct {
	Fortunes []Fortune
}

type Fortune struct {
	Id         uint64 `json:"id"`
	Text       string `json:"text"`
	Numbers    []int  `json:"numbers"`
	Color      string `json:"color"`
	Condition  uint64 `json:"condition"`
	Restaurant string `json:"restaurant"`
}

type FortuneResponse interface {
	GetFortune() Fortunes
}