package entity

type Currency struct {
	TokenFrom string `json:"token_from"`
	TokenTo   string `json:"token_to"`
}

type DiffCourse struct {
	DifCourse10  int `json:"dif_course_10"`
	DiffCourse60 int `json:"diff_course_60"`
}
