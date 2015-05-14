Template
-------

Create config files from templates files using environment variables as
context.

```bash
# Usage: template [TEMPLATE FILE]:[CONFIG FILE]
$ template config.tmpl:config.conf
```

You can reference environment variables in you template files with the
following notation:

```
I'm using {{ .Env.SHELL }} as my shell
```
