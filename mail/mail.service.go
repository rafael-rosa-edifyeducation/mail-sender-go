package mails

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type MailSendBody struct {
	To       string `json:"to"`
	Template string `json:"template"`
	Data     string `json:"data"`
}

func HandleSendMail(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Erro ao ler o corpo da requisição", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// Criando uma variável da struct
		var mailBody MailSendBody

		// Desserializando JSON para a struct
		err = json.Unmarshal(body, &mailBody)

		if err != nil {
			http.Error(w, "Erro ao ler o corpo da requisição", http.StatusInternalServerError)
			return
		}

		err = sendEmail(mailBody)

		if err != nil {
			http.Error(w, "Falha ao enviar email", http.StatusInternalServerError)
			return
		}
	default:
		fmt.Println("404")
	}
}

func sendEmail(teste MailSendBody) error {
	config := &aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			"localstack", // Access Key
			"localstack", // Secret Key
			"",           // Session Token (opcional, pode ser "")
		),
		Endpoint: aws.String("http://localhost:4566"),
	}
	sess := session.Must(session.NewSession(config))
	svc := ses.New(sess)
	from := "rafaelcruzrosa1@gmail.com"
	input := &ses.SendTemplatedEmailInput{
		Source:   &from,
		Template: &teste.Template,
		Destination: &ses.Destination{
			ToAddresses: []*string{&teste.To},
		},
		TemplateData: &teste.Data,
	}
	_, err := svc.SendTemplatedEmail(input)

	return err
}
