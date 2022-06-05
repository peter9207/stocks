package cmd

import (
	"fmt"

	"github.com/peter9207/stocks/indicators"
	"github.com/peter9207/stocks/readers"
	"github.com/spf13/cobra"
)

// relevanceCmd represents the relevance command
var relevanceCmd = &cobra.Command{
	Use:   "relevance <filename>",
	Short: "highlight the relevant dates of a stock data file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			return
		}

		filename := args[0]

		reader := readers.NewCSV()

		data, err := reader.Read(filename, true, 0, 3)
		if err != nil {
			fmt.Println(err)
			return
		}

		result := indicators.Relevance(data, 3)
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(relevanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// relevanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// relevanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
