/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quic-go/quic-go/http3"
	"github.com/spf13/cobra"
)

var (
	certificatePath string
	keyPath         string
	path            string
	port            int16
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	startCmd.Flags().StringVarP(&certificatePath, "cert", "c", "", "the path of certificate file")
	startCmd.Flags().StringVarP(&keyPath, "key", "k", "", "the path of key file")
	startCmd.Flags().StringVarP(&path, "path", "p", "", "the path of key and certificate files")
	startCmd.Flags().Int16VarP(&port, "port", "o", 0, "the port expose to public")
}

func run() {
	r := gin.Default()
	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, `{"status":"ok"}`)
	})
	po := fmt.Sprintf(":%d", port)
	// err := http3.ListenAndServeQUIC(po, path+"/server.pem", path+"/server.key", mux)
	err := http3.ListenAndServeTLS(po, path+"/server.pem", path+"/server.key", r.Handler())
	if err != nil {
		panic(err)
	}
}
