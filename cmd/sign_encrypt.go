package cmd

import (
	"github.com/masudur-rahman/jwt-demo/api"
	"github.com/spf13/cobra"
)

var signEncrypt = &cobra.Command{
	Use:   "signNencrypt",
	Short: "sign and encrypt at a time",
	Run: func(cmd *cobra.Command, args []string) {
		api.ExampleSignedAndEncrypted()
		println("\n")
		api.ExampleParseSignedAndEncrypted(data)
	},
}

func init() {
	parseSigned.PersistentFlags().StringVarP(&data, "datum", "D", `eyJhbGciOiJkaXIiLCJjdHkiOiJKV1QiLCJlbmMiOiJBMTI4R0NNIiwidHlwIjoiSldUIn0..TjjlcvAG5c0EMhEN.bLFfFcKD5NwLWSSXK1V1oRnkWGTPDKnu1dnN58SwwsCnVUvWXGQJ6DRKsvZqF32IvZEmy-Jm-bsDzmRem1_JLkGScxguA8qhRfVGbapOM41EQM3nycxQYHIwoVzNPRnhpgzvyL9Ai6yx9bsp-e6cyDCab_xE6IzVWWhGVnIQCmPLMGs6X43jdKf0zd3RyzGpeo7u2eso0FQCAlopgDn3IT8jBuh8Yyf1KcLX9p_vtvOCviZ2OnV_ZUKsN5RrkiAXKoEQ2_0zyFgk4wkH1pd8EaSXzVlBx7clMwSLDzb0Z3PAFGjFdqFejNKnCU_4a4E01dXCouDIKvhEps3H4aXNPH7rzBWJLce-6yzHoqVrQ_jIwjYmZNuUJfoNRu1VuVA-Mbkgi0sHIXkATOLs5fLFIeqMiLGe6p4i27E-NavExfhRYpolKU-zPIDNCyrYAAX4-fyIyW63BW5tffFttrhXU4DjDr9crqbTrknjvBdZBeCEAqd2uOWOJjZMfAZykUu8LCvcbAd-hTr9skWY2y46ab4zLrHBJct1-Yrh1i4zJg2Fvo3abdghE-7fV75gVg.AieRItQ75ReItvtyZ3HHzg`, "Data to parse")
	rootCmd.AddCommand(signEncrypt)
}
