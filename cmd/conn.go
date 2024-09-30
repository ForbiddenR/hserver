/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"github.com/spf13/cobra"
	"golang.org/x/net/http2"
)

var (
	h2 bool
	h3 bool
	address string
)

// connCmd represents the conn command
var connCmd = &cobra.Command{
	Use:   "conn",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		connRun()
	},
}

func init() {
	rootCmd.AddCommand(connCmd)

	connCmd.Flags().BoolVar(&h2, "h2", false, "enable h2")
	connCmd.Flags().BoolVar(&h3, "h3", false, "enable h3")
	connCmd.Flags().StringVarP(&address, "address", "a", "https://www.google.com" , "the target address")
}

type Client interface {
	Get(string) (*http.Response, error)
}

func connRun() {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
	}
	var roundTripper http.RoundTripper
	if h2 {
		roundTripper = &http2.Transport{
			TLSClientConfig: tlsConfig,
		}
	} else if h3 {
		roundTripper = &http3.RoundTripper{
			TLSClientConfig: tlsConfig,
			QUICConfig: &quic.Config{},
		}
	}
	client := &http.Client{
		Transport: roundTripper,
	}
	resp, err := client.Get(address)
	if err != nil {
		panic(err)
	}
	buffer := &bytes.Buffer{}
	io.Copy(buffer, resp.Body)
	fmt.Printf("%s\n", buffer.Bytes())
}
