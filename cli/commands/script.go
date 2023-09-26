package commands

import (
	"log"

	"github.com/GhostNet-Dev/GhostNet-Core/pkg/gcrypto"
	"github.com/spf13/cobra"
)

var (
	executeScript = false
	codeFilepath  = "./code.gs"
	scriptType    = 0
)

func ScriptCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "script",
		Short: `Register script in blockchain`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			return initializeConfig(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Working with OutOrStdout/OutOrStderr allows us to unit test our command easier
			if registerScriptCommand(username, password) {
				log.Println("Regist Complete")
			}
		},
	}
	cmd.Flags().StringVarP(&host, "ip", "i", "", "Host Address")
	cmd.Flags().StringVarP(&port, "port", "", "50129", "Port Number")
	cmd.Flags().StringVarP(&rpcPort, "rpc", "r", "50229", "GRPC Port Number")
	cmd.Flags().StringVarP(&username, "username", "u", "", "Ghost Account Nickname")
	cmd.Flags().StringVarP(&password, "password", "p", "", "Ghost Account Password")
	cmd.Flags().StringVarP(&codeFilepath, "code", "c", "", "script file path")
	cmd.Flags().Uint32VarP(&id, "id", "", 0, "Container Id, if not select, show all container")
	cmd.Flags().Uint32VarP(&timeout, "timeout", "t", 3, "rpc connection timeout")
	cmd.Flags().IntVarP(&scriptType, "script_type", "", 0, "script type")
	cmd.Flags().BoolVarP(&executeScript, "exe", "e", false, "execute script")

	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
	return cmd
}

func registerScriptCommand(username, password string) bool {
	cfg.Username = username
	// for encrypt passwd
	cfg.Password = gcrypto.PasswordToSha256(password)
	cfg.Ip = host
	cfg.Port = port
	return true
}
