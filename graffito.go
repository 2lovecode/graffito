package main

import (
	"graffito/algorithm"
	"graffito/experiment"
	"graffito/leetcode"
	"graffito/pattern"
	"graffito/practice"
	"graffito/tools"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{Use: "graffito"}

	rootCmd.AddCommand(tools.NewToolsCommand())
	rootCmd.AddCommand(algorithm.NewAlgorithmCommand())
	rootCmd.AddCommand(experiment.NewExperimentCommand())
	rootCmd.AddCommand(practice.NewPracticeCommand())
	rootCmd.AddCommand(pattern.NewPatternCommand())
	rootCmd.AddCommand(leetcode.NewLeetcodeCommand())

	rootCmd.Execute()
}
