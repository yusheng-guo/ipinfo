package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// {
// 	"ip": "119.91.204.226",
// 	"city": "Shenzhen",
// 	"region": "Guangdong",
// 	"country": "CN",
// 	"loc": "22.5455,114.0683",
// 	"org": "AS45090 Shenzhen Tencent Computer Systems Company Limited",
// 	"timezone": "Asia/Shanghai",
// 	"readme": "https://ipinfo.io/missingauth"
//   }

type IPinfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

// trackCmd represents the track command
var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Track the IP with this command.",
	Long:  `Track the IP with this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("track called")
		if len(args) > 0 {
			for _, ipAddress := range args {
				showInfo(getInfo(ipAddress))
			}
		} else {
			fmt.Println("Please provide an IP Address to track.")
		}
	},
}

func showInfo(ipinfo IPinfo) {
	fmt.Println("IP:", ipinfo.IP)
	fmt.Println("City:", ipinfo.City)
	fmt.Println("Country:", ipinfo.Country)
	fmt.Println("Loc:", ipinfo.Loc)
	fmt.Println("Org:", ipinfo.Org)
	fmt.Println("Timezone:", ipinfo.Timezone)
	fmt.Println("Readme:", ipinfo.Readme)
}

func getInfo(ipAddress string) IPinfo {
	url := "https://ipinfo.io/" + ipAddress + "/geo"
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var ipinfo IPinfo // 将 json 数据映射为 IPinfo结构体
	if resp.StatusCode == 200 {
		json.Unmarshal(body, &ipinfo)
	}
	return ipinfo
}

func init() {
	rootCmd.AddCommand(trackCmd)
}
