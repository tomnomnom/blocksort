# blocksort

A tool for sorting blocks of lines.

## Usage

Given some blocks of lines on `stdin` (in this example there 3 lines in each block):

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

They can be sorted by a given line (defaulting to the first line):

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

Get help with `-h`:

```
▶ blocksort -h
Usage of blocksort:
  -i, --index int    Which line to sort by (0-indexed)
  -n, --number int   Number of lines per block (default 3)
  -r, --reverse      Reverse sort order
pflag: help requested
```

## Install
```
go install github.com/tomnomnom/blocksort@latest
```

## Credit
Built for @digininja because of [this Twitter thread](https://twitter.com/digininja/status/1577331330567766016).
