package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tbonus/mhp3rd_damage_calculator/calculator"
)

var rootCmd = &cobra.Command{
	Use:   "mhp3rd_damage_calculator",
	Short: "Calculate average weapon damage based on its statistics",
	Long: `Calculate average weapon damage based on its statistics taking into account:
	- attack power
	- affinity
	- sharpness
	- stat buffs.`,
	Example: `mhp3rd_weapon_calculator <weapon damage: int> <sharpness: string> <affinity: int>
	Sharpness values: r (red), o (orange), y (yellow), g (green), b (blue), w (white)`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println(cmd.Long)
			fmt.Println("Usage:", cmd.Example)
			return
		}
		if len(args) < 3 {
			fmt.Printf("Not enough arguments.\nUsage: %v\n", cmd.Example)
			return
		}

		weapon_damage, _ := strconv.Atoi(args[0])
		sharpness := args[1]
		affinity, _ := strconv.Atoi(args[2])

		calculator.Damage_calculator(weapon_damage, sharpness, affinity)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
