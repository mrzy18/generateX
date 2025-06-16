/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "A brief description of your command",
	Long: `Provided length, it generates base64 string.
For example:
generateX base64 -l 10`,
	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")
		str, err := generateBase64String(length)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(str)
		}
	},
}

func generateBase64String(length int) (string, error) {
	// generate bytes
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	base64String := base64.StdEncoding.EncodeToString(bytes)
	return base64String, nil
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.Flags().IntP("length", "l", 8, "length of base64")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// base64Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// base64Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
