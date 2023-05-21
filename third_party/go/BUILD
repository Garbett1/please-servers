package(default_visibility = ["PUBLIC"])
subinclude("///go//build_defs:go")


go_toolchain(
    name = "toolchain",
    version = "1.20.2",
)

go_system_toolchain(
    name = "system_toolchain",
)

go_module(
    name = "logging",
    module = "gopkg.in/op/go-logging.v1",
    version = "b2cb9fa56473e98db8caba80237377e83fe44db5",
)

go_module(
    name = "terminal",
    install = ["ssh/terminal"],
    module = "golang.org/x/crypto",
    version = "7b85b097bf7527677d54d3220065e966a0e3b613",
)

go_module(
    name = "go-flags",
    module = "github.com/thought-machine/go-flags",
    version = "v1.5.1",
    deps = [
        ":unix",
    ],
)


go_repo(
    name = "net",
    download=":net_dl",
    module = "golang.org/x/net",
    install = ["..."],
)

go_mod_download(
    name = "net_dl",
    module = "golang.org/x/net",
    version = "1bab6f366d9ec8c152f8735b2e359822ef75e3bd",
    deps = [
        ":x_term",
        ":text",
    ],
)

go_module(
    name = "exp",
    install = [
        "constraints",
        "slices",
    ],
    module = "golang.org/x/exp",
    version = "v0.0.0-20230425010034-47ecfdc1ba53",
)

go_module(
    name = "tools",
    install = [
        "cover",
        "go/...",
    ],
    module = "golang.org/x/tools",
    strip = [
        "go/loader/testdata",
        "go/internal/gccgoimporter/testdata",
        "go/gcimporter15/testdata",
    ],
    version = "2ae76fd1560b622911f444c1e66b70a857e1de67",
)

go_module(
    name = "text",
    install = ["..."],
    module = "golang.org/x/text",
    version = "4e4a3210bb54bb31f6ab2cdca2edcc0b50c420c1",
    deps = [
        ":tools",
    ],
)

go_mod_download(
    name = "genproto_download",
    module = "google.golang.org/genproto",
    version = "1856144b1d9c726b541e4a75f9bf7fe259ff92de",
    deps = [":protobuf"],
)

go_repo(
    name = "rpc",
    download = ":genproto_download",
    install = [
        "googleapis/rpc/code",
        "googleapis/rpc/errdetails",
        "googleapis/rpc/status",
        "googleapis/type/date",
    ],
    module = "google.golang.org/genproto",
)

go_module(
    name = "protoc-gen-go-grpc",
    binary = True,
    module = "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
    version = "v1.1.0",
    visibility = ["PUBLIC"],
    deps = [":protobuf"],
)

go_module(
    name = "protoc-gen-go",
    binary = True,
    version = "v1.27.1",
    install = ["cmd/protoc-gen-go"],
    module = "google.golang.org/protobuf",
    deps = [":golang-protobuf"],
    visibility = ["PUBLIC"],
)

go_mod_download(
    name = "protobuf_download",
    module = "github.com/golang/protobuf",
    version = "v1.4.2",
     strip = [
            "proto/proto3_proto",
            "conformance",
        ],
        deps = [
            ":cmp",
            ":golang-protobuf",
        ],
)
go_repo(
name="protobuf",
module="github.com/golang/protobuf",
version="v1.4.2",
download=":protobuf_download",
install=["..."],
)

go_mod_download(
             name = "grpc_dl",
            module = "google.golang.org/grpc",
            version = "v1.48.0",
            deps = [
                    ":cmp",
                    ":oauth2",
                    ":protobuf",
                    ":golang-protobuf",
                    ":rpc",
                    ":unix",
                ],
            strip = [
                    "internal/credentials/xds",
                    "internal/testutils",
                    "internal/xds",
                ],
                )

go_repo(
name="grpc",
module="google.golang.org/grpc",
download=":grpc_dl",
install = [
    ".",
    "attributes",
    "backoff",
    "balancer",
    "balancer/base",
    "balancer/grpclb",
    "balancer/grpclb/...",
    "balancer/roundrobin",
    "binarylog/...",
    "channelz",
    "codes",
    "connectivity",
    "credentials",
    "credentials/alts",
    "credentials/alts/...",
    "credentials/google",
    "credentials/insecure",
    "credentials/oauth",
    "encoding",
    "encoding/gzip",
    "encoding/proto",
    "grpclog",
    "health",
    "health/grpc_health_v1",
    "internal",
    "internal/...",
    "keepalive",
    "metadata",
    "peer",
    "reflection",
    "reflection/...",
    "resolver",
    "resolver/dns",
    "resolver/passthrough",
    "serviceconfig",
    "stats",
    "status",
    "tap",
],
version="v1.48.0"
)


