package app

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
	Use:   "remove [city to remove]",
	Short: "remove city from query list",
	Long: `remove the city from query list
	and no longer to retrieve data from remote site`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("remove city %s from list", strings.Join(args, " "))
	},
}

func init() {

}
