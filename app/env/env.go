package env

// Config serves as a proxy to transfer information to the plugins
var Config = struct {
	Target string
	Params interface{}
}{}

// Channels collects the most important communication channels between the goroutines and the main thread
var Channels struct {
	Ok    chan string
	Bad   chan error
	Done  chan struct{}
	Start chan struct{}
	Stop  chan struct{}
}
