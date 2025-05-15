package gommon

var (
	NetworkBitcoinExplorer           = "https://blockchair.com/bitcoin/transaction/"
	NetworkEthereumExplorer          = "https://etherscan.io/tx/"
	NetworkBinanceSmartChainExplorer = "https://bscscan.com/tx/"
	NetworkBinanceCoinExplorer       = "https://explorer.binance.org/tx/"
	NetworkTronExplorer              = "https://blockchair.com/bitcoin/transaction/"
	NetworkTonExplorer               = "https://tonviewer.com/transaction/"

	ChainBTC = "bitcoin"
	ChainETH = "ethereum"
	ChainBSC = "smartchain"
	ChainBNB = "binance"
	ChainTRX = "tron"
	ChainTon = "ton"

	Chains = map[string]string{
		ChainBTC: NetworkBitcoinExplorer,
		ChainETH: NetworkEthereumExplorer,
		ChainBSC: NetworkBinanceSmartChainExplorer,
		ChainBNB: NetworkBinanceCoinExplorer,
		ChainTRX: NetworkTronExplorer,
		ChainTon: NetworkTonExplorer,
	}
)

// GeneralConvertor converts a TXID to a clickable link for explorers
func GeneralConvertor(chain, txID string) string {
	return Chains[chain] + txID
}

// ConvertTxidToBSCExplorerLink converts a TXid to a Binance Smart Chain clickable link
func ConvertTxidToBSCExplorerLink(txid string) (link string) {
	return NetworkBinanceSmartChainExplorer + txid
}

// ConvertTxidToBNBExplorerLink converts a TXid to a Binance Coin clickable link
func ConvertTxidToBNBExplorerLink(txid string) (link string) {
	return NetworkBinanceCoinExplorer + txid
}

// ConvertTxidToBTCExplorerLink converts a TXid to a Bitcoin clickable link
func ConvertTxidToBTCExplorerLink(txid string) (link string) {
	return NetworkBitcoinExplorer + txid
}

// ConvertTxidToETHExplorerLink converts a TXid to a Ethereum clickable link
func ConvertTxidToETHExplorerLink(txid string) (link string) {
	return NetworkEthereumExplorer + txid
}

// ConvertTxidToTRXExplorerLink converts a TXid to a TRON clickable link
func ConvertTxidToTRXExplorerLink(txid string) (link string) {
	return NetworkTronExplorer + txid
}

const (
	RegexAddressEvm                 = "^(0x)[0-9A-Fa-f]{40}$"
	RegexAddressBitcoin             = "^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$"
	RegexAddressBitcoinLegacy       = "^1[a-km-zA-HJ-NP-Z1-9]{25,34}$"                 // P2PKH (Legacy)
	RegexAddressBitcoinSegwit       = "^3[a-km-zA-HJ-NP-Z1-9]{25,34}$"                 // P2SH (P2WSH)
	RegexAddressBitcoinSegwitNative = "^bc1[qpzry9x8gf2tvdw0s3jn54khce6mua7l]{39,59}$" // Bech32 (P2WPKH/P2WSH)
	RegexAddressBitcoinCash         = "^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$|^[0-9A-Za-z]{42,42}$"
	RegexAddressCardano             = "^(([0-9A-Za-z]{57,59})|([0-9A-Za-z]{100,104}))$"
	RegexAddressDoge                = "^(D|A|9)[a-km-zA-HJ-NP-Z1-9]{33,34}$"
	RegexAddressLitecoin            = "^(L|M|3)[A-Za-z0-9]{33}$|^(ltc1)[0-9A-Za-z]{39}$"
	RegexAddressRipple              = "^r[1-9A-HJ-NP-Za-km-z]{25,34}$"
	RegexAddressStellar             = "^G[A-D]{1}[A-Z2-7]{54}$"
	RegexAddressTon                 = "^(EQ|UQ)[0-9A-Za-z_-]{46}$"
	RegexAddressTron                = "^T[1-9A-HJ-NP-Za-km-z]{33}$"
)

