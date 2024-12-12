package gommon

var (
	NetworkBitcoinExplorer           = "https://blockchair.com/bitcoin/transaction/"
	NetworkEthereumExplorer          = "https://etherscan.io/tx/"
	NetworkBinanceSmartChainExplorer = "https://bscscan.com/tx/"
	NetworkBinanceCoinExplorer       = "https://explorer.binance.org/tx/"
	NetworkTronExplorer              = "https://blockchair.com/bitcoin/transaction/"

	ChainBTC = "bitcoin"
	ChainETH = "ethereum"
	ChainBSC = "smartchain"
	ChainBNB = "binance"
	ChainTRX = "tron"

	Chains = map[string]string{
		ChainBTC: NetworkBitcoinExplorer,
		ChainETH: NetworkEthereumExplorer,
		ChainBSC: NetworkBinanceSmartChainExplorer,
		ChainBNB: NetworkBinanceCoinExplorer,
		ChainTRX: NetworkTronExplorer,
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
