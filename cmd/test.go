package cmd

import (
	"github.com/masudur-rahman/jwt-demo/api"
	"github.com/spf13/cobra"
)

var example = &cobra.Command{
	Use:   "example",
	Short: "Parse JWT Signed data",
	//Example: "jwt-demo private_claim",
	Run: func(cmd *cobra.Command, args []string) {
		api.ExampleClaims_Validate_withParse()
	},
}

func init() {
	//parseSigned.PersistentFlags().StringVarP(&data, "data", "d", `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiTm9uZSIsIkV2ZXJ5b25lIl0sImN1c3RvbSI6IlRoZXJlIGlzIG5vIGN1c3RvbSB2YWx1ZSBhdmFpbGFibGUiLCJleHAiOjE1NTY0NTMzODAsImlzcyI6Ik1hc3VkdXIgUmFobWFuIiwibmJmIjoxNDUxNjA2NDAwLCJzdWIiOiJUZXN0aW5nIn0.DBQoFkbHshjCig_534GVZ3M450ZclvbDW-RUm8shUOg`, "Data to parse")
	rootCmd.AddCommand(example)
}
