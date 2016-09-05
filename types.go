package main

type Meshcraft struct {
	Auth		Auth		`json:"auth"`
	Application	Application	`json:"application"`
}

type Auth struct {
	TSPName	string	`json:"tsp_name,omitempty" `
	TSPUrl	string	`json:"tsp_url,omitempty" `
	Email	string	`json:"email,omitempty" `
	APIKey	string	`json:"api_key,omitempty" `
}

type Application struct {
	Name	string	`json:"anem" `
	Version	string	`json:"version" `
	Uid	string	`json:"uid" `
}

func (mcraft *Meshcraft) Validate() error {

	return nil
}