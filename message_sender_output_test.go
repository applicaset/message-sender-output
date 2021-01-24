package messagesenderoutput_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/applicaset/message-sender-output"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMessageSender_Send(t *testing.T) {
	var output bytes.Buffer

	phoneNumber := "+1234567890"
	msg := "Dummy Text"

	expected := fmt.Sprintf("%s\n\n%s\n", phoneNumber, msg)

	ctx := context.Background()

	ms := messagesenderoutput.New(messagesenderoutput.WithOutput(&output))

	err := ms.Send(ctx, phoneNumber, msg)
	require.NoError(t, err)
	require.Equal(t, expected, output.String())
}
