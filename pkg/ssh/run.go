package ssh

func (c *Client) Run(cmd string) (string, error) {
	client, err := c.connect()
	if err != nil {
		return "", err
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	result, err := session.Output(cmd)

	return string(result), err
}
