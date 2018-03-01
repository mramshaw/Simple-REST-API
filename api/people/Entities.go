package people

// The Person entity is used to marshall/unmarshall JSON.
type Person struct {
	ID        string   `json:"id,       omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname, omitempty"`
	Address   *Address `json:"address,  omitempty"`
}

// The Address entity is used to marshall/unmarshall JSON.
type Address struct {
	City  string `json:"city, omitempty"`
	State string `json:"state,omitempty"`
}
