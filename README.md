
## Registry Cleaner

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
* `-repository="repository-name"`, repository name
* `tag1 tag2 tag3`, add as arguments the tags that you wan to keep or save. The process will prevend deleting these tags

```
$ ./registry-cleaner -repository="<repository>" tag1 tag2 tag3
```

### Collaborators

* Max Guzman <https://github.com/maxguzman>
* Miguel herrera <migueljherrera@gmail.com>