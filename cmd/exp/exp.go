package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// // Read config info from .env
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// host := os.Getenv("SMTP_HOST")
	// portStr := os.Getenv("SMTP_PORT")
	// port, err := strconv.Atoi(portStr)
	// if err != nil {
	// 	panic(err)
	// }
	// username := os.Getenv("SMTP_USERNAME")
	// password := os.Getenv("SMTP_PASSWORD")

	// // email := models.Email{
	// // 	From:      "test@lenslocked.com",
	// // 	To:        "hello@lenslocked.com",
	// // 	Subject:   "this is a test email",
	// // 	Plaintext: "test email body in plaintext",
	// // 	HTML:      `<h1>hi there!</h1><p>test email body in html</p>`,
	// // }

	// es := models.NewEmailService(models.SMTPConfig{
	// 	Host:     host,
	// 	Port:     port,
	// 	Username: username,
	// 	Password: password,
	// })

	// err = es.ForgotPassword("test1@example.com", "https://localhost:3000/reset-password?token=abc123")

	// // err = es.Send(email)
	// // if err != nil {
	// // 	panic(err)
	// // }
	// fmt.Println("Reset password email sent!")

	fmt.Println("On windows:")
	fmt.Println(filepath.Base("/foo.js"))
	fmt.Println(filepath.Base("/foo/bar.js"))
	fmt.Println(filepath.Base("/foo/bar/baz.js"))
	fmt.Println(filepath.Base("dev.txt"))
	fmt.Println(filepath.Base("../todo.txt"))
	fmt.Println(filepath.Base(".."))
	fmt.Println(filepath.Base("."))
	fmt.Println(filepath.Base("/galleries/new"))
	fmt.Println(filepath.Base(""))
}
