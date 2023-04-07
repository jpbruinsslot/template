Template
========

Use environment variables in Go templates.

Installation
------------

#### Binary installation

[Download](https://github.com/jpbruinsslot/template/releases) a
compatible binary for your system. For convenience, place `template` in a
directory where you can access it from the command line.

#### Via Go:

```
$ go get -u github.com/jpbruinsslot/template
```

Usage
-----

#### Command line usage:

```
NAME:
    template - use environment variables and data files (like .env) in Go templates

USAGE:
    template -t [template-file] -o [output-file] -d [data-file]

EXAMPLES:

    # Use system wide environment variables
    $ template -t input.tpml -o output.txt
    $ template -t input.tmpl > output.txt

    # Use data files (support for env and json files)
    $ template -t input.tpml -o output.txt -d data.env
    $ template -t input.tpml -o output.txt -d data.json

    # Use stdin for template file
    $ cat input.tmpl | template -o output.txt                   # output txt file
    $ cat input.tmpl | template -o -                            # output to stdout
    $ cat input.tmpl | template                                 # output to stdout

    # Use stdin for data file
    $ cat data.env | template -t input.tmpl -o output.txt -d -  # output txt file
    $ cat data.env | template -t input.tmpl -d -                # output to stdout

VERSION:
    0.2.1

WEBSITE:
    https://github.com/jpbruinsslot/template		

GLOBAL OPTIONS:
    -t, -template [template-file]     template file
    -o, -output [output-file]         output file
    -d, -data [data-file]             data file
    -h, -help
```
