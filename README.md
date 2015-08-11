Template
-------

Create config files from templates files using environment variables as
context. This will be very helpful for the creation of config files inside
docker containers that only have environment variables available to them,

# Installation

```bash
wget https://github.com/erroneousboat/template/raw/master/bin/template
chmod +x template
```

# Usage

```bash
# Usage: template [TEMPLATE FILE]:[CONFIG FILE]
$ template config.tmpl:config.conf
```

You can reference environment variables in you template files with the
following notation:

```
I'm using {{ .Env.SHELL }} as my shell
```
