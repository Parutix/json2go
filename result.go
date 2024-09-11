type Person struct {
	IsActive    bool `json:"isActive"`
	Preferences struct {
		Newsletter    bool     `json:"newsletter"`
		Notifications []string `json:"notifications"`
	} `json:"preferences"`
	Projects []struct {
		Name      string  `json:"name"`
		Budget    float64 `json:"budget"`
		Completed bool    `json:"completed"`
	} `json:"projects"`
	Name    string  `json:"name"`
	Age     float64 `json:"age"`
	Email   string  `json:"email"`
	Address struct {
		Street  string `json:"street"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
	} `json:"address"`
	PhoneNumbers []struct {
		Number string `json:"number"`
		Type   string `json:"type"`
	} `json:"phoneNumbers"`
}