go_module(
    name = "spew",
    install = ["spew"],
    module = "github.com/davecgh/go-spew",
    version = "d8f796af33cc11cb798c1aaeb27a4ebc5099927d",
)

go_module(
    name = "difflib",
    install = ["difflib"],
    module = "github.com/pmezard/go-difflib",
    version = "5d4384ee4fb2527b0a1256a821ebfc92f91efefc",
)

go_module(
    name = "testify",
    install = [
        "assert",
        "require",
    ],
    module = "github.com/stretchr/testify",
    version = "v1.3.0",
    deps = [
        ":difflib",
        ":spew",
    ],
)

go_module(
    name = "grpc-middleware",
    install = [
        "interceptors/recovery",
        "interceptors",
        "metadata",
        "util/backoffutils",
    ],
    module = "github.com/grpc-ecosystem/go-grpc-middleware/v2",
    version = "v2.0.0-rc.5",
    deps = [
        ":grpc",
        ":net",
        ":protobuf","//third_party/go:golang-protobuf",
    ],
)

go_module(
    name = "errgroup",
    install = [
        "errgroup",
        "semaphore",
    ],
    module = "golang.org/x/sync",
    version = "09787c993a3ab68e3d1f5c9b2394ab9433f391be",
    deps = [":net"],
)

go_module(
    name = "unix",
    install = [
        "cpu",
        "execabs",
        "internal/unsafeheader",
        "unix",
    ],
    module = "golang.org/x/sys",
    version = "v0.0.0-20210423082822-04245dca01da",
)

go_module(
    name = "errors",
    module = "github.com/pkg/errors",
    version = "v0.9.1",
)

go_module(
    name = "cli-init",
    install = [
        "flags",
        "logging",
    ],
    module = "github.com/peterebden/go-cli-init/v4",
    version = "v4.0.2",
    deps = [
        ":go-flags",
        ":humanize",
        ":logging",
        ":terminal",
    ],
)

go_module(
    name = "remote-apis",
    install = [
        "build/bazel/...",
    ],
    module = "github.com/bazelbuild/remote-apis",
    version = "35aee1c4a4250d3df846f7ba3e4a4e66cb014ecd",
    deps = [
        ":genproto",
        ":grpc",
        ":protobuf","//third_party/go:golang-protobuf",
        ":rpc",
    ],
)

go_module(
    name = "remote-apis-sdks",
    download = ":remote-apis-sdks_download",
    install = ["go/..."],
    module = "github.com/bazelbuild/remote-apis-sdks",
    deps = [
        ":api",
        ":cmp",
        ":errgroup",
        ":errors",
        ":genproto",
        ":glog",
        ":golang-protobuf",
        ":grpc",
        ":oauth2",
        ":pborman_uuid",
        ":protobuf","//third_party/go:golang-protobuf",
        ":remote-apis",
        ":rpc",
        ":xattr",
        ":zstd",
        ":zstdpool",
    ],
)

go_mod_download(
    name = "remote-apis-sdks_download",
    module = "github.com/peterebden/remote-apis-sdks",
    version = "2a9420921957c3485fc59c6cae6adc849591ba48",
)

go_mod_download(
 name="golang-protobuf_download",
 module = "google.golang.org/protobuf",
 deps=[":cmp"],
 version="v1.30.0",
)

go_repo(
    name = "golang-protobuf",
    install = ["..."],
    module = "google.golang.org/protobuf",
    version = "v1.30.0",
    download=":golang-protobuf_download"
)

go_module(
    name = "cmp",
    install = [
        "cmp",
        "cmp/...",
    ],
    module = "github.com/google/go-cmp",
    version = "b1c9c4891a6525d98001fea424c8926c6d77bb56",
)

go_module(
    name = "glog",
    download = ":glog_download",
    module = "github.com/golang/glog",
)

go_mod_download(
    name = "glog_download",
    module = "github.com/peterebden/glog",
    version = "78e7b285a1f26e8ca142f6f7b584511bc551988b",
)

go_module(
    name = "oauth2",
    install = [
        ".",
        "google",
        "internal",
        "jws",
        "jwt",
        "authhandler",
        "google/internal/externalaccount",
    ],
    module = "golang.org/x/oauth2",
    version = "fd043fe589d2d1486b6af56f44a691e819752a23",
    deps = [
        ":cloud-metadata",
        ":net",
    ],
)

go_module(
    name = "genproto",
    download = ":genproto_download",
    install = [
        "googleapis/api/annotations",
        "googleapis/longrunning",
        "googleapis/bytestream",
        "googleapis/iam/v1",
        "googleapis/iam/credentials/v1",
        "googleapis/pubsub/v1",
        "protobuf/field_mask",
        "googleapis/type/expr",
    ],
    module = "google.golang.org/genproto",
    deps = [
        ":grpc",
        ":protobuf","//third_party/go:golang-protobuf",
        ":rpc",
    ],
)

