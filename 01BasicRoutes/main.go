package main

import (
	"fmt"
	"time"
)

// type users struct {
// 	Name   string `json:"name"`
// 	Age    int32  `json:"Age"`
// 	Salary int32  `json:"salary"`
// }

func main() {
	// Users := []users{
	// 	{Name: "umapathi", Age: 21, Salary: 25000},
	// 	{Name: "solomon", Age: 22, Salary: 66000},
	// 	{Name: "deena", Age: 22, Salary: 30000},
	// 	{Name: "suresh babu", Age: 21, Salary: 45000},
	// }
	// fmt.Println("Basic Routes")
	// app := fiber.New()
	// app.Use(logger.New())
	// Home := app.Group("/")

	// Home.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(Users)
	// })

	// Home.Get("/about", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(fiber.Map{"Message": "about"})
	// })

	// Home.Get("/contect", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(fiber.Map{"Message": "Contect"})
	// })

	// app.Listen(":3000")
	time3 := time.Now().Format(time.RFC3339)
	fmt.Println(time3)
	

}
