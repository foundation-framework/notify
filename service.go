package notify

type Service interface {
	Send(text string, attachments ...Attachment) error
}

func CombineServices(providers ...Service) Service {
	return &serviceGroup{
		services: providers,
	}
}

type serviceGroup struct {
	services []Service
}

func (p *serviceGroup) Send(text string, attachments ...Attachment) error {
	for _, service := range p.services {
		if err := service.Send(text, attachments...); err != nil {
			return err
		}
	}

	return nil
}
