package messagesenderoutput

import (
	"context"
	"github.com/applicaset/sms-svc"
	"github.com/pkg/errors"
	"io"
	"os"
	"text/template"
)

const DefaultTemplate = `{{ .PhoneNumber }}

{{ .Message }}
`

type messageSender struct {
	output io.Writer
	tpl    *template.Template
}

func (ms *messageSender) Send(ctx context.Context, phoneNumber, message string) error {
	data := struct {
		PhoneNumber string
		Message     string
	}{
		PhoneNumber: phoneNumber,
		Message:     message,
	}

	err := ms.tpl.Execute(ms.output, data)
	if err != nil {
		return errors.Wrap(err, "error on execute template")
	}

	return nil
}

type Option func(sender *messageSender)

func WithTemplate(tpl *template.Template) Option {
	return func(ms *messageSender) {
		ms.tpl = tpl
	}
}

func WithOutput(output io.Writer) Option {
	return func(ms *messageSender) {
		ms.output = output
	}
}

func New(options ...Option) smssvc.MessageSender {
	ms := messageSender{
		output: os.Stdout,
		tpl:    template.Must(template.New("_").Parse(DefaultTemplate)),
	}

	for i := range options {
		options[i](&ms)
	}

	return &ms
}