go_module(
    name = "uuid",
    module = "github.com/google/uuid",
    version = "v1.1.1",
)

go_module(
    name = "pborman_uuid",
    module = "github.com/pborman/uuid",
    version = "v1.2.0",
    deps = [":uuid"],
)

go_module(
    name = "ristretto",
    install = [
        ".",
        "...",
    ],
    module = "github.com/dgraph-io/ristretto",
    version = "v0.0.3",
    deps = [":xxhash"],
)

go_module(
    name = "xxhash",
    module = "github.com/cespare/xxhash",
    version = "v1.1.0",
)

go_module(
    name = "humanize",
    module = "github.com/dustin/go-humanize",
    version = "v1.0.0",
)

go_module(
    name = "retryablehttp",
    module = "github.com/hashicorp/go-retryablehttp",
    version = "v0.6.1",
    deps = [
        ":cleanhttp",
        ":hclog",
    ],
)

go_module(
    name = "sri",
    module = "github.com/peterebden/go-sri",
    version = "v1.1.1",
)

go_module(
    name = "go-multierror",
    module = "github.com/hashicorp/go-multierror",
    version = "v1.1.1",
    deps = [":errwrap"],
)

go_module(
    name = "errwrap",
    module = "github.com/hashicorp/errwrap",
    version = "v1.0.0",
)

go_module(
    name = "cleanhttp",
    module = "github.com/hashicorp/go-cleanhttp",
    version = "v0.5.1",
)

go_module(
    name = "hclog",
    module = "github.com/hashicorp/go-hclog",
    version = "v0.12.0",
    deps = [
        ":color",
        ":isatty",
    ],
)

go_module(
    name = "isatty",
    module = "github.com/mattn/go-isatty",
    version = "v0.0.12",
    deps = [":unix"],
)

go_module(
    name = "colorable",
    module = "github.com/mattn/go-colorable",
    version = "v0.1.8",
    deps = [
        ":isatty",
    ],
)

go_module(
    name = "color",
    module = "github.com/fatih/color",
    version = "v1.9.0",
    deps = [
        ":colorable",
        ":isatty",
        ":unix",
    ],
)

go_module(
    name = "copyfile",
    module = "github.com/peterebden/go-copyfile",
    version = "bc0baf74909cb55565cf9536cb003f04da55372d",
)

go_module(
    name = "godirwalk",
    module = "github.com/karrick/godirwalk",
    strip = [
        "examples",
        "testdata",
        "vendor",
    ],
    version = "v1.7.0",
    deps = [":errors"],
)

go_module(
    name = "http-admin",
    module = "github.com/thought-machine/http-admin",
    version = "v1.1.0",
    deps = [
        ":logging",
        ":mux",
        ":net",
        "//third_party/go/prometheus",
    ],
)

go_module(
    name = "mux",
    module = "github.com/gorilla/mux",
    version = "v1.7.4",
)

go_module(
    name = "xattr",
    module = "github.com/pkg/xattr",
    version = "v0.4.9",
    deps = [":unix"],
)

go_module(
    name = "psutil",
    install = [
        "internal/common",
        "mem",
    ],
    module = "github.com/shirou/gopsutil",
    version = "v2.17.09",
    deps = [":unix"],
)

go_module(
    name = "zstd",
    install = [
        ".",
        "fse",
        "huff0",
        "snappy",
        "zstd",
        "zstd/...",
    ],
    module = "github.com/klauspost/compress",
    version = "v1.11.0",
)

go_module(
    name = "go-grpc-compression",
    install = ["zstd"],
    module = "github.com/mostynb/go-grpc-compression",
    version = "v1.1.2",
    deps = [
        ":grpc",
        ":zstd",
    ],
)

go_mod_download(
    name = "cloud_download",
    module = "cloud.google.com/go",
    version = "v0.78.0",
)

go_module(
    name = "cloud-metadata",
    download = ":cloud_download",
    install = ["compute/metadata"],
    module = "cloud.google.com/go",
)

go_module(
    name = "google-cloud-go",
    download = ":cloud_download",
    install = [
        "iam",
        "iam/credentials/...",
        "internal",
        "internal/optional",
        "internal/pubsub",
        "internal/trace",
        "internal/version",
    ],
    module = "cloud.google.com/go",
    deps = [
        ":api",
        ":errgroup",
        ":gax-go",
        ":genproto",
        ":go.opencensus.io",
    ],
)

go_module(
    name = "cloud_pubsub",
    install = [
        ".",
        "apiv1",
    ],
    module = "cloud.google.com/go/pubsub",
    version = "v1.10.0",
    deps = [
        ":api",
        ":gax-go",
        ":genproto",
        ":go.opencensus.io",
        ":google-cloud-go",
    ],
)

