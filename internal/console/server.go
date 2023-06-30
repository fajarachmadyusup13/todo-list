package console

import (
	"github.com/fajarachmadyusup13/todo-list/internal/db"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  "This subcommand start the server",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	db.InitializeCockroachConn()

}
