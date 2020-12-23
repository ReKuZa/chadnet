package types

type Transaction struct {
	Outputs []Transaction
	Inputs []Transaction
	Signature string
}

type Block struct {
	Index string
	Hash string
	Transaction Transaction
	Nonce int
	Difficulty int
	PreviousHash string
}