go_module(
    name = "cloud-storage",
    install = ["..."],
    module = "cloud.google.com/go/storage",
    version = "v1.25.0",
    deps = [
        ":api",
        ":errgroup",
        ":gax-go",
        ":genproto",
        ":go.opencensus.io",
        ":google-cloud-go",
    ],
)

go_module(
    name = "gax-go",
    module = "github.com/googleapis/gax-go/v2",
    version = "v2.0.5",
    deps = [
        ":grpc",
    ],
)

go_module(
    name = "go.opencensus.io",
    install = [
        ".",
        "...",
    ],
    module = "go.opencensus.io",
    version = "v0.22.1",
    deps = [
        ":groupcache",
        ":grpc",
        ":net",
    ],
)

go_module(
    name = "groupcache",
    install = [
        "lru",
    ],
    module = "github.com/golang/groupcache",
    version = "404acd9df4cc9859d64fb9eed42e5c026187287a",
)

go_module(
    name = "api",
    install = [
        "calendar/v3",
        "googleapi",
        "googleapi/transport",
        "iamcredentials/v1",
        "idtoken",
        "internal",
        "internal/gensupport",
        "internal/impersonate",
        "internal/third_party/uritemplates",
        "iterator",
        "option",
        "option/internaloption",
        "storage/v1",
        "sheets/v4",
        "support/bundler",
        "transport",
        "transport/bytestream",
        "transport/cert",
        "transport/grpc",
        "transport/http",
        "transport/http/internal/propagation",
        "transport/internal/dca",
        "dataflow/...",
        "compute/v1",
        "container/v1",
        "drive/v3",
    ],
    module = "google.golang.org/api",
    version = "v0.92.0",
    deps = [
        ":errgroup",
        ":gax-go",
        ":genproto",
        ":go.opencensus.io",
        ":grpc",
        ":oauth2",
        ":uuid",
        ":enterprise-certificate-proxy",
        ":net",
    ],
)

go_module(
    name = "redis",
    install = [
        "...",
    ],
    module = "github.com/go-redis/redis/v8",
    version = "v8.3.3",
    deps = [
        ":opentelemetry",
        ":rendezvous",
        ":xxhash_v2",
    ],
)

go_module(
    name = "opentelemetry",
    install = [
        ".",
        "api/global",
        "api/global/internal",
        "api/metric",
        "api/metric/registry",
        "api/trace",
        "codes",
        "internal/...",
        "label",
        "unit",
    ],
    licences = ["apache-2.0"],
    module = "go.opentelemetry.io/otel",
    version = "v0.13.0",
)

go_module(
    name = "xxhash_v2",
    module = "github.com/cespare/xxhash/v2",
    version = "v2.1.1",
)

go_module(
    name = "rendezvous",
    licences = ["mit"],
    module = "github.com/dgryski/go-rendezvous",
    version = "9f7001d12a5f0021fd3283525f888b5814ccee27",
)

go_module(
    name = "automaxprocs",
    install = [
        "...",
    ],
    licences = ["MIT"],
    module = "go.uber.org/automaxprocs",
    version = "v1.4.0",
)

go_module(
    name = "json-iterator",
    install = ["..."],
    licences = ["MIT"],
    module = "github.com/json-iterator/go",
    version = "v1.1.6",
    deps = [
        ":concurrent",
        ":reflect2",
    ],
)

go_module(
    name = "concurrent",
    install = ["..."],
    licences = ["Apache-2.0"],
    module = "github.com/modern-go/concurrent",
    version = "1.0.3",
)

go_module(
    name = "reflect2",
    install = ["..."],
    licences = ["Apache-2.0"],
    module = "github.com/modern-go/reflect2",
    version = "v1.0.1",
    deps = [":concurrent"],
)

go_module(
    name = "kingpin",
    licences = ["MIT"],
    module = "gopkg.in/alecthomas/kingpin.v2",
    version = "v2.2.6",
    deps = [
        ":template",
        ":units",
    ],
)

go_module(
    name = "template",
    install = [
        "",
        "parse",
    ],
    licences = ["BSD-3-Clause"],
    module = "github.com/alecthomas/template",
    version = "a0175ee3bccc567396460bf5acd36800cb10c49c",
)

go_module(
    name = "units",
    licences = ["MIT"],
    module = "github.com/alecthomas/units",
    version = "2efee857e7cfd4f3d0138cc3cbb1b4966962b93a",
)

go_module(
    name = "logrus",
    licences = ["MIT"],
    module = "github.com/sirupsen/logrus",
    version = "v1.6.0",
    deps = [
        ":terminal",
        ":unix",
    ],
)

go_module(
    name = "x_term",
    install = ["..."],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/term",
    version = "v0.0.0-20210615171337-6886f2dfbf5b",
    deps = [":unix"],
)

go_module(
    name = "zstdpool",
    licences = ["Apache-2.0"],
    module = "github.com/mostynb/zstdpool-syncpool",
    version = "v0.0.7",
    visibility = ["PUBLIC"],
    deps = [":zstd"],
)

