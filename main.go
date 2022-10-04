package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/spf13/pflag"
)

type options struct {
	BlockSize int
	SortIndex int

	Reverse bool
}

func main() {
	var opts options
	pflag.IntVarP(&opts.BlockSize, "number", "n", 3, "Number of lines per block")
	pflag.IntVarP(&opts.SortIndex, "index", "i", 0, "Which line to sort by (0-indexed; default 0)")
	pflag.BoolVarP(&opts.Reverse, "reverse", "r", false, "Reverse sort order")
	pflag.Parse()

	if opts.SortIndex >= opts.BlockSize || opts.SortIndex < 0 {
		fmt.Fprintln(os.Stderr, "index cannot be negative, or exceed number of lines per block")
		return
	}

	sc := bufio.NewScanner(os.Stdin)

	blocks := newBlockList(opts.SortIndex, opts.Reverse)

	for {
		block := make([]string, 0)
		for i := 0; i < opts.BlockSize; i++ {
			if !sc.Scan() {
				break
			}
			line := sc.Text()
			block = append(block, line)
		}
		if len(block) == 0 {
			break
		}
		blocks.append(block)
	}

	sort.Sort(blocks)

	for _, block := range blocks.blocks {
		for _, line := range block {
			fmt.Println(line)
		}
	}
}

type blockList struct {
	blocks    [][]string
	sortIndex int
	reverse   bool
}

func newBlockList(i int, r bool) *blockList {
	return &blockList{
		blocks:    make([][]string, 0),
		sortIndex: i,
		reverse:   r,
	}
}

func (b *blockList) append(block []string) {
	b.blocks = append(b.blocks, block)
}

func (b *blockList) Len() int {
	return len(b.blocks)
}

func (b *blockList) Less(i, j int) bool {
	if b.reverse {
		i, j = j, i
	}
	return b.blocks[i][b.sortIndex] < b.blocks[j][b.sortIndex]
}

func (b *blockList) Swap(i, j int) {
	b.blocks[i], b.blocks[j] = b.blocks[j], b.blocks[i]
}
