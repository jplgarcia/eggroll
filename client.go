// Copyright (c) Gabriel de Quadros Ligneul
// SPDX-License-Identifier: MIT (see LICENSE)

package eggroll

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gligneul/eggroll/blockchain"
	"github.com/gligneul/eggroll/internal/sunodo"
	"github.com/gligneul/eggroll/reader"
)

// Result of an advance input.
type AdvanceResult struct {
	*reader.Input

	// Result returned by the contract advance method.
	Result []byte

	// Logs generated during the advance method.
	Logs []string
}

func newAdvanceResult(input *reader.Input) *AdvanceResult {
	var result AdvanceResult
	result.Input = input
	for _, report := range input.Reports {
		tag, payload, err := decodeReport(report.Payload)
		if err != nil {
			// TODO how do we report this?
			continue
		}
		switch tag {
		case reportTagResult:
			result.Result = payload
		case reportTagLog:
			result.Logs = append(result.Logs, string(payload))
		}
	}
	return &result
}

// Read the rollups state off chain.
// For more details, see the eggroll/reader package.
type readerAPI interface {
	Input(ctx context.Context, index int) (*reader.Input, error)
}

// Communicate with the blockchain.
type blockchainAPI interface {
	SendInput(ctx context.Context, dappAddress common.Address, input []byte) error
}

// Configuration for the Client.
type ClientConfig struct {
	DAppAddress      common.Address
	GraphqlEndpoint  string
	ProviderEndpoint string
}

// The Client interacts with the DApp contract off chain.
type Client struct {
	ClientConfig
	reader     readerAPI
	blockchain blockchainAPI
}

// Create the Client with a custom config.
func NewClient(config ClientConfig) *Client {
	return &Client{
		ClientConfig: config,
		reader:       reader.NewGraphQLReader(config.GraphqlEndpoint),
		blockchain:   blockchain.NewETHClient(config.ProviderEndpoint),
	}
}

// Create the Client loading the config from environment variables.
// This function uses the context when building the client but do not store it.
func NewLocalClient() *Client {
	dappAddress, err := sunodo.GetDAppAddress()
	if err != nil {
		panic(fmt.Sprintf("failed to get DApp address: %v", err))
	}
	config := ClientConfig{
		DAppAddress:      dappAddress,
		GraphqlEndpoint:  "http://localhost:8080/graphql",
		ProviderEndpoint: "http://localhost:8545",
	}
	return NewClient(config)
}

//
// Send functions
//

// Send the input as bytes to the DApp contract.
func (c *Client) SendInputBytes(ctx context.Context, inputBytes []byte) error {
	return c.blockchain.SendInput(ctx, c.DAppAddress, inputBytes)
}

// Send a generic input to the DApp contract.
func (c *Client) SendInputJson(ctx context.Context, input any) error {
	inputBytes, err := EncodeJSONInput(input)
	if err != nil {
		return err
	}
	return c.SendInputBytes(ctx, inputBytes)
}

//
// Reader functions
//

// Wait until the DApp contract processes a given input.
// Returns the advance result of that input.
func (c *Client) WaitFor(ctx context.Context, inputIndex int) (*AdvanceResult, error) {
	for {
		input, err := c.reader.Input(ctx, inputIndex)
		if err != nil {
			if _, ok := err.(reader.NotFound); ok {
				goto wait
			}
			return nil, fmt.Errorf("faild to read input: %v", err)
		}
		if input.Status != reader.CompletionStatusUnprocessed {
			return newAdvanceResult(input), nil
		}
	wait:
		time.Sleep(time.Second)
	}
}

// Sync to the latest Dapp state.
// Return the updated slice of Advance results.
func (c *Client) Sync(ctx context.Context, results []*AdvanceResult) ([]*AdvanceResult, error) {
	inputIndex := 0
	if len(results) != 0 {
		inputIndex = results[len(results)-1].Index
	}
	for {
		input, err := c.reader.Input(ctx, inputIndex)
		if err != nil {
			if _, ok := err.(reader.NotFound); ok {
				break
			}
			return nil, fmt.Errorf("failed to read input: %v", err)
		}
		if input.Status == reader.CompletionStatusUnprocessed {
			break
		}
		results = append(results, newAdvanceResult(input))
		inputIndex++
	}
	return results, nil
}
