package chain

type Chain struct {
	ID   uint
	Name string
}

var (
	ArbitrumOne     Chain = Chain{ID: 42161, Name: "Arbitrum One"}
	Aurora          Chain = Chain{ID: 1313161554, Name: "Aurora Mainnet"}
	AvalancheCChain Chain = Chain{ID: 43114, Name: "Avalanche C-Chain"}
	Base            Chain = Chain{ID: 8453, Name: "Base"}
	BSC             Chain = Chain{ID: 56, Name: "BNB Smart Chain Mainnet"}
	BTTC            Chain = Chain{ID: 199, Name: "BitTorrent Chain Mainnet"}
	Cronos          Chain = Chain{ID: 25, Name: "Cronos Mainnet"}
	Ethereum        Chain = Chain{ID: 1, Name: "Ethereum Mainnet"}
	FantomOpera     Chain = Chain{ID: 250, Name: "Fantom Opera"}
	Linea           Chain = Chain{ID: 59144, Name: "Linea"}
	OasisEmerald    Chain = Chain{ID: 42262, Name: "Oasis Emerald"}
	OP              Chain = Chain{ID: 10, Name: "OP Mainnet"}
	Polygon         Chain = Chain{ID: 137, Name: "Polygon Mainnet"}
	PolygonZKEVM    Chain = Chain{ID: 1101, Name: "Polygon zkEVM"}
	Scroll          Chain = Chain{ID: 534352, Name: "Scroll"}
	Velas           Chain = Chain{ID: 106, Name: "Velas EVM Mainnet"}
	ZKSync          Chain = Chain{ID: 324, Name: "zkSync Mainnet"}
)
