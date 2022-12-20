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
	plot the path to an ending Wikipedia page. User can select either DFS or BFS as the
	algorithm to perform the search (default is BFS). For example: 
	wikiConnections --start Ballet --goal Tennis --algorithm BFS
	
	Returned back will be the number of hops it takes to get to Tennis`,

	Run: func(cmd *cobra.Command, args []string) {
		start, _ := cmd.Flags().GetString("start")
		goal, _ := cmd.Flags().GetString("goal")
		algorithm, _ := cmd.Flags().GetString("algorithm")
		pageCount, _ := cmd.Flags().GetInt("pageLimit")

		wiki := newWikiLinkSearch()
		graph := NewGraph(pageCount, &wiki)
		var goalReached bool

		if algorithm == "DFS" {
			_, goalReached = graph.depthFirstSearch(start, goal)
		} else {
			_, goalReached = graph.breathFirstSearch(start, goal)
		}
		if goalReached {
			fmt.Printf("wikiConnections found a path from start: %s to goal: %s\n", start, goal)
			count := graph.printPath(start, goal)
			fmt.Println("Number of hops is: ", count)
		} else {
			fmt.Printf("wikiConnections did not find a path from start: %s to goal: %s", start, goal)
		}
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
	rootCmd.Flags().IntP("pageLimit", "p", maxPagesVisited, "Maximum number of pages visited")
	rootCmd.Flags().StringP("algorithm", "a", "", "Search algorithm either BFS or DFS")
}
