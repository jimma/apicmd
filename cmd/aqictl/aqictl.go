package main

import (
	goflag "flag"

	"github.com/golang/glog"
	"github.com/jimma/aqicmd/cmd/app"
	"github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
)

var flagvar int

func init() {
	pflag.IntVar(&flagvar, "flagname", 10, "help message for flagname")
}

//type executefn func(cmd *cobra.Command, args []string)

func main() {
	//add goflag command line to pflag
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()
	glog.Info("command flag is parsed")
	var rootCmd = &cobra.Command{Use: "aqictl",
		Short: "Utility to query aqi",
		Long:  "Command line to query city's air quality index",
		//Run:   Execute,
	}

	rootCmd.AddCommand(app.CityCmd, app.RemoveCmd)

	rootCmd.Execute()
}

func Execute(cmd *cobra.Command, args []string) {
	glog.Info("Execute root method")
}
