# kuu

[![CI](https://github.com/winebarrel/kuu/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/kuu/actions/workflows/ci.yml)

Remove blank lines from the beginning and end of the file. Also, make duplicate blank lines into one.

## Installation

```sh
brew install winebarrel/kuu/kuu
```

## Usage

```
Usage: kuu [<files> ...] [flags]

Arguments:
  [<files> ...]    Input files.

Flags:
  -h, --help                     Show help.
  -i, --inplace                  Edit files in place.
  -b, --inplace-backup=STRING    Edit files in place (makes backup with specified suffix).
  -u, --unique                   Make duplicate blank lines unique.
      --version
```

```sh
$ echo -e "\nfoo\nbar\n\nzoo\n"

foo
bar

zoo

$ echo -e "\nfoo\nbar\n\nzoo\n" | kuu
foo
bar

zoo
```

```sh
$ cat london-bridge.txt | nl -ba
     1	
     2
     3	London Bridge is broken down,
     4	Broken down, broken down.
     5
     6	London Bridge is broken down,
     7
     8
     9	My fair lady.
    10
    11

$ kuu london-bridge.txt | nl -ba # or `cat london-bridge.txt | kuu`
     1	London Bridge is broken down,
     2	Broken down, broken down.
     3
     4	London Bridge is broken down,
     5
     6
     7	My fair lady.

$ kuu -i london-bridge.txt # when backing up `kuu -b.bak london-bridge.txt`

$ cat london-bridge.txt | nl -ba
     1	London Bridge is broken down,
     2	Broken down, broken down.
     3
     4	London Bridge is broken down,
     5
     6
     7	My fair lady.
```

### Make duplicate blank lines unique

To make duplicate blank lines unique, use the `-u` option.

```sh
$ cat london-bridge.txt | nl -ba
     1
     2
     3		London Bridge is broken down,
     4		Broken down, broken down.
     5
     6		London Bridge is broken down,
     7
     8
     9		My fair lady.
    10

$ cat london-bridge.txt | kuu | nl -ba
     1		London Bridge is broken down,
     2		Broken down, broken down.
     3
     4		London Bridge is broken down,
     5
     6
     7		My fair lady.

$ cat london-bridge.txt | kuu -u | nl -ba
     1		London Bridge is broken down,
     2		Broken down, broken down.
     3
     4		London Bridge is broken down,
     5
     6		My fair lady.
```
