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
    template -t [template-file] -o [output-file] -d [data-file]

EXAMPLES:

    $ template -t input.tpml -o output.txt

    $ template -t input.tpml -o output.txt -d data.json

    $ echo "{{ .PWD }}" | template -o output.txt

    $ echo "{{ .PWD }}" | template

    $ template -t input.tmpl > output.txt

VERSION:
    0.2.0

WEBSITE:
    https://github.com/erroneousboat/template		

GLOBAL OPTIONS:
    -t, -template [template-file]     template file
    -o, -output [output-file]   	  output file
	-d, -data [data-file]	    	  data file
    -h, -help
```