go_module(
    name = "enterprise-certificate-proxy",
    install = ["..."],
    licences = ["Apache-2.0"],
    module = "github.com/googleapis/enterprise-certificate-proxy",
    version = "v0.1.0",
)

go_repo(
name="otelhttp",
install=[".", "..."],
module="go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp", version="v0.41.1")

go_repo(name="otel-config-go", module="github.com/honeycombio/otel-config-go", install=["...",], version="v1.9.0")
go_repo(module="google.golang.org/appengine", version="v1.6.7")
go_repo(module="cloud.google.com/go/scheduler", version="v1.7.0")
go_repo(module="github.com/google/martian/v3", version="v3.2.1")
go_repo(module="cloud.google.com/go/recaptchaenterprise/v2", version="v2.5.0")
go_repo(module="cloud.google.com/go/datastore", version="v1.10.0")
go_repo(module="cloud.google.com/go/managedidentities", version="v1.4.0")
go_repo(module="go.opentelemetry.io/otel", version="v1.15.1")
go_repo(module="go.opentelemetry.io/contrib/instrumentation/runtime", version="v0.41.1")
go_repo(module="golang.org/x/tools", version="v0.6.0")
go_repo(module="go.opentelemetry.io/otel/exporters/stdout/stdouttrace", version="v1.15.1")
go_repo(module="github.com/shoenig/go-m1cpu", version="v0.1.4")
go_repo(module="github.com/tklauser/go-sysconf", version="v0.3.11")
go_repo(module="github.com/iancoleman/strcase", version="v0.2.0")
go_repo(module="github.com/pkg/errors", version="v0.9.1")
go_repo(module="cloud.google.com/go/compute/metadata", version="v0.2.3")
go_repo(module="github.com/stretchr/testify", version="v1.8.2")
go_repo(module="github.com/googleapis/gax-go/v2", version="v2.7.0")
go_repo(module="github.com/envoyproxy/protoc-gen-validate", version="v0.9.1")
go_repo(module="cloud.google.com/go/certificatemanager", version="v1.4.0")
go_repo(module="cloud.google.com/go/filestore", version="v1.4.0")
go_repo(module="cloud.google.com/go/phishingprotection", version="v0.6.0")
go_repo(module="go.opentelemetry.io/contrib/propagators/ot", version="v1.16.1")
go_repo(module="cloud.google.com/go/datafusion", version="v1.5.0")
go_repo(module="gopkg.in/errgo.v2", version="v2.1.0")
go_repo(module="cloud.google.com/go/servicemanagement", version="v1.5.0")
go_repo(module="github.com/go-logr/logr", version="v1.2.4")
go_repo(module="github.com/shoenig/test", version="v0.6.3")
go_repo(module="golang.org/x/mod", version="v0.8.0")
go_repo(module="github.com/spf13/afero", version="v1.9.2")
go_repo(module="github.com/golang/snappy", version="v0.0.3")
go_repo(module="cloud.google.com/go/area120", version="v0.6.0")
go_repo(module="cloud.google.com/go/orgpolicy", version="v1.5.0")
go_repo(module="cloud.google.com/go/bigquery", version="v1.44.0")
go_repo(module="github.com/spaolacci/murmur3", version="v0.0.0-20180118202830-f09979ecbc72")
go_repo(module="cloud.google.com/go/dataflow", version="v0.7.0")
go_repo(module="cloud.google.com/go/trace", version="v1.4.0")
go_repo(module="cloud.google.com/go/websecurityscanner", version="v1.4.0")
go_repo(module="github.com/go-logr/stdr", version="v1.2.2")
go_repo(module="github.com/kr/pretty", version="v0.2.1")
go_repo(module="cloud.google.com/go/ids", version="v1.2.0")
go_repo(module="cloud.google.com/go/pubsublite", version="v1.5.0")
go_repo(module="cloud.google.com/go/vmwareengine", version="v0.1.0")
go_repo(module="go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp", version="v0.38.1")
go_repo(module="github.com/ghodss/yaml", version="v1.0.0")
go_repo(module="github.com/chzyer/logex", version="v1.1.10")
go_repo(module="github.com/ianlancetaylor/demangle", version="v0.0.0-20181102032728-5e5cf60278f6")
go_repo(module="cloud.google.com/go/aiplatform", version="v1.27.0")
go_repo(module="github.com/cespare/xxhash", version="v1.1.0")
go_repo(module="cloud.google.com/go/workflows", version="v1.9.0")
go_repo(module="cloud.google.com/go/dataproc", version="v1.8.0")
go_repo(module="cloud.google.com/go/videointelligence", version="v1.9.0")
go_repo(module="github.com/sethvargo/go-envconfig", version="v0.9.0")
go_repo(module="cloud.google.com/go/edgecontainer", version="v0.2.0")
go_repo(module="github.com/tklauser/numcpus", version="v0.6.0")
go_repo(module="github.com/chzyer/test", version="v0.0.0-20180213035817-a1ea475d72b1")
go_repo(module="cloud.google.com/go/contactcenterinsights", version="v1.4.0")
go_repo(module="cloud.google.com/go/servicedirectory", version="v1.7.0")
go_repo(module="cloud.google.com/go/apigateway", version="v1.4.0")
go_repo(module="cloud.google.com/go/appengine", version="v1.5.0")
go_repo(module="cloud.google.com/go/assuredworkloads", version="v1.9.0")
go_repo(module="cloud.google.com/go/retail", version="v1.11.0")
go_repo(module="go.opentelemetry.io/otel/exporters/otlp/internal/retry", version="v1.15.1")
go_repo(module="cloud.google.com/go/essentialcontacts", version="v1.4.0")
go_repo(module="cloud.google.com/go/gkemulticloud", version="v0.4.0")
go_repo(module="golang.org/x/xerrors", version="v0.0.0-20220907171357-04be3eba64a2")
go_repo(module="github.com/kr/fs", version="v0.1.0")
go_repo(module="cloud.google.com/go/datalabeling", version="v0.6.0")
go_repo(module="cloud.google.com/go/dlp", version="v1.7.0")
go_repo(module="cloud.google.com/go/monitoring", version="v1.8.0")
go_repo(module="golang.org/x/sys", version="v0.7.0")
go_repo(module="google.golang.org/grpc/cmd/protoc-gen-go-grpc", version="v1.1.0")
go_repo(module="cloud.google.com/go/security", version="v1.10.0")
go_repo(module="cloud.google.com/go/dataqna", version="v0.6.0")
go_repo(module="cloud.google.com/go/gkebackup", version="v0.3.0")
go_repo(module="cloud.google.com/go/tpu", version="v1.4.0")
go_repo(module="cloud.google.com/go/gkehub", version="v0.10.0")
go_repo(
name="honey_otel",
install=[".", "..."],
module="github.com/honeycombio/honeycomb-opentelemetry-go", version="v0.6.0")
go_repo(module="github.com/felixge/httpsnoop", version="v1.0.3")
go_repo(module="github.com/lyft/protoc-gen-star", version="v0.6.1")
go_repo(module="cloud.google.com/go/mediatranslation", version="v0.6.0")
go_repo(module="cloud.google.com/go/networkmanagement", version="v1.5.0")
go_repo(module="cloud.google.com/go/webrisk", version="v1.7.0")
go_repo(module="github.com/golang/glog", version="v1.0.0")
go_repo(module="cloud.google.com/go/iap", version="v1.5.0")
go_repo(module="cloud.google.com/go/metastore", version="v1.8.0")
go_repo(module="cloud.google.com/go/networkconnectivity", version="v1.7.0")
go_repo(module="go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc", version="v0.38.1")
go_repo(module="cloud.google.com/go/pubsub", version="v1.27.1")
go_repo(module="cloud.google.com/go/baremetalsolution", version="v0.4.0")
go_repo(module="cloud.google.com/go/vmmigration", version="v1.3.0")


