package hysteria

import (
	"math/rand"
	"time"

	"github.com/sagernet/quic-go"
	"github.com/sagernet/sing/common/logger"
	M "github.com/sagernet/sing/common/metadata"
)

// HopStrategy is the strategy for port hopping.
type HopStrategy uint8

const (
	// HopBoth changes both client port and server port when hopping.
	// When using it, the underlying conn of Hysteria will be special
	// HopPacketConn, which changes UDP connection transparently.
	//
	// By the way, if you set multiple ports but their actually just
	// one port. you can just change client port but not change server port.
	HopBoth HopStrategy = iota

	// HopServer just changes server port when hopping.
	// This is the behavior of mihomo.
	HopServer
)

func LoopUpdateServerPort(done0, done1, done2 <-chan struct{},
	interval time.Duration,
	logger logger.Logger,
	serverAddr M.Socksaddr, serverPorts []uint16,
	quicConn *quic.Conn,
) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	logger.Trace("start hop loop")
	defer logger.Trace("stop hop loop")
	for {
		select {
		case <-done0:
			return
		case <-done1:
			return
		case <-done2:
			return
		case <-ticker.C:
		}

		newAddr := serverAddr.UDPAddr()
		newAddr.Port = int(serverPorts[rand.Intn(len(serverPorts))])
		quicConn.SetRemoteAddr(newAddr)
		logger.Debug("hop addr to: ", newAddr)
	}
}
