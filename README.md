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
      --deep             [Experimental feature]in deep mode,you will get all(include really all dependency) third party related pkg name
  -e, --ed stringArray   exclude the dir (default [vendor,.git])
  -d, --graph string     dump graphviz graph (default "mew.dot")
  -g, --grep string      grep the pkg list
  -h, --help             help for mew
      --json             show json format
  -r, --repo string      input repo name
```

### e.g.

* with normal mode

```shell
mew -r github.com/scbizu/mew -g 'github'  -e '.git' -e 'vendor' --json -d 'mew.dot' |jq .
[
  "github.com/scbizu/mew/cmd",
  "github.com/scbizu/mew/drawer",
  "github.com/scbizu/mew/filter",
  "github.com/scbizu/mew/linker",
  "github.com/sirupsen/logrus",
  "github.com/spf13/cobra",
  "github.com/awalterschulze/gographviz"
]
```

* with deep mode

```shell
❯ mew -r github.com/scbizu/mew -g 'github'  -e '.git' -e 'vendor' --json -d 'mew.dot' --deep

{
  "github.com/awalterschulze/gographviz": [
    "github.com/awalterschulze/gographviz/ast",
    "github.com/awalterschulze/gographviz/internal/parser"
  ],
  "github.com/awalterschulze/gographviz/ast": [
    "github.com/awalterschulze/gographviz/internal/token"
  ],
  "github.com/awalterschulze/gographviz/internal/parser": [
    "github.com/awalterschulze/gographviz/internal/errors",
    "github.com/awalterschulze/gographviz/internal/lexer"
  ],
  "github.com/coreos/bbolt": [],
  "github.com/coreos/etcd/auth": [
    "github.com/dgrijalva/jwt-go",
    "github.com/coreos/etcd/pkg/adt",
    "github.com/coreos/etcd/mvcc/backend"
  ],
  "github.com/coreos/etcd/client": [
    "github.com/coreos/etcd/version",
    "github.com/coreos/etcd/pkg/testutil",
    "github.com/coreos/etcd/pkg/pathutil",
    "github.com/ugorji/go/codec",
    "github.com/coreos/etcd/pkg/types",
    "github.com/coreos/etcd/pkg/srv",
    "github.com/coreos/etcd/integration"
  ],
  "github.com/coreos/etcd/clientv3": [
    "github.com/coreos/etcd/pkg/logutil",
    "github.com/coreos/etcd/etcdserver/etcdserverpb",
    "github.com/coreos/etcd/clientv3/balancer",
    "github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes",
    "github.com/coreos/etcd/pkg/transport",
    "github.com/grpc-ecosystem/go-grpc-prometheus",
    "github.com/coreos/etcd/clientv3/clientv3util",
    "github.com/coreos/etcd/clientv3/concurrency",
    "github.com/coreos/etcd/embed",
    "github.com/coreos/etcd/clientv3/mirror",
    "github.com/coreos/etcd/clientv3/namespace",
    "github.com/coreos/etcd/clientv3/leasing"
  ],
  "github.com/coreos/etcd/embed": [
    "github.com/coreos/etcd/etcdserver",
    "github.com/coreos/etcd/etcdserver/api/v2http",
    "github.com/coreos/etcd/etcdserver/api/v3client",
    "github.com/coreos/etcd/pkg/debugutil",
    "github.com/soheilhy/cmux",
    "github.com/ghodss/yaml",
    "github.com/coreos/etcd/pkg/flags",
    "github.com/coreos/etcd/etcdserver/api/v3election",
    "github.com/coreos/etcd/etcdserver/api/v3election/v3electionpb/gw",
    "github.com/coreos/etcd/etcdserver/api/v3lock",
    "github.com/coreos/etcd/etcdserver/api/v3lock/v3lockpb/gw",
    "github.com/coreos/etcd/etcdserver/etcdserverpb/gw",
    "github.com/tmc/grpc-websocket-proxy/wsproxy"
  ],
  "github.com/coreos/etcd/etcdserver": [
    "github.com/coreos/etcd/auth",
    "github.com/coreos/etcd/lease",
    "github.com/coreos/etcd/mvcc",
    "github.com/coreos/etcd/etcdserver/api/membership",
    "github.com/coreos/etcd/etcdserver/api/v2discovery",
    "github.com/coreos/etcd/etcdserver/api/v3alarm",
    "github.com/coreos/etcd/etcdserver/api/v3compactor",
    "github.com/coreos/etcd/pkg/idutil",
    "github.com/coreos/etcd/pkg/runtime",
    "github.com/coreos/etcd/pkg/wait",
    "github.com/coreos/etcd/wal",
    "github.com/coreos/etcd/pkg/mock/mockstorage",
    "github.com/coreos/etcd/pkg/contention",
    "github.com/coreos/etcd/pkg/mock/mockwait"
  ],
  "github.com/coreos/etcd/etcdserver/api": [
    "github.com/coreos/etcd/etcdserver/api/rafthttp",
    "github.com/coreos/etcd/lease/leasehttp",
    "github.com/coreos/etcd/etcdserver/api/v2http/httptypes",
    "github.com/coreos/etcd/pkg/netutil",
    "github.com/coreos/etcd/pkg/mock/mockstore",
    "github.com/coreos/etcd/etcdserver/api/etcdhttp",
    "github.com/coreos/etcd/etcdserver/api/v2auth",
    "github.com/coreos/etcd/etcdserver/api/v3rpc",
    "github.com/coreos/etcd/proxy/grpcproxy/adapter"
  ],
  "github.com/coreos/etcd/etcdserver/api/membership": [
    "github.com/coreos/etcd/etcdserver/api/v2store"
  ],
  "github.com/coreos/etcd/etcdserver/api/rafthttp": [
    "github.com/coreos/etcd/etcdserver/api/snap",
    "github.com/coreos/etcd/etcdserver/api/v2stats",
    "github.com/xiang90/probing"
  ],
  "github.com/coreos/etcd/etcdserver/api/snap": [
    "github.com/coreos/etcd/pkg/fileutil",
    "github.com/coreos/etcd/pkg/ioutil",
    "github.com/coreos/etcd/etcdserver/api/snap/snappb",
    "github.com/coreos/etcd/pkg/pbutil"
  ],
  "github.com/coreos/etcd/etcdserver/api/v2store": [
    "github.com/coreos/etcd/etcdserver/api/v2error",
    "github.com/jonboulle/clockwork",
    "github.com/coreos/etcd/etcdserver/api/v2v3"
  ],
  "github.com/coreos/etcd/etcdserver/api/v2v3": [
    "github.com/coreos/etcd/etcdserver/api"
  ],
  "github.com/coreos/etcd/etcdserver/etcdserverpb": [
    "github.com/coreos/etcd/mvcc/mvccpb",
    "github.com/coreos/etcd/auth/authpb",
    "github.com/grpc-ecosystem/grpc-gateway/runtime"
  ],
  "github.com/coreos/etcd/integration": [
    "github.com/coreos/etcd/clientv3",
    "github.com/coreos/etcd/contrib/recipes",
    "github.com/coreos/etcd/proxy/grpcproxy"
  ],
  "github.com/coreos/etcd/lease": [
    "github.com/coreos/etcd/lease/leasepb",
    "github.com/coreos/etcd/pkg/httputil"
  ],
  "github.com/coreos/etcd/mvcc": [
    "github.com/coreos/etcd/pkg/schedule",
    "github.com/google/btree"
  ],
  "github.com/coreos/etcd/mvcc/backend": [
    "github.com/coreos/bbolt",
    "github.com/dustin/go-humanize"
  ],
  "github.com/coreos/etcd/pkg/logutil": [
    "github.com/coreos/pkg/capnslog",
    "github.com/coreos/etcd/pkg/systemd",
    "github.com/coreos/etcd/raft"
  ],
  "github.com/coreos/etcd/pkg/netutil": [
    "github.com/coreos/etcd/pkg/cpuutil"
  ],
  "github.com/coreos/etcd/pkg/transport": [
    "github.com/coreos/etcd/pkg/tlsutil"
  ],
  "github.com/coreos/etcd/proxy/grpcproxy": [
    "github.com/coreos/etcd/clientv3/naming",
    "github.com/coreos/etcd/proxy/grpcproxy/cache"
  ],
  "github.com/coreos/etcd/proxy/grpcproxy/adapter": [
    "github.com/coreos/etcd/etcdserver/api/v3lock/v3lockpb",
    "github.com/coreos/etcd/etcdserver/api/v3election/v3electionpb"
  ],
  "github.com/coreos/etcd/proxy/grpcproxy/cache": [
    "github.com/golang/groupcache/lru"
  ],
  "github.com/coreos/etcd/raft": [
    "github.com/coreos/etcd/raft/raftpb"
  ],
  "github.com/coreos/etcd/raft/raftpb": [
    "github.com/golang/protobuf/proto",
    "github.com/gogo/protobuf/gogoproto"
  ],
  "github.com/coreos/etcd/version": [
    "github.com/coreos/go-semver/semver"
  ],
  "github.com/coreos/etcd/wal": [
    "github.com/coreos/etcd/wal/walpb",
    "github.com/coreos/etcd/pkg/crc"
  ],
  "github.com/coreos/go-systemd/journal": [],
  "github.com/coreos/pkg/capnslog": [
    "github.com/coreos/go-systemd/journal"
  ],
  "github.com/cpuguy83/go-md2man/md2man": [],
  "github.com/davecgh/go-spew/spew": [
    "github.com/davecgh/go-spew/spew/testdata"
  ],
  "github.com/dgrijalva/jwt-go": [
    "github.com/dgrijalva/jwt-go/request"
  ],
  "github.com/dgrijalva/jwt-go/request": [
    "github.com/dgrijalva/jwt-go/test"
  ],
  "github.com/dustin/go-humanize": [],
  "github.com/ghodss/yaml": [],
  "github.com/gogo/protobuf/gogoproto": [
    "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
  ],
  "github.com/gogo/protobuf/proto": [
    "github.com/gogo/protobuf/proto/testdata",
    "github.com/gogo/protobuf/proto/proto3_proto"
  ],
  "github.com/gogo/protobuf/proto/proto3_proto": [
    "github.com/gogo/protobuf/types"
  ],
  "github.com/gogo/protobuf/protoc-gen-gogo/descriptor": [
    "github.com/gogo/protobuf/proto"
  ],
  "github.com/gogo/protobuf/types": [
    "github.com/gogo/protobuf/sortkeys"
  ],
  "github.com/golang/groupcache/lru": [],
  "github.com/golang/protobuf/jsonpb": [
    "github.com/golang/protobuf/ptypes/struct",
    "github.com/golang/protobuf/jsonpb/jsonpb_test_proto"
  ],
  "github.com/golang/protobuf/jsonpb/jsonpb_test_proto": [
    "github.com/golang/protobuf/ptypes/wrappers"
  ],
  "github.com/golang/protobuf/proto": [
    "github.com/golang/protobuf/proto/proto3_proto",
    "github.com/golang/protobuf/ptypes"
  ],
  "github.com/golang/protobuf/proto/proto3_proto": [
    "github.com/golang/protobuf/ptypes/any",
    "github.com/golang/protobuf/proto/testdata"
  ],
  "github.com/golang/protobuf/ptypes": [
    "github.com/golang/protobuf/protoc-gen-go/descriptor",
    "github.com/golang/protobuf/ptypes/duration",
    "github.com/golang/protobuf/ptypes/timestamp"
  ],
  "github.com/google/btree": [],
  "github.com/grpc-ecosystem/go-grpc-prometheus": [
    "github.com/prometheus/client_golang/prometheus",
    "github.com/grpc-ecosystem/go-grpc-prometheus/examples/testproto",
    "github.com/stretchr/testify/suite",
    "github.com/grpc-ecosystem/go-grpc-prometheus/examples/grpc-server-with-prometheus/protobuf"
  ],
  "github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb": [
    "github.com/grpc-ecosystem/grpc-gateway/examples/proto/sub",
    "github.com/grpc-ecosystem/grpc-gateway/examples/proto/sub2",
    "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
  ],
  "github.com/grpc-ecosystem/grpc-gateway/runtime": [
    "github.com/grpc-ecosystem/grpc-gateway/utilities",
    "github.com/golang/protobuf/jsonpb",
    "github.com/grpc-ecosystem/grpc-gateway/runtime/internal",
    "github.com/golang/protobuf/ptypes/empty",
    "github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb"
  ],
  "github.com/hashicorp/hcl": [
    "github.com/hashicorp/hcl/hcl/ast",
    "github.com/hashicorp/hcl/hcl/parser",
    "github.com/hashicorp/hcl/json/parser",
    "github.com/hashicorp/hcl/testhelper",
    "github.com/hashicorp/hcl/hcl/printer"
  ],
  "github.com/hashicorp/hcl/hcl/ast": [
    "github.com/hashicorp/hcl/hcl/token"
  ],
  "github.com/hashicorp/hcl/hcl/parser": [
    "github.com/hashicorp/hcl/hcl/scanner"
  ],
  "github.com/hashicorp/hcl/hcl/token": [
    "github.com/hashicorp/hcl/hcl/strconv"
  ],
  "github.com/hashicorp/hcl/json/parser": [
    "github.com/hashicorp/hcl/json/scanner"
  ],
  "github.com/hashicorp/hcl/json/scanner": [
    "github.com/hashicorp/hcl/json/token"
  ],
  "github.com/jonboulle/clockwork": [],
  "github.com/magiconair/properties": [
    "github.com/magiconair/properties/assert"
  ],
  "github.com/matttproud/golang_protobuf_extensions/pbutil": [
    "github.com/matttproud/golang_protobuf_extensions/testdata"
  ],
  "github.com/pelletier/go-toml": [
    "github.com/BurntSushi/toml"
  ],
  "github.com/pkg/sftp": [
    "github.com/kr/fs",
    "github.com/pkg/errors"
  ],
  "github.com/prometheus/client_golang/prometheus": [
    "github.com/prometheus/client_model/go",
    "github.com/prometheus/common/expfmt",
    "github.com/prometheus/procfs",
    "github.com/beorn7/perks/quantile",
    "github.com/prometheus/client_golang/prometheus/promhttp",
    "github.com/prometheus/client_golang/prometheus/push"
  ],
  "github.com/prometheus/common/expfmt": [
    "github.com/matttproud/golang_protobuf_extensions/pbutil",
    "github.com/prometheus/common/internal/bitbucket.org/ww/goautoneg",
    "github.com/prometheus/common/model"
  ],
  "github.com/prometheus/procfs": [
    "github.com/prometheus/procfs/nfs",
    "github.com/prometheus/procfs/xfs",
    "github.com/prometheus/procfs/bcache"
  ],
  "github.com/prometheus/procfs/nfs": [
    "github.com/prometheus/procfs/internal/util"
  ],
  "github.com/scbizu/mew": [
    "github.com/scbizu/mew/cmd"
  ],
  "github.com/scbizu/mew/cmd": [
    "github.com/scbizu/mew/drawer",
    "github.com/scbizu/mew/filter",
    "github.com/scbizu/mew/linker",
    "github.com/spf13/cobra"
  ],
  "github.com/scbizu/mew/drawer": [
    "github.com/awalterschulze/gographviz",
    "github.com/sirupsen/logrus"
  ],
  "github.com/sirupsen/logrus": [
    "github.com/stretchr/testify/assert"
  ],
  "github.com/soheilhy/cmux": [],
  "github.com/spf13/afero": [
    "github.com/spf13/afero/mem",
    "github.com/pkg/sftp"
  ],
  "github.com/spf13/cobra": [
    "github.com/inconshreveable/mousetrap",
    "github.com/spf13/pflag",
    "github.com/spf13/cobra/cobra/cmd",
    "github.com/cpuguy83/go-md2man/md2man",
    "github.com/spf13/cobra/doc"
  ],
  "github.com/spf13/cobra/cobra/cmd": [
    "github.com/mitchellh/go-homedir",
    "github.com/spf13/viper"
  ],
  "github.com/spf13/jwalterweatherman": [
    "github.com/stretchr/testify/require"
  ],
  "github.com/spf13/viper": [
    "github.com/spf13/afero",
    "github.com/spf13/cast",
    "github.com/spf13/jwalterweatherman",
    "github.com/fsnotify/fsnotify",
    "github.com/hashicorp/hcl",
    "github.com/magiconair/properties",
    "github.com/mitchellh/mapstructure",
    "github.com/pelletier/go-toml",
    "github.com/xordataexchange/crypt/config"
  ],
  "github.com/stretchr/testify/assert": [
    "github.com/davecgh/go-spew/spew",
    "github.com/pmezard/go-difflib/difflib"
  ],
  "github.com/tmc/grpc-websocket-proxy/wsproxy": [],
  "github.com/xiang90/probing": [],
  "github.com/xordataexchange/crypt/backend": [
    "github.com/armon/consul-api",
    "github.com/coreos/etcd/client"
  ],
  "github.com/xordataexchange/crypt/backend/mock": [
    "github.com/xordataexchange/crypt/backend"
  ],
  "github.com/xordataexchange/crypt/config": [
    "github.com/xordataexchange/crypt/backend/mock",
    "github.com/xordataexchange/crypt/backend/consul",
    "github.com/xordataexchange/crypt/backend/etcd",
    "github.com/xordataexchange/crypt/encoding/secconf"
  ]
}
```

it will generate the DOT file by [gographviz](https://github.com/awalterschulze/gographviz)

and then generate the PNG file by

```shell
# deep mode dot file
dot -Tpng mew.dot -o mew.png
```

![](mew.png)