const (
	/* Tron network*/

	TronContractUsdt = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	TronContractBtt  = "TAFjULxiVgT4qWk6UZwjqwZXTSaGaqnVp4"
	TronContractJst  = "TCFLL5dx5ZJdKnWuesXxi1VPwjLVmWZZy9"
	TronContractWin  = "TLa2f6VPqDgRE67v1736s7bJ8Ray5wYjU7"

	/* Ton network*/

	ToncoinContractUsdt = "EQCxE6mUtQJKFnGfaROTKOt1lZbDiiX1kCixRv7Nw2Id_sDs"
	ToncoinContractNot  = "EQAvlWFDxGF2lXm67y4yzC17wYKD9A0guwPkMs1gOsM__NOT"
	ToncoinContractDogs = "EQCvxJy4eG8hyHBFsZ7eePxrRsUQSFE_jpptRAYBmcG_DOGS"

	/* Binance smart chain */

	BscContractUsdt  = "0x55d398326f99059ff775485246999027b3197955"
	BscContractBtc   = "0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c"
	BscContractEth   = "0x2170ed0880ac9a755fd29b2688956bd959f933f8"
	BscContractShib  = "0x2859e4544c4bb03966803b044a93563bd2d0dd4d"
	BscContractMatic = "0xcc42724c6683b7e57334c4e856f4c9965ed682bd"
	BscContractSol   = "0x570a5d26f7765ecb712c0924e4de545b89fd43df"
	BscContractUsdc  = "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d"
	BscContractDoge  = "0xba2ae424d960c26247dd6c32edc70b295c744c43"
	BscContractXrp   = "0x1d2f0da169ceb9fc7b3144628db156f3f6c60dbe"
	BscContractAda   = "0x3ee2200efb3400fabb9aacf31297cbdd1d435d47"
	BscContractFloki = "0xfb5b838b6cfeedc2873ab27866079ac55363d37e"
	BscContractTon   = "0x76a797a59ba2c17726896976b7b3747bfd1d220f"
	BscContractFtm   = "0xad29abb318791d579433d831ed122afeaf29dcfe"
	BscContractCake  = "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82"

	/* Ethereum */

	EthereumContractUsdt  = "0xdac17f958d2ee523a2206206994597c13d831ec7"
	EthereumContractShib  = "0x95ad61b0a150d79219dcf64e1e6cc01f0b64c4ce"
	EthereumContractLink  = "0x514910771af9ca656af840dff83e8264ecf986ca"
	EthereumContractUsdc  = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
	EthereumContractPepe  = "0x6982508145454ce325ddbe47a25d4ec3d2311933"
	EthereumContractFloki = "0xcf0c122c6b73ff809c693db761e7baebe62b6a2e"
	EthereumContractFtm   = "0x4e15361fd6b4bb609fa63c81a2be19d873717870"
	EthereumContractElon  = "0x761d38e5ddf6ccf6cf7c55759d5210750b5d60f3"
	EthereumContractGala  = "0xd1d2eb1b1e90b638588728b4130137d262c87cae"
	EthereumContractFet   = "0xaea46a60368a7bd060eec7df8cba43b7ef41ad85"
	EthereumContractLadys = "0x12970e6868f88f6557b76120662c1b3e50a646bf"
	EthereumContractPol   = "0x455e53cbb86018ac2b8092fdcd39d8444affc3f6"
	EthereumContractNpxs  = "0xa15c7ebe1f07caf6bff097d8a589fb8ac49ae5b3"
	EthereumContractTusd  = "0x0000000000085d4780b73119b644ae5ecd22b376"
	EthereumContractPax   = "0x8e870d67f660d95d5be530380d0ec0bd388289e1"
	EthereumContractSnt   = "0x744d70fdbe2ba4cf95131626614a1763df805b9e"
	EthereumContractStorm = "0xd0a4b8946cb52f0661273bfbc6fd0e0c75fc6433"
	EthereumContractStorj = "0xb64ef51c888972c908cfacf59b47c1afbc0ab8ac"
	EthereumContractGnt   = "0xa74476443119a942de498590fe1f2454d7d4ac0d"
	EthereumContractBat   = "0x0d8775f648430679a709e98d2b0cb6250d2887ef"
	EthereumContractZrx   = "0xe41d2489571d322189246dafa5ebde1f4699f498"
	EthereumContractDlt   = "0x07e3c70653548b04f0a75970c1f81b4cbbfb606f"
)
