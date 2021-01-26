load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "rules_proto",
    sha256 = "602e7161d9195e50246177e7c55b2f39950a9cf7366f74ed5f22fd45750cd208",
    strip_prefix = "rules_proto-97d8af4dc474595af3900dd85cb3a29ad28cc313",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
    ],
)
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
rules_proto_dependencies()
rules_proto_toolchains()

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "7904dbecbaffd068651916dce77ff3437679f9d20e1a7956bff43826e7645fcc",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.25.1/rules_go-v0.25.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.25.1/rules_go-v0.25.1.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

http_archive(
    name = "bazel_gazelle",
    sha256 = "222e49f034ca7a1d1231422cdb67066b885819885c356673cb1f72f748a3c9d4",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.3/bazel-gazelle-v0.22.3.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.3/bazel-gazelle-v0.22.3.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    sum = "h1:Iwh0ba2kTgq2Q6mJiXhzrrjD7h11nEVnbMHFmp0/HsQ=",
    version = "v0.0.0-20210122163508-8081c04a3579",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:TwIQcH3es+MojMVojxxfQ3l3OF2KzlRxML2xZq0kRo8=",
    version = "v1.35.0",
)

go_repository(
    name = "org_golang_google_protobuf",
    importpath = "google.golang.org/protobuf",
    sum = "h1:Ejskq+SyPohKW+1uil0JJMtmHCgJPJ/qWTxr8qp+R4c=",
    version = "v1.25.0",
)

gazelle_dependencies(go_sdk = "go_sdk")

#http_archive(
#    name = "com_google_protobuf",
#    sha256 = "9748c0d90e54ea09e5e75fb7fac16edce15d2028d4356f32211cfa3c0e956564",
#    strip_prefix = "protobuf-3.11.4",
#    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.11.4.zip"],
#)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:oWX7TPOiFAMXLq8o0ikBYfCJVlRHBcsciT5bXOrH628=",
    version = "v0.0.0-20190311183353-d8887717615a",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:g61tztE5qeGQ89tm6NTjjM9VPIm088od1l6aSorWRWg=",
    version = "v0.3.0",
)

#load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
#protobuf_deps()

http_archive(
    name = "build_bazel_rules_nodejs",
    sha256 = "84b1d11b1f3bda68c24d992dc6e830bca9db8fa12276f2ca7fcb7761c893976b",
    urls = ["https://github.com/bazelbuild/rules_nodejs/releases/download/3.0.0-rc.1/rules_nodejs-3.0.0-rc.1.tar.gz"],
)

#load("@build_bazel_rules_nodejs//:index.bzl", "node_repositories", "yarn_install")

#node_repositories(
#    package_json = ["//:package.json"],
#    node_version = "8.11.1",
#    yarn_version = "1.5.1",
#)
#yarn_install(
    # Name this npm so that Bazel Label references look like @npm//package
#    name = "npm",
#    package_json = "//:package.json",
#    yarn_lock = "//:yarn.lock",
#    symlink_node_modules = True,  # Expose installed packages for the IDE and the developer. See managed_directories.
#)

go_register_toolchains(version = "1.15.7")
