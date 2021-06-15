load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "rules_proto",
    sha256 = "8e7d59a5b12b233be5652e3d29f42fba01c7cbab09f6b3a8d0a57ed6d1e9a0da",
    strip_prefix = "rules_proto-7e4afce6fe62dbff0a4a03450143146f9f2d7488",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/7e4afce6fe62dbff0a4a03450143146f9f2d7488.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/7e4afce6fe62dbff0a4a03450143146f9f2d7488.tar.gz",
    ],
)
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
rules_proto_dependencies()
rules_proto_toolchains()

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "69de5c704a05ff37862f7e0f5534d4f479418afc21806c887db544a316f3cb6b",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.27.0/rules_go-v0.27.0.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.27.0/rules_go-v0.27.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
go_rules_dependencies()
go_register_toolchains(version = "1.16")

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

http_archive(
    name = "build_bazel_rules_nodejs",
    sha256 = "4a5d654a4ccd4a4c24eca5d319d85a88a650edf119601550c95bf400c8cc897e",
    urls = ["https://github.com/bazelbuild/rules_nodejs/releases/download/3.5.1/rules_nodejs-3.5.1.tar.gz"],
)

load("@build_bazel_rules_nodejs//:index.bzl", "node_repositories", "yarn_install")
load("@build_bazel_rules_nodejs//internal/js_library:js_library.bzl", "js_library")

node_repositories()
yarn_install(
    # Name this npm so that Bazel Label references look like @npm//package
    name = "npm",
    package_json = "//:package.json",
    yarn_lock = "//:yarn.lock",
    frozen_lockfile = False,
    symlink_node_modules = True,  # Expose installed packages for the IDE and the developer. See managed_directories.
)
#load("@npm//@bazel/labs:package.bzl", "npm_bazel_labs_dependencies")
#npm_bazel_labs_dependencies()

#rules_kotlin_version = "legacy-1.3.0"
#rules_kotlin_sha = "4fd769fb0db5d3c6240df8a9500515775101964eebdf85a3f9f0511130885fde"
#http_archive(
#    name = "io_bazel_rules_kotlin",
#    urls = ["https://github.com/bazelbuild/rules_kotlin/archive/%s.zip" % rules_kotlin_version],
#    type = "zip",
#    strip_prefix = "rules_kotlin-%s" % rules_kotlin_version,
#    sha256 = rules_kotlin_sha,
#)

#load("@io_bazel_rules_kotlin//kotlin:kotlin.bzl", "kotlin_repositories", "kt_register_toolchains")
#kotlin_repositories()
#kt_register_toolchains()

#http_archive(
#    name = "com_github_grpc_grpc_kotlin",
#    sha256 = "6c9642710cc4ca0ebe18732c12880c4fc112253aa38c9d5855249f98fb0d4d02",
#    urls = ["https://github.com/grpc/grpc-kotlin/archive/v1.0.0.zip"],
#)
#load("@com_github_grpc_grpc_kotlin//:repositories.bzl", "grpc_kt_repositories")
#grpc_kt_repositories()

#git_repository(
#    name = "com_github_grpc_grpc_kotlin",
#    branch = "master",
#    remote = "https://github.com/grpc/grpc-kotlin.git",
#)
#load("@com_github_grpc_grpc_kotlin//:repositories.bzl", "grpc_kt_repositories", "IO_GRPC_GRPC_KOTLIN_ARTIFACTS", "IO_GRPC_GRPC_KOTLIN_OVERRIDE_TARGETS")
#grpc_kt_repositories()

#http_archive(
#    name = "rules_jvm_external",
#    sha256 = "62133c125bf4109dfd9d2af64830208356ce4ef8b165a6ef15bbff7460b35c3a",
#    strip_prefix = "rules_jvm_external-3.0",
#    url = "https://github.com/bazelbuild/rules_jvm_external/archive/3.0.zip",
#)

#SPRING_FRAMEWORK_VERSION = "5.2.3.RELEASE"
#SPRING_BOOT_VERSION = "2.2.4.RELEASE"
#SPRING_CLOUD_CONSUL_VERSION = "2.2.1.RELEASE"
#GRPC_JAVA_VERSION = "1.27.2"
#GRPC_SPRING_BOOT_VERSION = "3.5.2"

#http_archive(
#    name = "io_grpc_grpc_java",
#    sha256 = "92ffb4391f847e02e115933a761e243dd1423f3fcafdc9b7ae0327eca102d76b",
#    strip_prefix = "grpc-java-%s" % GRPC_JAVA_VERSION,
#    url = "https://github.com/grpc/grpc-java/archive/v%s.zip" % GRPC_JAVA_VERSION,
#)
#load("@rules_jvm_external//:defs.bzl", "maven_install")
#load("@io_grpc_grpc_java//:repositories.bzl", "IO_GRPC_GRPC_JAVA_ARTIFACTS")
#load("@io_grpc_grpc_java//:repositories.bzl", "IO_GRPC_GRPC_JAVA_OVERRIDE_TARGETS")
#load("@io_grpc_grpc_java//:repositories.bzl", "grpc_java_repositories")
#grpc_java_repositories()
#
#maven_install(
#    artifacts = [
#        "io.grpc:grpc-netty-shaded:%s" % GRPC_JAVA_VERSION,
#        "io.grpc:grpc-protobuf:%s" % GRPC_JAVA_VERSION,
#        "io.grpc:grpc-stub:%s" % GRPC_JAVA_VERSION,
#        "io.github.lognet:grpc-spring-boot-starter:%s" % GRPC_SPRING_BOOT_VERSION,
#        "org.springframework.boot:spring-boot-autoconfigure:%s" % SPRING_BOOT_VERSION,
#        "org.springframework.boot:spring-boot-test-autoconfigure:%s" % SPRING_BOOT_VERSION,
#        "org.springframework.boot:spring-boot-test:%s" % SPRING_BOOT_VERSION,
#        "org.springframework.boot:spring-boot:%s" % SPRING_BOOT_VERSION,
#        "org.springframework.boot:spring-boot-starter-web:%s" % SPRING_BOOT_VERSION,
#        "org.springframework.cloud:spring-cloud-starter-consul-discovery:%s" % SPRING_CLOUD_CONSUL_VERSION,
#    ] + IO_GRPC_GRPC_JAVA_ARTIFACTS + IO_GRPC_GRPC_KOTLIN_ARTIFACTS,
#    generate_compat_repositories = True,
#    override_targets = dict(IO_GRPC_GRPC_JAVA_OVERRIDE_TARGETS.items() + IO_GRPC_GRPC_KOTLIN_OVERRIDE_TARGETS.items()),
#    repositories = [
#        "https://repo.maven.apache.org/maven2/",
#    ],
#    maven_install_json = "//:maven_install.json",
#)
#load("@maven//:defs.bzl", "pinned_maven_install")
#pinned_maven_install()
#
#load("@maven//:compat.bzl", "compat_repositories")
#compat_repositories()
