## mew
---
mew shows Go packages that used in your repo


### Usage

```shell
❯ mew -h
mew - Show your Go repo related pkgs

Usage:
  mew [flags]

Flags:
  -e, --ed stringArray   exclude the dir
  -g, --grep string      grep the pkg list
  -h, --help             help for mew
  -r, --repo string      input repo name
```

### e.g.

```shell
❯ mew -r github.com/scbizu/mew  -g mew -e vendor -e filter -e .git

```
