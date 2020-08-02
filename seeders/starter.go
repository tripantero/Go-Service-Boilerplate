package seeders

// Starter return dummy JSON response to test the api
func Starter() interface{} {
	return struct {
		Status string        `json:"status"`
		Data   []interface{} `json:"data"`
	}{Status: "ok"}
}
