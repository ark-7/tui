/*
Copyright © 2023 Aadhitya A <aadhitya864@gmail.com>

*/
package cmd

import (
	"fmt"
	"log"
	"time"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/briandowns/spinner"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install dependencies",
	Long: `Install dependencies for project ARK`,
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[25], 100*time.Millisecond)
		s.Suffix = " Installing dependencies ←_←"
		s.FinalMSG = "Complete! q(≧▽≦q)\n"
		s.Start()
		// Install dependencies
		if runtime.GOOS == "linux" {
			curl := exec.Command("curl", "-O", "https://raw.githubusercontent.com/ark-7/arkLB/main/install.sh")
			_, err := curl.Output()
			if err != nil {
				log.Println("\nScript downloaded")
			}
			
			install := exec.Command("bash", "install.sh")
			_, err = install.Output()

			rm := exec.Command("rm", "-f", "install.sh")
			_, err = rm.Output()
		} else {
			fmt.Println("This project is only supported in linux OS. :(")
			os.Exit(1)
		}
		s.Stop()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