go_repo(module="cloud.google.com/go/analytics", version="v0.12.0")
go_repo(module="cloud.google.com/go/container", version="v1.7.0")
go_repo(module="cloud.google.com/go/memcache", version="v1.7.0")
go_repo(module="github.com/googleapis/google-cloud-go-testing", version="v0.0.0-20200911160855-bcd43fbb19e8")
go_repo(module="cloud.google.com/go/maps", version="v0.1.0")
go_repo(module="cloud.google.com/go/domains", version="v0.7.0")
go_repo(module="github.com/client9/misspell", version="v0.3.4")
go_repo(module="cloud.google.com/go/iam", version="v0.8.0")


go_repo(module="github.com/kr/text", version="v0.1.0")
go_repo(module="cloud.google.com/go/channel", version="v1.9.0")
go_repo(module="cloud.google.com/go/gaming", version="v1.8.0")
go_repo(module="github.com/cncf/udpa/go", version="v0.0.0-20220112060539-c52dc94e7fbe")
go_repo(module="go.opentelemetry.io/proto/otlp", version="v0.19.0")
go_repo(module="cloud.google.com/go/accesscontextmanager", version="v1.4.0")
go_repo(module="golang.org/x/sync", version="v0.1.0")
go_repo(module="github.com/envoyproxy/go-control-plane", version="v0.10.3")
go_repo(module="github.com/kisielk/gotool", version="v1.0.0")
go_repo(module="cloud.google.com/go/eventarc", version="v1.8.0")
go_repo(module="cloud.google.com/go/translate", version="v1.4.0")
go_repo(module="cloud.google.com/go/vpcaccess", version="v1.5.0")
go_repo(module="github.com/rogpeppe/go-internal", version="v1.3.0")
go_repo(module="cloud.google.com/go/beyondcorp", version="v0.3.0")
go_repo(module="cloud.google.com/go/datastream", version="v1.5.0")
go_repo(module="cloud.google.com/go/deploy", version="v1.5.0")
go_repo(module="cloud.google.com/go/functions", version="v1.9.0")
go_repo(module="cloud.google.com/go/texttospeech", version="v1.5.0")
go_repo(module="go.uber.org/goleak", version="v1.2.1")
go_repo(module="github.com/chzyer/readline", version="v0.0.0-20180603132655-2972be24d48e")
go_repo(module="cloud.google.com/go/kms", version="v1.6.0")
go_repo(module="cloud.google.com/go/securitycenter", version="v1.16.0")
go_repo(module="go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp", version="v1.15.1")
go_repo(module="go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc", version="v0.41.1")
go_repo(module="honnef.co/go/tools", version="v0.0.1-2020.1.4")
go_repo(module="cloud.google.com/go/cloudbuild", version="v1.4.0")
go_repo(module="cloud.google.com/go/notebooks", version="v1.5.0")
go_repo(module="cloud.google.com/go/oslogin", version="v1.7.0")
go_repo(module="github.com/shirou/gopsutil/v3", version="v3.23.3")
go_repo(module="github.com/grpc-ecosystem/grpc-gateway/v2", version="v2.12.0")

