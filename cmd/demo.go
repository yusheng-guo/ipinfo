package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/ipinfo/go/ipinfo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(demoCmd)
}

func demo() {
	// const token = "fa204583ea0443"
	client := ipinfo.NewClient(nil)
	info, err := client.GetInfo(net.ParseIP("8.8.8.8"))
	// info, err := client.GetIP(nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(info)
}

// demoCmd represents the demo command
var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "A Demo to get information about domain 8.8.8.8.",
	Long:  `A Demo to get information about domain 8.8.8.8.`,
	Run: func(cmd *cobra.Command, args []string) {
		demo()
	},
}
