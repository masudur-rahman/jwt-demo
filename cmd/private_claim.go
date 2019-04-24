package cmd

import (
	"github.com/masudur-rahman/jwt-demo/api"
	"github.com/spf13/cobra"
)

var privateClaim = &cobra.Command{
	Use:     "private_claim",
	Short:   "Private claim",
	Example: "jwt-demo private_claim",
	Run: func(cmd *cobra.Command, args []string) {
		api.ExampleSigned_privateClaims()
	},
}

func init() {
	rootCmd.AddCommand(privateClaim)
}