go_repo(module="github.com/census-instrumentation/opencensus-proto", version="v0.4.1")
go_repo(module="cloud.google.com/go/networksecurity", version="v0.6.0")
go_repo(module="cloud.google.com/go/video", version="v1.9.0")
go_repo(module="github.com/BurntSushi/toml", version="v0.3.1")
go_repo(module="github.com/google/go-cmp", version="v0.5.9")
go_repo(module="go.opentelemetry.io/otel/trace", version="v1.15.1")
go_repo(module="cloud.google.com/go/orchestration", version="v1.4.0")
go_repo(module="cloud.google.com/go/speech", version="v1.9.0")
go_repo(module="gopkg.in/check.v1", version="v1.0.0-20190902080502-41f04d3bba15")
go_repo(module="cloud.google.com/go/recommendationengine", version="v0.6.0")
go_repo(module="go.opentelemetry.io/contrib/instrumentation/host", version="v0.41.1")
go_repo(module="github.com/kr/pty", version="v1.1.1")
go_repo(module="github.com/googleapis/enterprise-certificate-proxy", version="v0.2.0")
go_repo(module="cloud.google.com/go/clouddms", version="v1.4.0")
go_repo(module="github.com/golang/groupcache", version="v0.0.0-20200121045136-8c9f03a8e57e")
go_repo(module="cloud.google.com/go/redis", version="v1.10.0")
go_repo(module="cloud.google.com/go/servicecontrol", version="v1.5.0")
go_repo(module="go.opentelemetry.io/otel/exporters/otlp/otlptrace", version="v1.15.1")

