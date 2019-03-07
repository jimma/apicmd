package app

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var CityCmd = &cobra.Command{
	Use:   "info [city to query]",
	Short: "query city aqi",
	Long:  `Query city air quality index.`,
	Args:  cobra.MaximumNArgs(1),
	Run:   Query,
}

type AQI struct {
	/* 	"data": {
	        "idx":7397,
	        "aqi":71,
	        "time":{
	            "v":1481396400,
	            "s":"2016-12-10 19:00:00",
	            "tz":"-06:00"
	        },
	        "city":{
	            "name":"Chi_sp, Illinois",
	            "url":"https://aqicn.org/city/usa/illinois/chi_sp/",
	            "geo":["41.913600","-87.723900"]
	        },
	        "iaqi":{
	            "pm25":{
	                "v":71
	            }
			} */
	Data struct {
		AQI int `json:"aqi"`
	} `json:"data"`
	//Pm21  string `json:"data.iaqi.pm25.v"`
}

func Query(cmd *cobra.Command, args []string) {
	city := args[0]
	url := "https://api.waqi.info/feed/" + city + "/?"
	tokenValue := "eacf929e8553b2ad7364db39ad73ef258e1ab58f"
	res, err := resty.R().SetQueryParam("token", tokenValue).Get(url)
	if err != nil {
		glog.Fatal("Error when send request to retrieve index data :  %s", err)
	} else {
		body := res.Body()
		aqi := AQI{}
		jsonErr := json.Unmarshal(body, &aqi)
		if jsonErr != nil {
			glog.Fatal("Umarshal json data err", jsonErr)
		}
		fmt.Printf("City=%s, AQI=%d\n", city, aqi.Data.AQI)
	}

}
