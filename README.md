Used as command runner for the commands running in background
for [Estuary-Agent](https://github.com/dinuta/estuary-agent)

## runcmd

Examples:

```bash
runcmd --args="ls;;echo 4"
runcmd --cid=2 --args="ls;;echo 4"
runcmd --cid myId --args "ls;;echo 4"
runcmd --cid myId --args "ls;;echo 4" --enableStreams=true
```

## For private pkg

```access transformers
go env -w GOPRIVATE=github.com/estuaryoss/*
```
