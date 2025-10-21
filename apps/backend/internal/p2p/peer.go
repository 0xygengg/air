package p2p

import (
	"fmt"
	"net"
	"sync"
)

var (
	Peers     = make(map[string]net.Conn)
	peerMutex = sync.Mutex{}
)

func AddPeer(addr string, conn net.Conn) {
	peerMutex.Lock()
	defer peerMutex.Unlock()
	Peers[addr] = conn
	fmt.Printf("Connected peer: %s\n", addr)
}

func RemovePeer(addr string) {
	peerMutex.Lock()
	defer peerMutex.Unlock()
	delete(Peers, addr)
	fmt.Printf("Disconnected peer: %s\n", addr)
}

func Broadcast(msg []byte) {
	peerMutex.Lock()
	defer peerMutex.Unlock()

	for addr, conn := range Peers {
		_, err := conn.Write(append(msg, '\n'))
		if err != nil {
			fmt.Printf("Failed to send %s: %v\n", addr, err)
		}
	}
}
