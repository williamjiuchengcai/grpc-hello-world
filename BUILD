load("@gazelle//:def.bzl", "gazelle")
load("@rules_go//go:def.bzl", "go_library")
load("@rules_oci//oci:defs.bzl", "oci_image", "oci_load", "oci_push")
load("@rules_pkg//pkg:tar.bzl", "pkg_tar")

cc_binary(
    name = "greeter_callback_client",
    srcs = ["greeter_callback_client.cc"],
    defines = ["BAZEL_BUILD"],
    deps = [
        "//protos:helloworld_cc_grpc",
        "@abseil-cpp//absl/flags:flag",
        "@abseil-cpp//absl/flags:parse",
        "@com_github_grpc_grpc//:grpc++",
    ],
)

cc_binary(
    name = "greeter_callback_server",
    srcs = ["greeter_callback_server.cc"],
    defines = ["BAZEL_BUILD"],
    deps = [
        "//protos:helloworld_cc_grpc",
        "@abseil-cpp//absl/flags:flag",
        "@abseil-cpp//absl/flags:parse",
        "@abseil-cpp//absl/strings:str_format",
        "@com_github_grpc_grpc//:grpc++",
        "@com_github_grpc_grpc//:grpc++_reflection",
    ],
)

pkg_tar(
    name = "greeter_callback_server_tar",
    srcs = [":greeter_callback_server"],
)

oci_image(
    name = "greeter_callback_server_image",
    base = "@docker_lib_ubuntu",
    entrypoint = ["/greeter_callback_server"],
    tars = [":greeter_callback_server_tar"],
)

# Use with 'bazel run' to load the oci image into a container runtime.
# The image is designated using `repo_tags` attribute.
#
# To load image: `bazel run //:greeter_callback_server_load`
# To run the docker image: `docker run --rm greeter_callback_server:latest`
oci_load(
    name = "greeter_callback_server_load",
    image = ":greeter_callback_server_image",
    repo_tags = ["greeter_callback_server:latest"],
)

# To push image: `bazel run :greeter_callback_server_push -- --tag latest`
oci_push(
    name = "greeter_callback_server_push",
    image = ":greeter_callback_server_image",
    repository = "us-central1-docker.pkg.dev/jcwc-summarization-dev/jcwc-oss-dev/helloworld",
)

# gazelle:prefix github.com/williamjiuchengcai/medlmpp-incubation-platform-jcwc-dev
gazelle(name = "gazelle")

go_library(
    name = "medlmpp-incubation-platform-jcwc-dev",
    srcs = ["greeter_callback_client.go"],
    importpath = "github.com/williamjiuchengcai/medlmpp-incubation-platform-jcwc-dev",
    visibility = ["//visibility:public"],
    deps = [
        "//protos:helloworld_go_proto",
        "@org_golang_google_api//idtoken:go_default_library",
        "@org_golang_google_grpc//:grpc",
        "@org_golang_google_grpc//credentials",
        "@org_golang_google_grpc//metadata",
    ],
)
