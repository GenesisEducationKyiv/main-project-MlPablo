package e2e

type eventer struct{}

func (e *eventer) Publish(topic, body string) error {
	return nil
}
