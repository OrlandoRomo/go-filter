package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var filters map[int]string

func init() {
	filters = map[int]string{
		1: "Grey filter",
		2: "Sepia filter",
		3: "Negative filter",
		4: "Flip or mirror filter",
		5: "Sketch filter",
		6: "Red filter",
		7: "Blue filter",
		8: "Green filter",
	}
}
func NewListCommand() *cli.Command {
	return &cli.Command{
		Name:    "filters",
		Aliases: []string{"fltrs"},
		Usage:   "filters lists of all available filters or tones.",
		Action:  listFilters,
	}
}

func listFilters(c *cli.Context) error {
	fmt.Println("List of supported filters:")
	for _, filter := range filters {
		fmt.Printf("\t%s %s\n", "-", filter)
	}
	return nil
}
