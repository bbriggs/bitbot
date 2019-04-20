// Copyright © 2018 Bren Briggs <code@fraq.io>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
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
	"github.com/bbriggs/bitbot/bitbot"
	"github.com/spf13/cobra"
	"log"
)

var server string
var channels []string
var nick string
var ssl bool = false
var nickservPass string
var operUser string
var operPass string
var admins []string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run bitbot",
	Run: func(cmd *cobra.Command, args []string) {
		c := bitbot.Config{
			NickservPass: nickservPass,
			OperUser:     operUser,
			OperPass:     operPass,
			Channels:     channels,
			Nick:         nick,
			Server:       server,
			SSL:          ssl,
			Admins: bitbot.ACL{
				Permitted: admins,
			},
		}
		log.Println("Starting bitbot...")
		bitbot.Run(c)
	},
}

func init() {
	// runCmd represents the run command

	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&server, "server", "s", server, "target server")
	runCmd.Flags().StringVarP(&nickservPass, "nickserv", "", nickservPass, "nickserv password")
	runCmd.Flags().StringVarP(&operUser, "operUser", "", operUser, "oper username")
	runCmd.Flags().StringVarP(&operPass, "operPass", "", operPass, "oper password")
	runCmd.Flags().StringSliceVarP(&channels, "channels", "c", channels, "channels to join")
	runCmd.Flags().StringSliceVarP(&admins, "admins", "", admins, "Hostmasks of administrators")
	runCmd.Flags().StringVarP(&nick, "nick", "n", nick, "nickname")
	runCmd.Flags().BoolVarP(&ssl, "ssl", "", ssl, "enable ssl")

}
