package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var compareFacesCmd = &cobra.Command{
	Use:   "compare-faces",
	Short: "See if the same person is in both images",
	Long:  `Takes a source and target image and copares the largest face from each image to see if they are the same person`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Compare faces command")
	},
}

func init() {
	RootCmd.AddCommand(compareFacesCmd)
}
