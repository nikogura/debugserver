/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
)

var port int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "debugserver",
	Short: "Simple HTTP debug server",
	Long: `
Simple HTTP debug server.

Just prints out what you send it.  Useful for debugging http clients that may or may not do what their docs say they do.
`,
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			dump, err := httputil.DumpRequest(r, true)
			if err != nil {
				http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
				return
			}

			fmt.Printf("%s", dump)
		})

		if len(args) > 0 {
			port, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("Can't convert %d to an integer", port)
			}
		}

		addr := fmt.Sprintf(":%d", port)

		fmt.Printf("Debug Server Listening on %s\nAll requests will be logged here.\n", addr)

		http.ListenAndServe(addr, nil)
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
	rootCmd.Flags().IntVarP(&port, "port", "p", 8888, "Port on which to listen")
}
