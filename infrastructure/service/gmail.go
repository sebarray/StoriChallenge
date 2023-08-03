package service


import (
	"os"
	"net/smtp"
	"storie/pkg/handler/helpers"
	"storie/pkg/domain"
)

type Gmail struct {



}



func sendHTMLEmail( recipientEmail, subject, htmlBody string) error {


	senderEmail :=  os.Getenv("sendMail")
	senderPassword := os.Getenv("pswmail")
	// Configura la información del servidor SMTP
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)

	// Crea el mensaje
	to := []string{recipientEmail}
	msg := []byte("To: " + recipientEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		htmlBody + "\r\n")

	// Envía el correo electrónico
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, senderEmail, to, msg)
	if err != nil {
		return err
	}

	return nil
}

func (g Gmail) 	Send(mail domain.Mail) error{
	
	recipientEmail := mail.SenderEmail
	htmlBody :=helpers.SendHtml(mail)
	subject:= "Credit card statement"
	// Envía el correo electrónico
	err:= sendHTMLEmail( recipientEmail, subject, htmlBody)
	return err


	
}