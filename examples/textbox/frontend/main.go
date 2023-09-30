// Copyright (c) Gabriel de Quadros Ligneul
// SPDX-License-Identifier: MIT (see LICENSE)

package main

import (
	"fmt"
	"log"

	"github.com/gligneul/eggroll"
	"github.com/gligneul/eggroll/examples/textbox"
)

// Redefine the types to make the example cleaner
type (
	Append textbox.Append
	Clear  textbox.Clear
	State  textbox.State
)

func main() {
	client := eggroll.SetupClient[State]()

	indices, err := client.Send(
		&Clear{},
		&Append{Value: "egg"},
		&Append{Value: "roll"},
	)
	if err != nil {
		log.Panic(err)
	}

	lastInput := indices[len(indices)-1]
	if err := client.WaitFor(lastInput); err != nil {
		log.Panic(err)
	}

	state := client.State()
	fmt.Println(state.TextBox) // -> eggroll
}