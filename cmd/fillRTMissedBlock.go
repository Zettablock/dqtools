/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var blocks string
var file string
var autoyes bool

// fillRTMissedBlockCmd represents the fillMissedBlock command
var fillRTMissedBlockCmd = &cobra.Command{
	Use:   "fillRTMissedBlock",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("are you sure to apply fillMissedBlock to these blocks?")
		fmt.Println("19021997")
		fmt.Println("19168028")
		fmt.Println("19168031")
		fmt.Println("please type yes or no:")
		var yes string
		_, err := fmt.Scanln(&yes)
		if err != nil {
			fmt.Println("canceled")
			return
		}
		if yes == "yes" || yes == "y" {
			fmt.Println("19021997:OK")
			fmt.Println("19168028:OK")
			fmt.Println("19168031:OK")
		} else {
			fmt.Println("canceled")
		}

	},
}

func init() {
	realtimeCmd.AddCommand(fillRTMissedBlockCmd)

	fillRTMissedBlockCmd.Flags().StringVarP(&blocks, "blocks", "b", "", "Help message for toggle")
	fillRTMissedBlockCmd.Flags().StringVarP(&file, "file", "f", "", "Help message for toggle")
	fillRTMissedBlockCmd.MarkFlagsOneRequired("blocks", "file")
	fillRTMissedBlockCmd.MarkFlagsMutuallyExclusive("blocks", "file")
	fillRTMissedBlockCmd.Flags().BoolVarP(&autoyes, "autoyes", "a", false, "auto yes without confirm")

}
