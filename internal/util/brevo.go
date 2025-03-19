package util

import (
	"log"

	"github.com/go-resty/resty/v2"
)

func SendEmailWithTemplate(brevoApiKey string, templateId int, toEmail string, otp string) error {
	client := resty.New()

	response, err := client.R().
		SetHeader("accept", "application/json").
		SetHeader("api-key", brevoApiKey).
		SetHeader("content-type", "application/json").
		SetBody(map[string]interface{}{
			"to": []map[string]string{
				{"email": toEmail, "name": "User"},
			},
			"templateId": templateId,
			"params": map[string]string{
				"otp": otp,
			},
		}).
		Post("https://api.brevo.com/v3/smtp/email")

	if err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	log.Println("Email sent successfully:", response)
	return nil
}
