// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package bc

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/hpb-project/go-hpb/common"
	"github.com/hpb-project/go-hpb/common/hexutil"
	"github.com/hpb-project/go-hpb/common/math"
	"github.com/hpb-project/go-hpb/config"
)

func (g Genesis) MarshalJSON() ([]byte, error) {
	type Genesis struct {
		Config     *config.ChainConfig                         `json:"config"`
		Nonce      math.HexOrDecimal64                         `json:"nonce"`
		Timestamp  math.HexOrDecimal64                         `json:"timestamp"`
		ExtraData  hexutil.Bytes                               `json:"extraData"`
		//VoteIndex  math.HexOrDecimal64                         `json:"voteIndex"`
		
		//CandAddress common.Address                             `json:"candAddress"`
		GasLimit   math.HexOrDecimal64                         `json:"gasLimit"   gencodec:"required"`
		Difficulty *math.HexOrDecimal256                       `json:"difficulty" gencodec:"required"`
		Mixhash    common.Hash                                 `json:"mixHash"`
		Coinbase   common.Address                              `json:"coinbase"`
		
		Alloc      map[common.UnprefixedAddress]GenesisAccount `json:"alloc"      gencodec:"required"`
		Number     math.HexOrDecimal64                         `json:"number"`
		GasUsed    math.HexOrDecimal64                         `json:"gasUsed"`
		ParentHash common.Hash                                 `json:"parentHash"`
		HardwareRandom hexutil.Bytes                           `json:"hardwareRandom"`
	}
	
	
	var enc Genesis
	
	enc.Config= g.Config

	enc.Nonce = math.HexOrDecimal64(g.Nonce)
	enc.Timestamp = math.HexOrDecimal64(g.Timestamp)
	enc.ExtraData = g.ExtraData
   // enc.CandAddress = g.CandAddress
    
	//enc.VoteIndex = math.HexOrDecimal64(g.VoteIndex)
	enc.GasLimit = math.HexOrDecimal64(g.GasLimit)
	enc.Difficulty = (*math.HexOrDecimal256)(g.Difficulty)
	enc.Mixhash = g.Mixhash
	enc.Coinbase = g.Coinbase
	enc.HardwareRandom = g.HardwareRandom
	
	
	if g.Alloc != nil {
		enc.Alloc = make(map[common.UnprefixedAddress]GenesisAccount, len(g.Alloc))
		for k, v := range g.Alloc {
			enc.Alloc[common.UnprefixedAddress(k)] = v
		}
	}
	enc.Number = math.HexOrDecimal64(g.Number)
	enc.GasUsed = math.HexOrDecimal64(g.GasUsed)
	enc.ParentHash = g.ParentHash
	return json.Marshal(&enc)
}

func (g *Genesis) UnmarshalJSON(input []byte) error {
	type Genesis struct {
		Config     *config.ChainConfig                         `json:"config"`
		Nonce      *math.HexOrDecimal64                        `json:"nonce"`
		Timestamp  *math.HexOrDecimal64                        `json:"timestamp"`
		ExtraData  hexutil.Bytes                               `json:"extraData"`
		//VoteIndex  *math.HexOrDecimal64                        `json:"voteIndex"`
		
		//CandAddress *common.Address                            `json:"candAddress"`
		GasLimit   *math.HexOrDecimal64                        `json:"gasLimit"   gencodec:"required"`
		Difficulty *math.HexOrDecimal256                       `json:"difficulty" gencodec:"required"`
		Mixhash    *common.Hash                                `json:"mixHash"`
		Coinbase   *common.Address                             `json:"coinbase"`
		
		Alloc      map[common.UnprefixedAddress]GenesisAccount `json:"alloc"      gencodec:"required"`
		Number     *math.HexOrDecimal64                        `json:"number"`
		GasUsed    *math.HexOrDecimal64                        `json:"gasUsed"`
		ParentHash *common.Hash                                `json:"parentHash"`
		HardwareRandom hexutil.Bytes                           `json:"hardwareRandom"`
	}
	var dec Genesis
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Config != nil {
		g.Config = dec.Config
	}
	if dec.Nonce != nil {
		g.Nonce = uint64(*dec.Nonce)
	}
	if dec.Timestamp != nil {
		g.Timestamp = uint64(*dec.Timestamp)
	}
	//if dec.VoteIndex != nil {
	//	g.VoteIndex = uint64(*dec.VoteIndex)
	//}
	if dec.ExtraData != nil {
		g.ExtraData = dec.ExtraData
	}
	
	if dec.HardwareRandom != nil {
		g.HardwareRandom = dec.HardwareRandom
	}
	
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for Genesis")
	}
	g.GasLimit = uint64(*dec.GasLimit)
	if dec.Difficulty == nil {
		return errors.New("missing required field 'difficulty' for Genesis")
	}
	g.Difficulty = (*big.Int)(dec.Difficulty)
	if dec.Mixhash != nil {
		g.Mixhash = *dec.Mixhash
	}
	if dec.Coinbase != nil {
		g.Coinbase = *dec.Coinbase
	}
	//if dec.CandAddress != nil {
	//	g.CandAddress = *dec.CandAddress
	//}
	if dec.Alloc == nil {
		return errors.New("missing required field 'alloc' for Genesis")
	}
	g.Alloc = make(GenesisAlloc, len(dec.Alloc))
	for k, v := range dec.Alloc {
		g.Alloc[common.Address(k)] = v
	}
	if dec.Number != nil {
		g.Number = uint64(*dec.Number)
	}
	if dec.GasUsed != nil {
		g.GasUsed = uint64(*dec.GasUsed)
	}
	if dec.ParentHash != nil {
		g.ParentHash = *dec.ParentHash
	}
	return nil
}