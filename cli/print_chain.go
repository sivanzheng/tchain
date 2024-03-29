package cli

import (
	"fmt"
	"strconv"
	"tchain/blockchain"
)

func (cli *CLI) printChain(nodeID string) {
	bc := blockchain.NewBlockchain(nodeID)
	defer bc.DB.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("============ Block %x ============\n", block.Hash)
		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("Prev block: %x\n", block.PrevBlockHash)
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Printf("TX ID: %x\n\n", tx.ID)
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
