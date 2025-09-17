package main


import (
	"flag"
	"fmt"
	"os"

	"github.com/Astraxx04/godirview/internal/visualizer"
)

func main() {
	// Flags
	depth := flag.Int("L", -1, "Descend only level directories deep")
	showAll := flag.Bool("a", false, "Show hidden files")
	jsonOut := flag.Bool("json", false, "Output JSON instead of tree view")

	flag.Parse()

	// Get path
	args := flag.Args()
	path := "."
	if len(args) > 0 {
		path = args[0]
	}

	// Run visualizer
	err := visualizer.Run(path, *depth, *showAll, *jsonOut)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
