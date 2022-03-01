package fingers

type Operations struct{}

func (o Operations) GetSuccessor() (*Node, error) {
	return nil, nil
}

func (o Operations) FindSuccessor([]byte) (*Node, error) {
	return nil, nil
}

func (o Operations) GetPredecessor() (*Node, error) {
	return nil, nil
}

func (o Operations) Notify(*Node) error {
	return nil
}

func (o Operations) CheckPredecessor(*Node) error {
	return nil
}

func (o Operations) SetPredecessor(*Node) error {
	return nil
}

func (o Operations) SetSuccessor(succ *Node) error {
	return nil
}

func (o Operations) GetKey(string) (*GetResponse, error) {
	return nil, nil
}

func (o Operations) SetKey(string, string) error {
	return nil
}

func (o Operations) DeleteKey(string) error {
	return nil
}

func (o Operations) RequestKeys([]byte, []byte) ([]*KV, error) {
	return nil, nil
}

func (o Operations) DeleteKeys([]string) error {
	return nil
}
