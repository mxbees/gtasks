package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gtasks",
	Short:   "A CLI Tool for Google Tasks",
	Version: "0.9.0",
	Long: `A CLI Tool for managing your Google Tasks:

Made with ♥	by https://github.com/BRO3886


████████╗░█████╗░░██████╗██╗░░██╗░██████╗  ░█████╗░██╗░░░░░██╗
╚══██╔══╝██╔══██╗██╔════╝██║░██╔╝██╔════╝  ██╔══██╗██║░░░░░██║
░░░██║░░░███████║╚█████╗░█████═╝░╚█████╗░  ██║░░╚═╝██║░░░░░██║
░░░██║░░░██╔══██║░╚═══██╗██╔═██╗░░╚═══██╗  ██║░░██╗██║░░░░░██║
░░░██║░░░██║░░██║██████╔╝██║░╚██╗██████╔╝  ╚█████╔╝███████╗██║
░░░╚═╝░░░╚═╝░░╚═╝╚═════╝░╚═╝░░╚═╝╚═════╝░  ░╚════╝░╚══════╝╚═╝`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	viper.SetDefault("license", "apache")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".google-tasks-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".google-tasks-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
