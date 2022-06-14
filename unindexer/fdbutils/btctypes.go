package fdbutils

type Block struct {
	Hash              string   `json:"hash"`
	Confirmations     int64    `json:"confirmations"`
	Size              int64    `json:"size"`
	StrippedSize      int64    `json:"strippedSize"`
	Weight            int64    `json:"weight"`
	Height            int64    `json:"height"`
	Version           int64    `json:"version"`
	VersionHex        string   `json:"versionHex"`
	MerkleRoot        string   `json:"merkleRoot"`
	Time              int64    `json:"time"`
	MedianTime        int64    `json:"medianTime"`
	Nonce             int64    `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	ChainWork         string   `json:"chainWork"`
	NTx               int64    `json:"nTx"`
	PreviousBlockHash string   `json:"previousblockhash"`
	NextBlockHash     string   `json:"nextblockhash"`
	Tx                []string `json:"tx"`
}

type Transaction struct {
	InActiveChain bool              `json:"in_active_chain"`
	Hex           string            `json:"hex"`
	TxId          string            `json:"txid"`
	Hash          string            `json:"hash"`
	Size          int64             `json:"size"`
	VSize         int64             `json:"vsize"`
	Weight        int64             `json:"weight"`
	Version       int64             `json:"version"`
	LockTime      int64             `json:"locktime"`
	BlockHash     string            `json:"blockhash"`
	Confirmations int64             `json:"confirmations"`
	BlockTime     int64             `json:"blocktime"`
	Time          int64             `json:"time"`
	Vin           []TransactionVIn  `json:"vin"`
	Vout          []TransactionVOut `json:"vout"`
}

type TransactionVIn struct {
	TxId        string      `json:"txid"`
	Coinbase    string      `json:"coinbase,omitempty"`
	Vout        int64       `json:"vout"`
	ScriptSig   TxScriptSig `json:"scriptSig"`
	Sequence    int64       `json:"sequence"`
	TxInWitness []string    `json:"txinwitness"`
}

type TxScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

type TransactionVOut struct {
	Value        float64                 `json:"value"`
	N            int64                   `json:"n"`
	ScriptPubKey TransactionScriptPubKey `json:"scriptPubKey"`
}

type TransactionScriptPubKey struct {
	TxScriptSig
	ReqSigs   int64    `json:"reqSigs"`
	Type      string   `json:"type"`
	Addresses []string `json:"addresses"`
	Desc      string   `json:"desc,omitempty"`
	Address   string   `json:"address,omitempty"`
}
