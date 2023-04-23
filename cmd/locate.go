package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/ipinfo/go/ipinfo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(locateCmd)
}

var locateCmd = &cobra.Command{
	Use:   "locate",
	Short: "Locate the IP with this command.",
	Long: `Locate the IP with this command, 
	if no IP is provided, your own IP information will be printed.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("locate called")
		infos, err := locate(args...)
		if err != nil {
			log.Panicln(err)
			return
		}
		showLocationInfos(infos)
	},
}

func showLocationInfos(infos []*ipinfo.Info) {
	for _, v := range infos {
		fmt.Printf("IP: %s\n", v.IP.String())
		fmt.Printf("City: %s\n", v.City)
		fmt.Printf("Region: %s\n", v.Region)
		fmt.Printf("Country: %s\n", v.Country)
		fmt.Printf("Location: %s\n", v.Location)
		fmt.Printf("Phone: %s\n", v.Phone)
		fmt.Printf("Postal: %s\n", v.Postal)
		fmt.Printf("Hostname: %s\n", v.Hostname)
		fmt.Printf("Organization: %s\n", v.Organization)
		fmt.Println()
	}
}

// locate 获取 IP 位置信息
func locate(ips ...string) (infos []*ipinfo.Info, err error) {
	client := ipinfo.NewClient(nil)
	var info *ipinfo.Info
	if len(ips) == 0 { // 没有提供ip地址
		info, err = client.GetInfo(nil)
		infos = append(infos, info)
		return
	} else {
		for _, ip := range ips {
			info, err = client.GetInfo(net.ParseIP(ip))
			infos = append(infos, info)
		}
		return
	}
}
