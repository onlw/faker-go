package model

type Person struct {
	FirstName string `json:"first_name" xml:"first_name"`
	LastName  string `json:"last_name" xml:"last_name"`
	Gender    string `json:"gender" xml:"gender"`
	//SSN        string          `json:"ssn" xml:"ssn"`
	Image string `json:"image" xml:"image"`
	Hobby string `json:"hobby" xml:"hobby"`
	//Job        *JobInfo        `json:"job" xml:"job"`
	//Address    *AddressInfo    `json:"address" xml:"address"`
	//Contact    *ContactInfo    `json:"contact" xml:"contact"`
	//CreditCard *CreditCardInfo `json:"credit_card" xml:"credit_card"`
}
