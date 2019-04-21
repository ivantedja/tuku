package delivery

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/ivantedja/tuku"
	"github.com/ivantedja/tuku/config"

	"github.com/nlopes/slack"
)

type tukuSlackRtm struct {
	Config         config.Config
	UserUsecase    tuku.UserUsecase
	DepositUsecase tuku.DepositUsecase
}

func NewTukuSlackRtm(cfg config.Config, uu tuku.UserUsecase, du tuku.DepositUsecase) tukuSlackRtm {
	return tukuSlackRtm{cfg, uu, du}
}

func (tsr tukuSlackRtm) ListenAndServe() error {
	api := slack.New(
		tsr.Config.Slack.Token,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	rtm := api.NewRTM()

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)

		case *slack.MessageEvent:
			fmt.Printf("Message Event: %+v\n", ev)
			tsr.RegisterHandler(tsr.Config, rtm, ev)

		case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())
			return errors.New(ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return errors.New("Invalid credentials")

		default:
			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}

	return nil
}

func (tsr tukuSlackRtm) RegisterHandler(cfg config.Config, rtm *slack.RTM, ev *slack.MessageEvent) {
	regexGetUser := regexp.MustCompile(`^tuku getuser (\d+)`)
	if matches := regexGetUser.FindStringSubmatch(ev.Text); len(matches) > 1 {
		userId, err := strconv.ParseInt(matches[1], 10, 64)
		if err != nil {
			rtm.SendMessage(rtm.NewOutgoingMessage("Invalid User", cfg.Slack.ChannelID))
		}
		tsr.GetUser(cfg, rtm, userId)
	}
}

func (tsr tukuSlackRtm) GetUser(cfg config.Config, rtm *slack.RTM, userId int64) {
	user, err := tsr.UserUsecase.GetUser(userId)
	if err != nil {
		rtm.SendMessage(rtm.NewOutgoingMessage("Invalid user", cfg.Slack.ChannelID))
	}
	rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("%+v", user), cfg.Slack.ChannelID))
}
