/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

// randomStrCmd represents the randomStr command
var randomStrCmd = &cobra.Command{
	Use:   "randomStr",
	Short: "Generate random string",
	Long: `Provided length, it generates random string.
For example:
generateX randomStr -l 10`,
	Run: func(cmd *cobra.Command, args []string) {
		length, err := cmd.Flags().GetInt("length")
		str := randomString(length)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(str)
		}
	},
}

func randomString(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func init() {
	rootCmd.AddCommand(randomStrCmd)
	randomStrCmd.Flags().IntP("length", "l", 32, "length of random string")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomStrCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomStrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
