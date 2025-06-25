/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/leonardoTavaresM/algocli/internal/sort"
	"github.com/spf13/cobra"
)

// Flags
var input string
var alg string

// sortCmd represents the sort command
var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "Ordena uma lista de números",
	Run: func(cmd *cobra.Command, args []string) {
		alg, _ := cmd.Flags().GetString("alg")
		inputStr, _ := cmd.Flags().GetString("input")
		values := parseInput(inputStr)

		start := time.Now()

		var result []int
		switch alg {
		case "bubble":
			result = sort.BubbleSort(values)
		case "quick":
			result = sort.QuickSort(values)
		default:
			fmt.Println("Algoritmo não suportado. Use --alg bubble ou quick.")
			return
		}

		elapsed := time.Since(start)

		fmt.Println("Resultado:", result)
		fmt.Println("tempo do algoritmo:", elapsed)
	},
}

func parseInput(input string) []int {
	var out []int
	for _, s := range strings.Split(input, ",") {
		var n int
		fmt.Sscanf(s, "%d", &n)
		out = append(out, n)
	}
	return out
}

func init() {
	rootCmd.AddCommand(sortCmd)
	sortCmd.Flags().String("alg", "bubble", "Algoritmo de ordenação: bubble, quick")
	sortCmd.Flags().String("input", "", "Lista de números separada por vírgula (ex: 3,1,2)")
	sortCmd.MarkFlagRequired("input")
}
