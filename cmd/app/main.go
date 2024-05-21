package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// n^2

func CheckBallsDistribution(in *bufio.Reader, out *bufio.Writer) error {
	var n int
	if _, err := fmt.Fscanf(in, "%d\n", &n); err != nil {
		return fmt.Errorf("failed to read n: %w", err)
	}

	sumRows := make([]int, n)
	sumCols := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var ballsCount int
			if _, err := fmt.Fscanf(in, "%d", &ballsCount); err != nil {
				return fmt.Errorf("failed to read balls count: %w", err)
			}

			sumRows[i] += ballsCount
			sumCols[j] += ballsCount
		}

		if _, err := fmt.Fscanf(in, "\n"); err != nil {
			return fmt.Errorf("failed to read new line: %w", err)
		}
	}

	sort.Ints(sumRows)
	sort.Ints(sumCols)

	for i := 0; i < n; i++ {
		if sumRows[i] != sumCols[i] {
			fmt.Fprint(out, "no")
			return nil
		}
	}

	fmt.Fprint(out, "yes")
	return nil
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	err := CheckBallsDistribution(in, out)
	if err != nil {
		fmt.Fprintf(out, "error: %v\n", err)
	}
}
