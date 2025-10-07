package main

import (
	"ecommerce/cmd"
)

func main() {
	cmd.Serve()

	// jwt, err := utils.CreateJwt("hallucination", utils.Payload{
	// 	ID:          1,
	// 	FistName:    "John",
	// 	LastName:    "Doe",
	// 	Email:       "john@doe.com",
	// 	IsShopOwner: false,
	// })
	// if err != nil {
	// 	fmt.Println("Error creating JWT:", err)
	// 	return
	// }
	// fmt.Println("Generated JWT:", jwt)
}
