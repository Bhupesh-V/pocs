Sample Go project to demonstrate how to build a Go binary and print build information.

## Pre-requisites

- Go
- Make

## Build

```bash
make build
```

Sample output:

```bash
-buildmode exe
-compiler gc
DefaultGODEBUG asynctimerchan=1,gotypesalias=0,httplaxcontentlength=1,httpmuxgo121=1,httpservecontentkeepheaders=1,tls10server=1,tls3des=1,tlskyber=0,tlsrsakex=1,tlsunsafeekm=1,winreadlinkvolume=0,winsymlink=0,x509keypairleaf=0,x509negativeserial=1
CGO_ENABLED 1
CGO_CFLAGS 
CGO_CPPFLAGS 
CGO_CXXFLAGS 
CGO_LDFLAGS 
GOARCH arm64
GOOS darwin
GOARM64 v8.0
vcs git
vcs.revision 7034e09433ee73af065c839f54f84ddbf1f788da
vcs.time 2025-01-11T11:29:50Z
vcs.modified false
```
