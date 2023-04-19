package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVarP(&crawlerName, "crawler", "b", "", "The crawler name.")
	viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
}

var rootCmd = &cobra.Command{
	Use:   "newz",
	Short: "newz is a crawler",
	Long:  `newz is a crawler that crawls the web and stores the data in a database.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var crawlerName string
var configFile string

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
