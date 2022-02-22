package fingers

import "errors"

var (
	ERR_NO_SUCCESSOR  = errors.New("cannot find successor")
	ERR_NODE_EXISTS   = errors.New("node with id already exists")
	ERR_KEY_NOT_FOUND = errors.New("key not found")
)
