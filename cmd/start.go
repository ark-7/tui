/*
Copyright © 2023 Aadhitya A <aadhitya864@gmail.com>

*/
package cmd

import (
	"log"
	"time"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/briandowns/spinner"
)

var Package bool 

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start eBPF application",
	Long: `Start eBPF application in your system`,
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[25], 100*time.Millisecond)
		s.Suffix = " Starting app ←_←"
		s.FinalMSG = "Operation Complete (～￣▽￣)～\n"
		s.Start()

		if runtime.GOOS == "linux" {
			if Package == true {
				curl := exec.Command("curl", "-L", "-O", "https://github.com/ark-7/arkLB/archive/refs/heads/main.zip")
				_, err := curl.Output()
				if err != nil {
					log.Println("\nPackage downloaded")
				}
				extractZip := exec.Command("unzip", "main.zip")
				_, err = extractZip.Output()

				rm := exec.Command("rm", "main.zip")
				_, err = rm.Output()
			} else {
				curl := exec.Command("curl", "-L", "-O", "https://github.com/ark-7/arkLB/archive/refs/heads/main.zip")
				_, err := curl.Output()
				if err != nil {
					log.Println("\nPackage downloaded")
				}
				start := exec.Command("bash", "./arkLB-main/start.sh")
				_, err = start.Output()
			}
		} else {
			log.Fatal("Project ARK is not supported in your OS")
		}
		s.Stop()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolVarP(&Package, "package", "p", false, "Download app package alone")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
