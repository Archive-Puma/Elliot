package env

// Config TODO: Doc
var Config = struct {
	Target string
	Params interface{}
}{}

// Channels TODO: Doc
var Channels struct {
	Ok    chan string
	Bad   chan error
	Done  chan struct{}
	Start chan struct{}
	Stop  chan struct{}
}
