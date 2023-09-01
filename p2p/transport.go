package p2p

// Peer is a node in a network.
type Peer interface {
	Close() error
}

// Transport handles communication with a peer in a network.
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
