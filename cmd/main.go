package main

type Account struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	CPF         string `json:"cpf"`
	CarPlate    string `json:"carPlate"`
	IsPassenger bool   `json:"isPassenger"`
	IsDriver    bool   `json:"isDriver"`
	Password    string `json:"password"`
}

func main() {
	StartServer()
}
