package p2p

// Peer is an interface that represents the remote node in the network
type Peer interface{}

// Transport is anything that handeles the communication
// betwen the nodes in the network. This can be of the
// form of a TCP connection, a UDP connection, websockets...
type Transport interface {
	ListenAndAccept() error
}
