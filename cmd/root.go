// Copyright © 2018 Bren Briggs <code@fraq.io>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//			<3   suser was here   <3
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/bbriggs/bitbot/bitbot"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const VERSION = ""

var (
	cfgFile  string
	server   string
	channels []string
	nick     string
	ssl      bool
	nickserv string
	operUser string
	operPass string
	promAddr string
	prom     bool
)

var pluginMap = map[string]bitbot.NamedTrigger{
	"trackIdleUsers": bitbot.TrackIdleUsers,
	"invite":         bitbot.InviteTrigger,
	"part":           bitbot.PartTrigger,
	"skip":           bitbot.SkipTrigger,
	"info":           bitbot.InfoTrigger,
	"shrug":          bitbot.ShrugTrigger,
	"urlReader":      bitbot.URLReaderTrigger,
	"roll":           bitbot.RollTrigger,
	"decisions":      bitbot.DecisionsTrigger,
	"beef":           bitbot.BeefyTrigger,
	"help":           bitbot.HelpTrigger,
	"8ball":          bitbot.Magic8BallTrigger,
	"tarot":          bitbot.TarotTrigger,
	"markovResponse": bitbot.MarkovResponseTrigger,
	"markovInit":     bitbot.MarkovInitTrigger,
	"troll":          bitbot.TrollLauncherTrigger,
	"markovTrainer":  bitbot.MarkovTrainerTrigger,
	"epeen":          bitbot.EpeenTrigger,
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: VERSION,
	Use:     "bitbot [flags]",
	Short:   "A Golang IRC bot powered by Hellabot",
	Run: func(cmd *cobra.Command, args []string) {
		var plugins []bitbot.NamedTrigger
		for _, plugin := range viper.GetStringSlice("plugins") {
			if p, ok := pluginMap[plugin]; ok {
				plugins = append(plugins, p)
			}
		}
		config := bitbot.Config{
			NickservPass: viper.GetString("nickserv"),
			OperUser:     viper.GetString("operUser"),
			OperPass:     viper.GetString("operPass"),
			Channels:     viper.GetStringSlice("channels"),
			Nick:         viper.GetString("nick"),
			Server:       viper.GetString("server"),
			SSL:          viper.GetBool("ssl"),
			Prometheus:   viper.GetBool("prom"),
			PromAddr:     viper.GetString("promAddr"),
			Admins: bitbot.ACL{
				Permitted: viper.GetStringSlice("admins"),
			},
			Plugins: plugins,
		}
		log.Println("Starting bitbot...")
		bitbot.Run(config)
	},
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&server, "server", "s", server, "target server")
	rootCmd.PersistentFlags().StringVarP(&operUser, "operUser", "", server, "oper username")
	rootCmd.PersistentFlags().StringVarP(&server, "operPass", "", server, "oper password")
	rootCmd.PersistentFlags().StringVarP(&nickserv, "nickserv", "", nickserv, "nickserv password")
	rootCmd.PersistentFlags().StringSliceVarP(&channels, "channels", "c", channels, "channels to join")
	rootCmd.PersistentFlags().StringVarP(&nick, "nick", "n", nick, "nickname")
	rootCmd.PersistentFlags().BoolVarP(&ssl, "ssl", "", ssl, "enable ssl")
	rootCmd.PersistentFlags().BoolVarP(&prom, "prom", "", prom, "enable prometheus")
	rootCmd.PersistentFlags().StringVarP(&promAddr, "promAddr", "", promAddr, "Prometheus metrics address and port")

	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("nickserv", rootCmd.PersistentFlags().Lookup("nickserv"))
	viper.BindPFlag("operUser", rootCmd.PersistentFlags().Lookup("operUser"))
	viper.BindPFlag("operPass", rootCmd.PersistentFlags().Lookup("operPass"))
	viper.BindPFlag("channels", rootCmd.PersistentFlags().Lookup("channels"))
	viper.BindPFlag("nick", rootCmd.PersistentFlags().Lookup("nick"))
	viper.BindPFlag("ssl", rootCmd.PersistentFlags().Lookup("ssl"))
	viper.BindPFlag("promAddr", rootCmd.PersistentFlags().Lookup("promAddr"))

	// All plugins enabled by default
	var defaultPlugins []string
	for plugin, _ := range pluginMap {
		defaultPlugins = append(defaultPlugins, plugin)
	}
	viper.SetDefault("nick", "bitbot")
	viper.SetDefault("prom", false)
	viper.SetDefault("promAddr", "127.0.0.1:8080")
	viper.SetDefault("plugins", defaultPlugins)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
