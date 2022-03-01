package fingers

import "errors"

var (
	errNoSuccessor = errors.New("cannot find successor")
	errNodeExists  = errors.New("node with id already exists")
	errKeyNotFound = errors.New("key not found")
)
