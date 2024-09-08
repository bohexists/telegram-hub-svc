package telegram

type Bot struct {
	Client *Client
}

func NewBot(token string) *Bot {
	return &Bot{
		Client: NewClient(token),
	}
}

func (b *Bot) Start() {
	b.Client.HandleUpdates()
}
