/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"os/user"
	"strings"

	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Pallinder/go-randomdata"
	log "github.com/taylormonacelli/reactnut/cmd/logging"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "reactnut",
	Args: cobra.ExactArgs(1),

	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		basePath := args[0]
		dostuff(basePath)
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	rootCmd.PersistentFlags().String("log-level",
		"info", "Log level (trace, debug, info, warn, error, fatal, panic)",
	)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"config file (default is $HOME/.reactnut.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".reactnut" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".reactnut")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	viper.BindPFlag("log-level", rootCmd.Flags().Lookup("log-level"))
	logLevel := log.ParseLogLevel(viper.GetString("log-level"))
	log.Logger.SetLevel(logLevel)

	viper.BindPFlag("config", rootCmd.Flags().Lookup("config"))

	viper.AutomaticEnv() // read in environment variables that match
	viper.ReadInConfig()
}

func dostuff(basePath string) {
	adjective := randomdata.Adjective()
	noun := randomdata.Noun()
	concat := fmt.Sprintf("%s%s", adjective, noun)
	fullPath := filepath.Join(basePath, concat)

	for i := 0; i < 10; i++ {
		if !pathExists(fullPath) {
			break
		}
		adjective := randomdata.Adjective()
		noun := randomdata.Noun()
		concat := fmt.Sprintf("%s%s", adjective, noun)
		fullPath = filepath.Join(basePath, concat)
	}
	log.Logger.Traceln(fullPath)
}

func pathExists(path string) bool {
	path, err := expandTilde(path)
	if err != nil {
		panic(err)
	}
	log.Logger.Trace("")
	log.Logger.Traceln(path) // output: /Users/username/Documents/example.txt

	// Use os.Stat() to get information about the path
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	// Check if the error value is nil, which indicates that the path exists
	if err == nil {
		// Check if the path is a directory
		if fileInfo.Mode().IsDir() {
			log.Logger.Tracef("%s is a directory\n", path)
		} else {
			log.Logger.Tracef("%s is a file\n", path)
		}
	} else {
		log.Logger.Tracef("Path %s does not exist\n", path)
	}
	return true
}

func expandTilde(path string) (string, error) {
	if strings.HasPrefix(path, "~/") || path == "~" {
		currentUser, err := user.Current()
		if err != nil {
			return "", err
		}
		return strings.Replace(path, "~", currentUser.HomeDir, 1), nil
	}
	return path, nil
}
