# kuu

[![CI](https://github.com/winebarrel/kuu/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/kuu/actions/workflows/ci.yml)

Remove blank lines from the beginning and end of the file. Also, make duplicate blank lines into one.

## Usage

```
Usage: kuu [<files> ...] [flags]

Arguments:
  [<files> ...]    Input files.

Flags:
  -h, --help                     Show help.
  -i, --inplace                  Edit files in place.
  -b, --inplace-backup=STRING    Edit files in place (makes backup with specified suffix).
      --version
```

```sh
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
     7	My fair lady.
     8	---
```
