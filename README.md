
## manifest
```bash
# https://github.com/akavel/rsrc
rsrc -manifest ./main.manifest -ico ./src/assets/favicon.ico -o ./src/main.syso
```

## assets
```bash
# https://github.com/go-bindata/go-bindata/
go get -u github.com/go-bindata/go-bindata/...
cd src
go-bindata -o file/assets.go -pkg=assets ./assets
```

## development
```bash
cd src
go build -o ./main.exe && ./main.exe
```

## build
```bash
go build -ldflags="-H windowsgui" ./src/main.go
```