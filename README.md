Template
========

Create config files from templates using environment variables as context.
This will be very helpful for the creation of config files inside
docker containers that only have environment variables available to them.

Installation
------------

Two binaries are available in the `bin/` folder. You can download template from
there, but if you want to download them for your docker container you can use
the following command:

```bash
wget https://github.com/erroneousboat/template/raw/master/bin/template-linux-amd64
chmod +x template-linux-amd64
```

Usage
-----

Specify the `TEMPLATE_FILE` and the resulting `CONFIG_FILE`, in the following
command:

```bash
# Usage: template [TEMPLATE FILE]:[CONFIG FILE]
$ template config.tmpl:config.conf
```

You can reference environment variables in you template files with the
following notation:

```
I'm using {{ .Env.SHELL }} as my shell
```

And this will ouput:

```
I'm using /bin/bash as my shell
```
