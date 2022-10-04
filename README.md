# blocksort

A tool for sorting blocks of lines.

## Install
```
go install github.com/tomnomnom/blocksort@latest
```

## Usage

Given some blocks of lines on `stdin`:

```
▶ cat blocks.txt
Block C
    Item 2
    Item 1
Block A
    Item 3
    Item 2
Block B
    Item 5
    Item 3
```

They can be sorted by a specific line (defaulting to the first line):

```
▶ cat blocks.txt | blocksort -n3
Block A
    Item 3
    Item 2
Block B
    Item 5
    Item 3
Block C
    Item 2
    Item 1
```

The `-n`/`--number` flag sets how many lines are in each block.
The `-i`/`--index` flag sets which line in the block to use when sorting.

Get help with `-h`:

```
▶ blocksort -h
Usage of blocksort:
  -i, --index int    Which line to sort by (0-indexed)
  -n, --number int   Number of lines per block (default 3)
  -r, --reverse      Reverse sort order
pflag: help requested
```


## Credit
Built for [@digininja](https://github.com/digininja) as a result of
[this Twitter thread](https://twitter.com/digininja/status/1577331330567766016).
