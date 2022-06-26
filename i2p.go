package i2p

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"strings"

	sam3 "github.com/eyedeekay/sam3"
	libp2p "github.com/libp2p/go-libp2p"
	p2pcrypto "github.com/libp2p/go-libp2p-core/crypto"
	host "github.com/libp2p/go-libp2p-core/host"
	ma "github.com/multiformats/go-multiaddr"
	i2ptp "github.com/pilinsin/go-libp2p-i2p-transport"
)

const SAMHost = "localhost:7656"
const sam3Err = "Failed to create"

func getSeed(seeds ...io.Reader) io.Reader {
	if len(seeds) == 0 {
		return rand.Reader
	} else {
		return seeds[0]
	}
}

func makeI2pTpBuilder() (i2ptp.TransportBuilderFunc, ma.Multiaddr, error) {
	sam, err := sam3.NewSAM(SAMHost)
	if err != nil {
		return nil, nil, err
	}
	keys, err := sam.NewKeys()
	if err != nil {
		return nil, nil, err
	}

	seedBytes := make([]byte, 8)
	rand.Read(seedBytes)
	seed, _ := binary.Varint(seedBytes)
	return i2ptp.I2PTransportBuilder(sam, keys, "45793", int(seed))
}

type i2pHost struct{
	host.Host
}
func NewI2pHost(seeds ...io.Reader) (host.Host, error) {
	var tpBuilder i2ptp.TransportBuilderFunc
	var listenAddr ma.Multiaddr
	var err error
	for {
		tpBuilder, listenAddr, err = makeI2pTpBuilder()
		if err == nil {
			break
		}
		if strings.HasPrefix(err.Error(), sam3Err) {
			continue
		}
		return nil, err
	}

	seed := getSeed(seeds...)
	priv, _, _ := p2pcrypto.GenerateEd25519Key(seed)

	h, err := libp2p.New(
		libp2p.Transport(tpBuilder),
		libp2p.ListenAddrs(listenAddr),
		libp2p.Identity(priv),
		libp2p.DefaultSecurity,
		libp2p.ForceReachabilityPublic(),
		libp2p.EnableRelay(),
	)
	return &i2pHost{Host: h}, err
}

func (h *i2pHost) Close() error{
	var err error
	pids := h.Network().Peers()
	for _, pid := range pids{
		err0 := h.Network().ClosePeer(pid)
		if err0 != nil{
			err = err0
		}
	}

	return err
}

