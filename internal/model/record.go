package model

type Detail struct {
	Type     string
	Time     string
	Duration string
	Quality  uint
	Tags	 []string
	Location string
}

type Record struct {
	ID        string
	Details   []Detail
	Important bool
}
