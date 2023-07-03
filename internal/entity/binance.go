package entity

type Currency struct {
	TokenFrom string `json:"token_from"`
	TokenTo   string `json:"token_to"`
}

type DiffCourse struct {
	Now         float64 `json:"now"`
	DifCourse24 float64 `json:"dif_course_24"`
}
