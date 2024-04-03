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
$ echo -e "\n1\n2\n\n3\n"

1
2

3

$ echo -e "\n1\n2\n\n3\n" | ./kuu
1
2

3
$ (echo '---' ; cat london-bridge.txt ; echo '---') | nl -b a
     1	---
     2
     3
     4	London Bridge is broken down,
     5	Broken down, broken down.
     6
     7	London Bridge is broken down,
     8
     9
    10	My fair lady.
    11
    12
    13	---
$ kuu london-bridge.txt # or `cat london-bridge.txt | kuu`
London Bridge is broken down,
Broken down, broken down.

London Bridge is broken down,


My fair lady.
$ kuu -i london-bridge.txt # when backing up `kuu -b.bak london-bridge.txt`
$ (echo '---' ; cat london-bridge.txt ; echo '---') | nl -b a
     1	---
     2	London Bridge is broken down,
     3	Broken down, broken down.
     4
     5	London Bridge is broken down,
     6
     7
     8	My fair lady.
     9	---
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
$ cat london-bridge.txt | ./kuu | nl -ba
     1		London Bridge is broken down,
     2		Broken down, broken down.
     3
     4		London Bridge is broken down,
     5
     6
     7		My fair lady.
$ cat london-bridge.txt | ./kuu -u | nl -ba
     1		London Bridge is broken down,
     2		Broken down, broken down.
     3
     4		London Bridge is broken down,
     5
     6		My fair lady.
```
