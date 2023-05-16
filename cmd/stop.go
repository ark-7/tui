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

// stopCmd represents the start command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop eBPF application",
	Long: `Stop eBPF application in your system`,
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[25], 100*time.Millisecond)
		s.Suffix = " Stopping app ←_←"
		s.FinalMSG = "Operation Complete (～￣▽￣)～\n"
		s.Start()

		if runtime.GOOS == "linux" {
			if Package == true {
				stop := exec.Command("cd", "arkLB-main")
				_, err := stop.Output()

				make := exec.Command("make", "clean")
				_, err = make.Output()

				rm := exec.Command("rm", "-rf", "arkLB-main")
				_, err = rm.Output()
				if err != nil {
					log.Println("\nPackage removed")
				}
			} else {
				stop := exec.Command("cd", "arkLB-main")
				_, err := stop.Output()

				make := exec.Command("make", "clean")
				_, err = make.Output()

				if err != nil {
					log.Println("\nApp stopped")
				}
			}
		} else {
			log.Fatal("Project ARK is not supported in your OS")
		}
		s.Stop()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
	stopCmd.Flags().BoolVarP(&Package, "package", "p", false, "Stop and remove app package")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