go_repo(module="github.com/OneOfOne/xxhash", version="v1.2.2")
go_repo(module="cloud.google.com/go/containeranalysis", version="v0.6.0")
go_repo(module="cloud.google.com/go/gkeconnect", version="v0.6.0")
go_repo(module="github.com/cenkalti/backoff/v4", version="v4.2.1")
go_repo(module="github.com/antihax/optional", version="v1.0.0")
go_repo(module="github.com/golang/mock", version="v1.4.4")
go_repo(module="github.com/google/pprof", version="v0.0.0-20200708004538-1a94d8640e99")
go_repo(module="cloud.google.com/go/vision", version="v1.2.0")
go_repo(module="github.com/davecgh/go-spew", version="v1.1.1")
go_repo(module="go.opentelemetry.io/contrib/propagators/b3", version="v1.16.1")
go_repo(module="github.com/jstemmer/go-junit-report", version="v0.9.1")
go_repo(module="cloud.google.com/go/serviceusage", version="v1.4.0")
go_repo(module="cloud.google.com/go/talent", version="v1.4.0")
go_repo(module="golang.org/x/oauth2", version="v0.4.0")
go_repo(module="golang.org/x/term", version="v0.7.0")
go_repo(module="cloud.google.com/go/recommender", version="v1.8.0")
go_repo(module="github.com/pmezard/go-difflib", version="v1.0.0")
go_repo(module="cloud.google.com/go/storage", version="v1.27.0")
go_repo(module="cloud.google.com/go/automl", version="v1.8.0")
go_repo(module="cloud.google.com/go/batch", version="v0.4.0")
go_repo(module="cloud.google.com/go/privatecatalog", version="v0.6.0")
go_repo(module="github.com/go-ole/go-ole", version="v1.2.6")
go_repo(module="github.com/cespare/xxhash/v2", version="v2.2.0")
go_repo(module="cloud.google.com/go/dataform", version="v0.5.0")
go_repo(module="cloud.google.com/go/resourcesettings", version="v1.4.0")
go_repo(module="cloud.google.com/go/secretmanager", version="v1.9.0")
go_repo(module="github.com/google/renameio", version="v0.1.0")
go_repo(module="cloud.google.com/go/optimization", version="v1.2.0")
go_repo(module="go.opentelemetry.io/otel/exporters/otlp/otlpmetric", version="v0.38.1")
go_repo(module="cloud.google.com/go/accessapproval", version="v1.5.0")
go_repo(module="github.com/prometheus/client_model", version="v0.2.0")
go_repo(module="cloud.google.com/go/artifactregistry", version="v1.9.0")
go_repo(module="cloud.google.com/go/asset", version="v1.10.0")
go_repo(module="cloud.google.com/go/dataplex", version="v1.4.0")
go_repo(module="cloud.google.com/go/documentai", version="v1.10.0")
go_repo(module="cloud.google.com/go/shell", version="v1.4.0")
go_repo(module="go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc", version="v1.15.1")
go_repo(module="github.com/lufia/plan9stats", version="v0.0.0-20220913051719-115f729f3c8c")
go_repo(module="github.com/rogpeppe/fastuuid", version="v1.2.0")
go_repo(module="cloud.google.com/go/billing", version="v1.7.0")
go_repo(module="cloud.google.com/go/vision/v2", version="v2.5.0")
go_repo(module="github.com/stretchr/objx", version="v0.5.0")
go_repo(module="github.com/pkg/sftp", version="v1.13.1")
go_repo(module="cloud.google.com/go/osconfig", version="v1.10.0")
go_repo(module="cloud.google.com/go/spanner", version="v1.41.0")
go_repo(module="cloud.google.com/go", version="v0.107.0")
go_repo(module="golang.org/x/lint", version="v0.0.0-20210508222113-6edffad5e616")
go_repo(module="github.com/google/uuid", version="v1.3.0")
go_repo(module="cloud.google.com/go/datacatalog", version="v1.8.0")
go_repo(module="cloud.google.com/go/cloudtasks", version="v1.8.0")
go_repo(module="cloud.google.com/go/firestore", version="v1.9.0")
go_repo(module="cloud.google.com/go/resourcemanager", version="v1.4.0")
go_repo(module="gopkg.in/yaml.v3", version="v3.0.1")
go_repo(module="github.com/yusufpapurcu/wmi", version="v1.2.2")
go_repo(module="golang.org/x/crypto", version="v0.0.0-20211108221036-ceb1ce70b4fa")
go_repo(module="github.com/yuin/goldmark", version="v1.4.13")
go_repo(module="google.golang.org/api", version="v0.103.0")
go_repo(module="go.uber.org/multierr", version="v1.11.0")
go_repo(module="golang.org/x/exp", version="v0.0.0-20190121172915-509febef88a4")
go_repo(module="cloud.google.com/go/grafeas", version="v0.2.0")
go_repo(module="cloud.google.com/go/errorreporting", version="v0.3.0")
go_repo(module="cloud.google.com/go/lifesciences", version="v0.6.0")
go_repo(module="go.opentelemetry.io/otel/metric", version="v0.38.1")
go_repo(module="cloud.google.com/go/binaryauthorization", version="v1.4.0")
go_repo(module="cloud.google.com/go/recaptchaenterprise", version="v1.3.1")
go_repo(module="cloud.google.com/go/storagetransfer", version="v1.6.0")
go_repo(module="github.com/power-devops/perfstat", version="v0.0.0-20220216144756-c35f1ee13d7c")
go_repo(module="github.com/cncf/xds/go", version="v0.0.0-20230105202645-06c439db220b")
go_repo(module="cloud.google.com/go/apigeeconnect", version="v1.4.0")
go_repo(module="cloud.google.com/go/dialogflow", version="v1.19.0")
go_repo(module="cloud.google.com/go/language", version="v1.8.0")
go_repo(module="cloud.google.com/go/policytroubleshooter", version="v1.4.0")
go_repo(module="go.opentelemetry.io/otel/sdk", version="v1.15.1")
go_repo(module="go.opentelemetry.io/otel/sdk/metric", version="v0.38.1")
go_repo(module="cloud.google.com/go/compute", version="v1.15.1")
go_repo(module="cloud.google.com/go/longrunning", version="v0.3.0")
go_repo(module="cloud.google.com/go/gsuiteaddons", version="v1.4.0")
go_repo(module="cloud.google.com/go/iot", version="v1.4.0")
go_repo(module="cloud.google.com/go/logging", version="v1.6.1")
go_repo(module="cloud.google.com/go/run", version="v0.3.0")
