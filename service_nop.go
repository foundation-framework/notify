package notify

type nopService struct {
}

func NewNopService() Service {
	return &nopService{}
}

func (n *nopService) Send(text string, attachments ...Attachment) error {
	return nil
}
