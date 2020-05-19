package env

var (
	// Params TODO: Doc
	Params = struct {
		Target    string
		Plugin    string
		Arguments []interface{}
	}{}
	// Channels TODO: Doc
	Channels = struct {
		Ok   chan string
		Bad  chan error
		Done chan struct{}
	}{
		Ok:   make(chan string),
		Bad:  make(chan error),
		Done: make(chan struct{}),
	}
)
