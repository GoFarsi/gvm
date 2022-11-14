# gvm
Go version manager tool, for install, upgrade or downgrade easy you golang.

# install

1. install with go:
```shell
go install github.com/GoFarsi/gvm@latest
```

2. install from [release page](https://github.com/GoFarsi/gvm/releases).

# Example
example commands of gvm tool
### install
install latest version of Go
```shell
gvm install
```
install with backup downloaded Go compiler
```shell
gvm install -b /home/{user}
```
install specific version of Go compiler
```shell
gvm install -v 1.19.3
```

### upgrade
upgrade to latest version Go
```shell
gvm upgrade
```
upgrade with backup downloaded Go compiler
```shell
gvm upgrade -b /home/{user}
```
upgrade to specific version of Go compiler
```shell
gvm upgrade -v 1.19.3
```

### downgrade
downgrade to previous version of Go
```shell
gvm downgrade
```
downgrade with backup downloaded Go compiler
```shell
gvm downgrade -b /home/{user}
```
downgrade to specific version of Go compiler
```shell
gvm upgrade -v 1.19.2
```

### list
list of version with commit id Go compiler
```shell
gvm list
```
list of version with commit id Go compiler in 10 line
```shell
gvm list -l 10
```

### release
show release note link all version of Go compiler
```shell
gvm release
```
show release note specific of version Go compiler
```shell
gvm release -v 1.19
```

### notification
check new version of Go compiler base on your installed version
```shell
gvm notification
```
**note**:
> you can set cronjob for check new version Go compiler using command :
> `(crontab -l ; echo "* */6 * * * bash gvm notification >/dev/null 2>&1")| crontab -` every 6 hours.

# TODO
- [x] gvm for linux
- [ ] mac os
- [ ] windows
- [ ] test code
