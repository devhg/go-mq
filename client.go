package gomq

type Client struct {
	broker Broker
}

func NewClient() *Client {
	return &Client{broker: NewBroker()}
}

func (c *Client) Publish(topics string, msg interface{}) error {
	return c.broker.publish(topics, msg)
}

func (c *Client) Subscribe(topic string) (<-chan interface{}, error) {
	return c.broker.subscribe(topic)
}

func (c *Client) Unsubscribe(topic string, sub <-chan interface{}) error {
	return c.broker.unsubscribe(topic, sub)
}

func (c *Client) Close() {
	c.broker.close()
}

func (c *Client) SetConditions(cap int) {
	c.broker.setConditions(cap)
}


func (c Client) GetPayload(sub <-chan interface{}) interface{} {
	for val := range sub {
		if val != nil {
			return val
		}
	}
	return nil
}
