package fingers

type Operations struct{}

type KV struct {
	Key   []byte
	Value []byte
}

func (o *Operations) Echo(req *KV, resp *KV) error {
	resp = &KV{
		Key:   req.Key,
		Value: req.Value,
	}

	return nil
}

// func (o *Operations) GetSuccessor(_ *Args, reply *Node) error {
// 	return nil
// }

// func (o *Operations) FindSuccessor(_ *Args, reply *Node) error {
// 	return nil
// }

// func (o *Operations) GetPredecessor(_ *Args, reply *Node) error {
// 	return nil
// }

// func (o *Operations) Notify(_ *Args, reply *Node) error {
// 	return nil
// }

// func (o *Operations) CheckPredecessor(_ *Args, reply *Node) error {
// 	return nil
// }

// func (o *Operations) SetPredecessor(_ *Args, pred *Node) error {
// 	return nil
// }

// func (o *Operations) SetSuccessor(_ *Args, succ *Node) error {
// 	return nil
// }

// func (o *Operations) GetKey(k []byte, reply *KV) error {
// 	return nil
// }

// func (o *Operations) SetKey(with, _ *KV) error {
// 	return nil
// }

// func (o *Operations) DeleteKey(k []byte, _ *KV) error {
// 	return nil
// }

// func (o *Operations) RequestKeys(k []byte, reply *[]KV) error {
// 	return nil
// }

// func (o *Operations) DeleteKeys(keys [][]byte, reply *[]KV) error {
// 	return nil
// }
