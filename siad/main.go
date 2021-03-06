package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/NebulousLabs/Sia/build"
)

var (
	// globalConfig is used by the cobra package to fill out the configuration
	// variables.
	globalConfig Config
)

// The Config struct contains all configurable variables for siad. It is
// compatible with gcfg.
type Config struct {
	Siad struct {
		APIaddr  string
		RPCaddr  string
		HostAddr string

		Explorer          bool
		NoBootstrap       bool
		RequiredUserAgent string

		Profile    bool
		ProfileDir string
		SiaDir     string
	}
}

// versionCmd is a cobra command that prints the version of siad.
func versionCmd(*cobra.Command, []string) {
	fmt.Println("Sia Daemon v" + build.Version)
}

// main establishes a set of commands and flags using the cobra package.
func main() {
	root := &cobra.Command{
		Use:   os.Args[0],
		Short: "Sia Daemon v" + build.Version,
		Long:  "Sia Daemon v" + build.Version,
		Run:   startDaemonCmd,
	}

	root.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Long:  "Print version information about the Sia Daemon",
		Run:   versionCmd,
	})

	// Set default values, which have the lowest priority.
	root.PersistentFlags().StringVarP(&globalConfig.Siad.RequiredUserAgent, "agent", "A", "Sia-Agent", "required substring for the user agent")
	root.PersistentFlags().BoolVarP(&globalConfig.Siad.Explorer, "explorer", "E", false, "whether or not to run an explorer in the daemon")
	root.PersistentFlags().StringVarP(&globalConfig.Siad.HostAddr, "host-addr", "H", ":9982", "which port the host listens on")
	root.PersistentFlags().StringVarP(&globalConfig.Siad.ProfileDir, "profile-directory", "P", "profiles", "location of the profiling directory")
	root.PersistentFlags().StringVarP(&globalConfig.Siad.APIaddr, "api-addr", "a", "localhost:9980", "which host:port the API server listens on")
	root.PersistentFlags().StringVarP(&globalConfig.Siad.SiaDir, "sia-directory", "d", "", "location of the sia directory")
	root.PersistentFlags().BoolVarP(&globalConfig.Siad.NoBootstrap, "no-bootstrap", "n", false, "disable bootstrapping on this run")
	root.PersistentFlags().BoolVarP(&globalConfig.Siad.Profile, "profile", "p", false, "enable profiling")
	root.PersistentFlags().StringVarP(&globalConfig.Siad.RPCaddr, "rpc-addr", "r", ":9981", "which port the gateway listens on")

	// Parse cmdline flags, overwriting both the default values and the config
	// file values.
	root.Execute()
}
