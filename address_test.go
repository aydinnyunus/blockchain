package blockchain

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestGetAddress(t *testing.T) {
	fmt.Println("===== TESTING ADDRESS =====")

	c, e := New()
	resp, e := c.GetAddress("162FjqU7RYdojnejCDe6zrPDUpaLcv9Hhq")
	if e != nil {
		fmt.Print(e)
	}
	fmt.Println(resp.Address)
	fmt.Println(resp.Hash160)
	fmt.Println(resp.Address)
	fmt.Println(resp.NTx)
	fmt.Println(resp.TotalReceived)
	fmt.Println(resp.TotalSent)
	fmt.Println(resp.FinalBalance)

	for i := range resp.Txs {
		fmt.Println(resp.Txs[i].Result)

		for j := range resp.Txs[i].Inputs {
			fmt.Println(resp.Txs[i].Inputs[j].Sequence)
			fmt.Println(resp.Txs[i].Inputs[j].PrevOut.Spent)
		}
	}
}

func TestGetAddresses(t *testing.T) {
	fmt.Println("===== TESTING ADDRESSES =====")

	c, e := New()
	if e != nil {
		fmt.Print(e)
	}

	addresses := []string{
		"162FjqU7RYdojnejCDe6zrPDUpaLcv9Hhq",
		"1K3Vs8tPu2YkAoWmrkjUQVJuxr7wgPP3Wf"}

	resp, _ := c.GetAddresses(addresses)
	for i := range resp.Addresses {
		fmt.Println(resp.Addresses[i].Address)

		for j := range resp.Txs[i].Inputs {
			fmt.Println(resp.Txs[i].Inputs[j].Sequence)
			fmt.Println(resp.Txs[i].Inputs[j].PrevOut.Spent)
		}
	}
}

func TestGetETHAddress(t *testing.T) {
	fmt.Println("===== TESTING ADDRESS =====")

	c, e := New()
	resp, e := c.GetETHAddress("0x7444bce361ead96a737b351899d30f4df1e16900")
	if e != nil {
		fmt.Print(e)
	}

	for i, _ := range resp.Transactions{
		fmt.Println(resp.Transactions[i].Hash)
	}

}

func TestCheckAddress(t *testing.T) {
	//fmt.Println("===== TESTING CHECK ADDRESS =====")

	c, e := New()
	resp, e := c.CheckAddress("0xef4fad847e4078828e32327536f2c1b8a5f53630")
	if e != nil {
		fmt.Print(e)
	}

	exc := strings.Fields(resp)
	if "Uniswap" == exc[0]{
		fmt.Println("Successfull")
	}
}


func BenchmarkClient_CheckAddress(b *testing.B) {
	c, _ := New()

	for i := 0; i < b.N; i++ {
		resp, e := c.CheckAddress("0xef4fad847e4078828e32327536f2c1b8a5f53630")
		if e != nil {
			time.Sleep(5000)
			fmt.Print(e)
		}

		exc := strings.Fields(resp)
		if "Uniswap" == exc[0]{
			//fmt.Println("Successfull")
		}
	}
}
