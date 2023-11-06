// Copyright (c) Gabriel de Quadros Ligneul
// SPDX-License-Identifier: MIT (see LICENSE)

package cmd

import (
	"os"

	"github.com/gligneul/eggroll/internal/codegen"
	"github.com/spf13/cobra"
)

var schemaGenArgs struct {
	packageName string
	outputPath  string
}

var schemaGenCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate ABI bindings",
	Long:  `Generate the Go bindings for the given ABI yaml file.`,
	Run: func(cmd *cobra.Command, args []string) {
		abiJson := schemaLoad()
		code, err := codegen.Gen(abiJson, schemaGenArgs.packageName)
		cobra.CheckErr(err)

		outputFile, err := os.Create(schemaGenArgs.outputPath)
		cobra.CheckErr(err)
		defer outputFile.Close()

		_, err = outputFile.Write([]byte(code))
		cobra.CheckErr(err)
	},
}

func init() {
	schemaCmd.AddCommand(schemaGenCmd)

	schemaGenCmd.Flags().StringVar(
		&schemaGenArgs.packageName, "package", "main", "Name of the generated package")

	schemaGenCmd.Flags().StringVar(
		&schemaGenArgs.outputPath, "output", "schema.go", "Target file for the Go binding")
}
