package models

// RequestBody
type RequestBody struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

// ResponseBody
type ResponseBody struct {
	Solutions [][]int `json:"solutions"`
}
