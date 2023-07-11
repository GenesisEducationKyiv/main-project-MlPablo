package e2e

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"

	"github.com/antihax/optional"
	mailslurp "github.com/mailslurp/mailslurp-client-go"
)

func (suite *Suite) TestSendEmail() {
	// create test email
	client, ctx := getMailSlurpClient()
	inbox1, _, _ := client.InboxControllerApi.CreateInbox(ctx, nil)

	// subcribe new email
	rec := httptest.NewRecorder()
	suite.createUser(inbox1.EmailAddress, rec)

	// send emails to subscribed users
	rec = httptest.NewRecorder()
	suite.sendEmails(rec)

	// create wait options for email
	waitOpts := &mailslurp.WaitForLatestEmailOpts{
		InboxId:    optional.NewInterface(inbox1.Id),
		Timeout:    optional.NewInt64(30000),
		UnreadOnly: optional.NewBool(true),
	}
	email, _, err := client.WaitForControllerApi.WaitForLatestEmail(ctx, waitOpts)
	suite.Require().NoError(err)

	suite.Require().NotNil(email.Body)

	// var resp rate.Currency
	// err = json.Unmarshal([]byte(strings.TrimSpace(*email.Body)), &resp)
	btcPrice, err := strconv.ParseFloat(strings.TrimSpace(*email.Body), 64)
	suite.Require().NoError(err, *email.Body)
	suite.NotZero(btcPrice)
}

func getMailSlurpClient() (*mailslurp.APIClient, context.Context) {
	const apiKey = "2168e0bdf90b5195a3ed9ad7c94976e7ead45417315296b1c2f72f51ac386f66"

	// create a context with your api key
	ctx := context.WithValue(
		context.Background(),
		mailslurp.ContextAPIKey,
		mailslurp.APIKey{Key: apiKey},
	)

	// create mailslurp client
	config := mailslurp.NewConfiguration()
	client := mailslurp.NewAPIClient(config)

	return client, ctx
}

func (suite *Suite) sendEmails(w http.ResponseWriter) {
	req := httptest.NewRequest(http.MethodPost, "/api/sendEmails", nil)
	suite.e.ServeHTTP(w, req)
}
