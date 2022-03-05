package i2p

import(
	"io"
	"encoding/binary"
	"crypto/rand"

	sam3 "github.com/eyedeekay/sam3"
	i2ptp "github.com/allhailjarjar/go-libp2p-i2p-transport"
	ma "github.com/multiformats/go-multiaddr"
	libp2p "github.com/libp2p/go-libp2p"
	p2pcrypto "github.com/libp2p/go-libp2p-core/crypto"
	host "github.com/libp2p/go-libp2p-core/host"
)

const SAMHost = "localhost:7656"

func getSeed(seeds ...io.Reader) io.Reader{
	if len(seeds) == 0{
		return rand.Reader
	}else{
		return seeds[0]
	}
}

func makeI2pTpBuilder() (i2ptp.TransportBuilderFunc, ma.Multiaddr, error){
	sam, err := sam3.NewSAM(SAMHost)
	if err != nil{return nil, nil, err}
	keys, err := sam.NewKeys()
	if err != nil{return nil, nil, err}

	seedBytes := make([]byte, 8)
	rand.Read(seedBytes)
	seed, _ := binary.Varint(seedBytes)
	return i2ptp.I2PTransportBuilder(sam, keys, "45793", int(seed))
}

func NewI2pHost(seeds ...io.Reader) (host.Host, error){
	tpBuilder, listenAddr, err := makeI2pTpBuilder()
	if err != nil{return nil, err}

	seed := getSeed(seeds...)
	priv, _, _ := p2pcrypto.GenerateEd25519Key(seed)

	return libp2p.New(
		libp2p.Transport(tpBuilder),
		libp2p.ListenAddrs(listenAddr),
		libp2p.Identity(priv),
		libp2p.DefaultSecurity,
		libp2p.ForceReachabilityPublic(),
		libp2p.EnableRelay(),
	)
}
