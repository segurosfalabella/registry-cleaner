
## Registry Cleaner
[![Build Status](https://www.travis-ci.org/segurosfalabella/registry-cleaner.svg?branch=master)](https://www.travis-ci.org/segurosfalabella/registry-cleaner) [![Coverage Status](https://coveralls.io/repos/github/segurosfalabella/registry-cleaner/badge.svg?branch=master)](https://coveralls.io/github/segurosfalabella/registry-cleaner?branch=master)

Tool to delete unused tags in azure container registry.

### Requirements

Before use this tool, you have to install azure-cli and logged in with an service-principal account

```
az login --service-principal -u <azure-client-id> -p <azure-client-secret> --tenant <azure-tenant-id>
```

### Usage

Compile code and generate binary with

```
go build .
```

Then execute the binary with these flags
* `-registry="registry-name"`, registry name
* `-repository="repository-name"`, repository name
* `tag1 tag2 tag3`, add as arguments the tags that you wan to keep or save. The process will prevend deleting these tags

```
$ ./registry-cleaner -registry="<registry>" -repository="<repository>" tag1 tag2 tag3
```

### Collaborators

* Max Guzman <https://github.com/maxguzman>
* Miguel herrera <migueljherrera@gmail.com>