package blockchain

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

type Address struct {
	Hash160       string `json:"hash160"`
	Address       string `json:"address"`
	NTx           int    `json:"n_tx"`
	TotalReceived int    `json:"total_received"`
	TotalSent     int    `json:"total_sent"`
	FinalBalance  int    `json:"final_balance"`
	Txs           []*Tx  `json:"txs"`
}

type MultiAddr struct {
	Addresses []*Address `json:"addresses"`
	Txs       []*Tx      `json:"txs"`
}

type Tx struct {
	Result      int       `json:"result"`
	Ver         int       `json:"ver"`
	Size        int       `json:"size"`
	Inputs      []*Inputs `json:"inputs"`
	Time        int       `json:"time"`
	BlockHeight int       `json:"block_height"`
	TxIndex     int       `json:"tx_index"`
	VinSz       int       `json:"vin_sz"`
	Hash        string    `json:"hash"`
	VoutSz      int       `json:"vout_sz"`
	RelayedBy   string    `json:"relayed_by"`
	Out         []*Out    `json:"out"`
}

type Inputs struct {
	Sequence int      `json:"sequence"`
	Script   string   `json:"script"`
	PrevOut  *PrevOut `json:"prev_out"`
}

type PrevOut struct {
	Spent   bool   `json:"spent"`
	TxIndex int    `json:"tx_index"`
	Type    int    `json:"type"`
	Addr    string `json:"addr"`
	Value   int    `json:"value"`
	N       int    `json:"n"`
	Script  string `json:"script"`
}

type Out struct {
	Spent   bool   `json:"spent"`
	TxIndex int    `json:"tx_index"`
	Type    int    `json:"type"`
	Addr    string `json:"addr"`
	Value   int    `json:"value"`
	N       int    `json:"n"`
	Script  string `json:"script"`
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}

func GetHtmlTitle(r io.Reader) (string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		panic("Fail to parse html")
	}

	return traverse(doc)
}

func (c *Client) GetAddress(address string) (*Address, error) {
	rsp := &Address{}
	var path = "/address/" + address
	e := c.loadResponse(path, rsp, true)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetAddresses(addresses []string) (*MultiAddr, error) {
	rsp := &MultiAddr{}
	var path = "/multiaddr?active=" + strings.Join(addresses, "|")
	e := c.loadResponse(path, rsp, false)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) CheckAddress(address string) (string, error) {
	resp, err := http.Get("https://etherscan.io/address/" + address)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if t, ok := GetHtmlTitle(resp.Body); ok {
		return t, err
	} else {
		return "", err
	}
}
