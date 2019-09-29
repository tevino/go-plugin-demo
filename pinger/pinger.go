package pinger

// Pinger is the interface of a plugin.
type Pinger interface {
	Ping() error
}
