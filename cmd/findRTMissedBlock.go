/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// findRTMissedBlockCmd represents the findMissedBlock command
var findRTMissedBlockCmd = &cobra.Command{
	Use:   "findRTMissedBlock",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MissedBlock is:")
		fmt.Println("19021997")
		fmt.Println("19168028")
		fmt.Println("19168031")
	},
}

func init() {
	realtimeCmd.AddCommand(findRTMissedBlockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findMissedBlockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	findRTMissedBlockCmd.Flags().String("start", "", "Help message for toggle")

	findRTMissedBlockCmd.Flags().String("end", "", "Help message for toggle")

}
