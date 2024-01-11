package token

import (
	"github.com/KyberNetwork/blockchain-toolkit/chain"
	"github.com/ethereum/go-ethereum/common"
)

var WETH9 = map[chain.Chain]*Token{
	chain.ArbitrumOne: &Token{
		ChainID:  chain.ArbitrumOne.ID,
		Address:  common.HexToAddress("0x82af49447d8a07e3bd95bd0d56f35241523fbab1"),
		Decimals: 18,
		Symbol:   "WETH",
		Name:     "Wrapped Ether",
	},
	chain.Aurora: &Token{
		ChainID:  chain.Aurora.ID,
		Address:  common.HexToAddress("0xC9BdeEd33CD01541e1eeD10f90519d2C06Fe3feB"),
		Decimals: 18,
		Symbol:   "WETH",
		Name:     "Wrapped Ether",
	},
	chain.AvalancheCChain: &Token{
		ChainID:  chain.AvalancheCChain.ID,
		Address:  common.HexToAddress("0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7"),
		Decimals: 18,
		Symbol:   "WAVAX",
		Name:     "Wrapped AVAX",
	},
	chain.Base: &Token{
		ChainID:  chain.Base.ID,
		Address:  common.HexToAddress("0x4200000000000000000000000000000000000006"),
		Decimals: 18,
		Symbol:   "WETH",
		Name:     "Wrapped Ether",
	},
	chain.BSC: &Token{
		ChainID:  chain.BSC.ID,
		Address:  common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		Decimals: 18,
		Symbol:   "WBNB",
		Name:     "Wrapped BNB",
	},
	chain.BTTC: &Token{
		ChainID:  chain.BTTC.ID,
		Address:  common.HexToAddress("0x8D193c6efa90BCFf940A98785d1Ce9D093d3DC8A"),
		Decimals: 18,
		Symbol:   "WBTT",
		Name:     "Wrapped BTT",
	},
	chain.Cronos: &Token{
		ChainID:  chain.Cronos.ID,
		Address:  common.HexToAddress("0x5C7F8A570d578ED84E63fdFA7b1eE72dEae1AE23"),
		Decimals: 18,
		Symbol:   "WCRO",
		Name:     "Wrapped CRO",
	},
	chain.Ethereum: &Token{
		ChainID:  chain.Ethereum.ID,
		Address:  common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		Decimals: 18,
		Symbol:   "WETH",
		Name:     "Wrapped Ether",
	},
	chain.FantomOpera: &Token{
		ChainID:  chain.FantomOpera.ID,
		Address:  common.HexToAddress("0x21be370D5312f44cB42ce377BC9b8a0cEF1A4C83"),
		Decimals: 18,
		Symbol:   "WFTM",
		Name:     "Wrapped Fantom",
	},
	chain.Linea: &Token{
		ChainID:  chain.Linea.ID,
		Address:  common.HexToAddress("0xe5d7c2a44ffddf6b295a15c148167daaaf5cf34f"),
		Decimals: 18,
		Symbol:   "WETH",
		Name:     "Wrapped Ether",
	},
	chain.OasisEmerald: &Token{
		ChainID:  chain.OasisEmerald.ID,
		Address:  common.HexToAddress("0x21C718C22D52d0F3a789b752D4c2fD5908a8A733"),
		Decimals: 18,
		Symbol:   "wROSE",
		Name:     "Wrapped ROSE",
	},
	chain.OP: &Token{
		ChainID:  chain.OP.ID,
		Address:  common.HexToAddress("0x4200000000000000000000000000000000000006"),
		Decimals: 18,
		Symbol:   "WETH",
		Name:     "Wrapped Ether",
	},
	chain.Polygon: &Token{
		ChainID:  chain.Polygon.ID,
		Address:  common.HexToAddress("0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270"),
		Decimals: 18,
		Symbol:   "WMATIC",
		Name:     "Wrapped Matic",
	},
	chain.PolygonZKEVM: &Token{
		ChainID:  chain.PolygonZKEVM.ID,
		Address:  common.HexToAddress("0x4f9a0e7fd2bf6067db6994cf12e4495df938e6e9"),
		Decimals: 18,
		Symbol:   "WETH",
		Name:     "Wrapped Ether",
	},
	chain.Scroll: &Token{
		ChainID:  chain.Scroll.ID,
		Address:  common.HexToAddress("0x5300000000000000000000000000000000000004"),
		Decimals: 18,
		Symbol:   "WETH",
		Name:     "Wrapped Ether",
	},
	chain.Velas: &Token{
		ChainID:  chain.Velas.ID,
		Address:  common.HexToAddress("0xc579D1f3CF86749E05CD06f7ADe17856c2CE3126"),
		Decimals: 18,
		Symbol:   "WVLX",
		Name:     "Wrapped VLX",
	},
	chain.ZKSync: &Token{
		ChainID:  chain.ZKSync.ID,
		Address:  common.HexToAddress("0x5aea5775959fbc2557cc8789bc1bf90a239d9a91"),
		Decimals: 18,
		Symbol:   "WETH",
		Name:     "Wrapped Ether",
	},
}
