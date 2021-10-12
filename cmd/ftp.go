/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ftpCmd represents the ftp command
var ftpCmd = &cobra.Command{
	Use:   "ftp",
	Short: "to set/update/test ftp connection",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ftp called")
		if len(args) > 0 {
			fmt.Println("argument len: ", len(args))
			fmt.Println("argument content: ", args)
		}
		fmt.Println("Hostname: ", FtpHostname)
		fmt.Println("Username: ", FtpUsername)
		fmt.Println("UserPass: ", FtpUserPass)
		fmt.Println("Port: ", FtpPort)
		fmt.Println("Directory: ", FtpDirectory)
	},
}

var FtpHostname string
var FtpUsername string
var FtpUserPass string
var FtpPort string
var FtpDirectory string

func init() {
	rootCmd.AddCommand(ftpCmd)

	ftpCmd.Flags().StringVarP(&FtpHostname, "hostname", "s", "10.1.5.101",
		"--hostname 10.1.5.101")
	ftpCmd.Flags().StringVarP(&FtpUsername, "username", "u", "ftp2",
		"--username MyUsername")
	ftpCmd.Flags().StringVarP(&FtpUserPass, "password", "p", "Password123$$'",
		"--password MyPassword")
	ftpCmd.Flags().StringVarP(&FtpPort, "port", "P", "21", "-P 21")
	ftpCmd.Flags().StringVarP(&FtpDirectory, "directory", "d",
		"\\\\devsrv\\\\ftp_sap\\\\ftp_mii\\\\invoice",
		"--directory my_directory")

	//make flags permanent and required
	//ftpCmd.MarkFlagRequired("hostname")
	//ftpCmd.MarkFlagRequired("username")
	//ftpCmd.MarkFlagRequired("password")
	//ftpCmd.MarkFlagRequired("P")
	//ftpCmd.MarkFlagRequired("directory")
}
