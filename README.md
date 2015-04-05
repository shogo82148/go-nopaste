# go-nopaste
very simple nopaste written in golang.

## Install

```
go get github.com/shogo82148/go-nopaste/cmd/nopaste
```

## Usage

```
$ nopaste -config config.yaml
$ curl -v -F text=hogehoge http://localhost:3000/
...
< Location: /3b2c6c10d0
...
$ curl http://localhost:3000/3b2c6c10d0
hogehoge
```

## Configuration

```
root: /
listen: ":3000"
data_dir: data
```
