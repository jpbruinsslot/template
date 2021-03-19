Template
========

Use environment variables in Go templates.

Installation
------------

#### Via Go:

```
$ go get -u github.com/erroneousboat/template
```

#### Via Docker:

```
$ docker run --rm -it erroneousboat/template
```


Usage
-----

#### Command line usage:

```
NAME:
    template - use environment variables in Go templates

USAGE:
    template -i [input-file] -o [output-file]

EXAMPLES:

    $ template -i input.tpml -o output.txt

    $ echo "{{ .PWD }}" | template -o output.txt

    $ echo "{{ .PWD }}" | template

    $ template -i input.tmpl > output.txt

VERSION:
        0.1.0

WEBSITE:
        https://github.com/erroneousboat/template

GLOBAL OPTIONS:
    -i, -input [input-file]     input file
    -o, -output [output-file]   output file
    -h, -help
```
