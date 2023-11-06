// Code generated by EggRoll - DO NOT EDIT.

package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gligneul/eggroll/pkg/eggtypes"
)

var (
	_ = big.NewInt
	_ = common.Big1
	_ = eggtypes.MustAddSchema
)

//
// Types
//

// Schema with selector 09df59d0
type TextBox struct {
	Value string
}

// Schema with selector 86cdbae9
type Append struct {
	Value string
}

// Schema with selector 52efea6e
type Clear struct {
}

//
// IDs
//

// textBox ID
var TextBoxID = [4]byte{0x9, 0xdf, 0x59, 0xd0}

// append ID
var AppendID = [4]byte{0x86, 0xcd, 0xba, 0xe9}

// clear ID
var ClearID = [4]byte{0x52, 0xef, 0xea, 0x6e}

//
// Encode
//

// Encode textBox into binary data.
func EncodeTextBox(
	Value string,
) []byte {
	values := make([]any, 1)
	values[0] = Value
	data, err := _abi.Methods["textBox"].Inputs.PackValues(values)
	if err != nil {
		panic(fmt.Sprintf("failed to encode textBox: %v", err))
	}
	return append(TextBoxID[:], data...)
}

// Encode textBox into binary data.
func (v TextBox) Encode() []byte {
	return EncodeTextBox(
		v.Value,
	)
}

// Encode append into binary data.
func EncodeAppend(
	Value string,
) []byte {
	values := make([]any, 1)
	values[0] = Value
	data, err := _abi.Methods["append"].Inputs.PackValues(values)
	if err != nil {
		panic(fmt.Sprintf("failed to encode append: %v", err))
	}
	return append(AppendID[:], data...)
}

// Encode append into binary data.
func (v Append) Encode() []byte {
	return EncodeAppend(
		v.Value,
	)
}

// Encode clear into binary data.
func EncodeClear() []byte {
	values := make([]any, 0)
	data, err := _abi.Methods["clear"].Inputs.PackValues(values)
	if err != nil {
		panic(fmt.Sprintf("failed to encode clear: %v", err))
	}
	return append(ClearID[:], data...)
}

// Encode clear into binary data.
func (v Clear) Encode() []byte {
	return EncodeClear()
}

//
// Decode
//

func _decode_TextBox(values []any) (any, error) {
	if len(values) != 1 {
		return nil, fmt.Errorf("wrong number of values")
	}
	var ok bool
	var v TextBox
	v.Value, ok = values[0].(string)
	if !ok {
		return nil, fmt.Errorf("failed to decode textBox.value")
	}
	return v, nil
}

func _decode_Append(values []any) (any, error) {
	if len(values) != 1 {
		return nil, fmt.Errorf("wrong number of values")
	}
	var ok bool
	var v Append
	v.Value, ok = values[0].(string)
	if !ok {
		return nil, fmt.Errorf("failed to decode append.value")
	}
	return v, nil
}

func _decode_Clear(values []any) (any, error) {
	if len(values) != 0 {
		return nil, fmt.Errorf("wrong number of values")
	}
	var v Clear
	return v, nil
}

//
// Init
//

const _JSON_ABI = `[
  {
    "name": "append",
    "type": "function",
    "stateMutability": "nonpayable",
    "inputs": [
      {
        "name": "value",
        "type": "string",
        "internalType": "string",
        "components": null
      }
    ],
    "outputs": null
  },
  {
    "name": "clear",
    "type": "function",
    "stateMutability": "nonpayable",
    "inputs": null,
    "outputs": null
  },
  {
    "name": "textBox",
    "type": "function",
    "stateMutability": "nonpayable",
    "inputs": [
      {
        "name": "value",
        "type": "string",
        "internalType": "string",
        "components": null
      }
    ],
    "outputs": null
  }
]
`

var _abi abi.ABI

func init() {
	var err error
	_abi, err = abi.JSON(strings.NewReader(_JSON_ABI))
	if err != nil {
		// This should not happen
		panic(fmt.Sprintf("failed to decode ABI: %v", err))
	}
	eggtypes.MustAddSchema(eggtypes.MessageSchema{
		ID:        TextBoxID,
		Kind:      "textBox",
		Arguments: _abi.Methods["textBox"].Inputs,
		Decoder:   _decode_TextBox,
	})
	eggtypes.MustAddSchema(eggtypes.MessageSchema{
		ID:        AppendID,
		Kind:      "append",
		Arguments: _abi.Methods["append"].Inputs,
		Decoder:   _decode_Append,
	})
	eggtypes.MustAddSchema(eggtypes.MessageSchema{
		ID:        ClearID,
		Kind:      "clear",
		Arguments: _abi.Methods["clear"].Inputs,
		Decoder:   _decode_Clear,
	})
}