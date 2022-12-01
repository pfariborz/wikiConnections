/*
Copyright © 2022 PARI GARAY pjgaray617@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wikiConnections",
	Short: "wikiConnections is a program that will route a path between two subjects in Wikipedia",
	Long: `wikiConnections is a program that will take a starting subject page in Wikipedia and
	plot the path to an ending Wikipedia page. For example: 
	wikiConnections --start Ballet --goal Tennis 
	
	Returned back will be the number of hops it takes to get to Tennis`,

	Run: func(cmd *cobra.Command, args []string) {
		start, _ := cmd.Flags().GetString("start")
		goal, _ := cmd.Flags().GetString("goal")

		wiki := newWikiLinkSearch()
		graph := NewGraph(3, &wiki)
		graph.depthFirstSearch(start, goal)

		fmt.Println("Number of hops is: ", len(graph.mapPath))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Flags for wikiConnetions command
	rootCmd.Flags().StringP("start", "s", "", "Starting subject in WikiPedia")
	rootCmd.Flags().StringP("goal", "g", "", "Final Wikipedia page")
}
