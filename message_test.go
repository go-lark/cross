package cross

import (
	"os"
	"testing"

	clark "github.com/chyroc/lark"
	golark "github.com/go-lark/lark/v2"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// IDs for test use
var (
	testAppID     string
	testAppSecret string
	testUserEmail string
)

func loadTestEnv() {
	testMode := os.Getenv("GO_LARK_TEST_MODE")
	if testMode == "" {
		testMode = "testing"
	}
	if testMode == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}
	testAppID = os.Getenv("LARK_APP_ID")
	testAppSecret = os.Getenv("LARK_APP_SECRET")
	testUserEmail = os.Getenv("LARK_USER_EMAIL")
	if len(testAppID) == 0 ||
		len(testAppSecret) == 0 ||
		len(testUserEmail) == 0 {
		panic("insufficient test environment")
	}
}

func TestCrossMsg(t *testing.T) {
	loadTestEnv()
	b := golark.NewCardBuilder()
	card := b.I18N.
		Card(
			b.I18N.WithLocale(
				golark.LocaleEnUS,
				b.Div(
					b.Field(b.Text("English Content")),
				),
			),
			b.I18N.WithLocale(
				golark.LocaleZhCN,
				b.Div(
					b.Field(b.Text("中文内容")),
				),
			),
			b.I18N.WithLocale(
				golark.LocaleJaJP,
				b.Div(
					b.Field(b.Text("日本語コンテンツ")),
				),
			),
		).
		Title(
			b.I18N.LocalizedText(golark.LocaleEnUS, "English Title"),
			b.I18N.LocalizedText(golark.LocaleZhCN, "中文标题"),
			b.I18N.LocalizedText(golark.LocaleJaJP, "日本語タイトル"),
		).
		Red().
		UpdateMulti(true)
	msg := golark.NewMsgBuffer(golark.MsgInteractive)
	om := msg.BindEmail(testUserEmail).Card(card.String()).Build()
	crossMsg, _ := BuildMessage(golark.UIDEmail, om)
	cli := clark.New(clark.WithAppCredential(testAppID, testAppSecret))
	resp, _, err := cli.Message.SendRawMessage(t.Context(), crossMsg)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, resp.MessageID)
	}
}
