package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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
	- basic stat buffs.`,
	Example: `mhp3rd_weapon_calculator <weapon damage: int> <sharpness: string> <affinity: int>
	Weapon damage should be just weapon damage, not attack from player status.
	Sharpness values: r (red), o (orange), y (yellow), g (green), b (blue), w (white).
	Affinity should be passed as percentage value, negative affinity should start with 'n', eg: 'n20' is -20% affinity`,
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
		affinity, _ := strconv.Atoi(strings.Replace(args[2], "n", "-", 1))

		calculator.Damage_calculator(weapon_damage, sharpness, affinity)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
