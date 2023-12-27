package email

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"gopkg.in/gomail.v2"

	//"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

//verificar no aws ses o email que ira enviar os email

func CreateSession(region, awsPK, awsSK string) *ses.SES {
	session, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				awsPK,
				awsSK,
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}

	sesMail := ses.New(session)

	return sesMail
}

type SeSMailSender struct {
	SES       *ses.SES
	Wait      *sync.WaitGroup
	ErrorChan chan error
	//DoneChan  chan bool
}

func NewSesMailSender(sesMail *ses.SES, wait *sync.WaitGroup, errChan chan error) *SeSMailSender {
	return &SeSMailSender{
		SES: sesMail,
		Wait: wait,
		ErrorChan: errChan,
	}
}

func (s *SeSMailSender) sendMail(toAddres, htmlBody, textBody, subject string) error {
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(toAddres),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(htmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(textBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(os.Getenv("EMAIL_SENDER")),
	}

	_, err := s.SES.SendEmail(input)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *SeSMailSender) SendInvoiceMail(toAddress, textBody string) {
	defer s.Wait.Done()
	html := "<p>Sua Nota fiscal esta no pdf em anexo</p>"
	err := s.sendInvoicePDF(toAddress, html, textBody, "nota fiscal referente a compra")
	if err != nil {
		s.ErrorChan <- err
	}
	return 
}

func(s *SeSMailSender) ListenForMail() {
	for {
		select {
		case errMsg := <- s.ErrorChan:
			log.Println(errMsg)
		//case <- s.DoneChan:
		//	return
		}
	}
}

func(s *SeSMailSender) sendInvoicePDF(toAddres, htmlBody, textBody, subject string) error {
	source := aws.String(os.Getenv("EMAIL_SENDER"))
	destination := []*string{&toAddres}

	msg := gomail.NewMessage()
	msg.SetHeader("From", os.Getenv("EMAIL_SENDER"))
	msg.SetHeader("To", toAddres)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", htmlBody)

	//substituir por bucket aws
	data,err := os.ReadFile(os.Getenv("PDF_INVOICE_PATH"))
	if err != nil {
		log.Println(err)
		return err
	}
	msg.Attach("invoice.pdf", gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(data)
		log.Println(err)
		return err
	}))

	var emailRaw bytes.Buffer
	msg.WriteTo(&emailRaw)

	message := &ses.SendRawEmailInput{
		Source: source,
		Destinations: destination,
		RawMessage: &ses.RawMessage{
			Data: emailRaw.Bytes(),
		},
	}
	_,err = s.SES.SendRawEmail(message)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}