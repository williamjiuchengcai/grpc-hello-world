load("@rules_oci//oci:defs.bzl", "oci_image", "oci_load", "oci_push")
load("@rules_pkg//pkg:tar.bzl", "pkg_tar")

cc_binary(
    name = "greeter_callback_client",
    srcs = ["greeter_callback_client.cc"],
    defines = ["BAZEL_BUILD"],
    deps = [
        "@com_github_grpc_grpc//:grpc++",
        "//protos:helloworld_cc_grpc",
        "@abseil-cpp//absl/flags:flag",
        "@abseil-cpp//absl/flags:parse",
    ],
)

cc_binary(
    name = "greeter_callback_server",
    srcs = ["greeter_callback_server.cc"],
    defines = ["BAZEL_BUILD"],
    deps = [
      "@com_github_grpc_grpc//:grpc++",
      "@com_github_grpc_grpc//:grpc++_reflection",
      "//protos:helloworld_cc_grpc",
      "@abseil-cpp//absl/flags:flag",
      "@abseil-cpp//absl/flags:parse",
      "@abseil-cpp//absl/strings:str_format",
    ],
)

pkg_tar(
  name = "greeter_callback_server_tar",
  srcs = [":greeter_callback_server"],
)

oci_image(
  name = "greeter_callback_server_image",
  base = "@docker_lib_ubuntu",
  tars = [":greeter_callback_server_tar"],
  entrypoint = ["/greeter_callback_server"],
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
  repository = "us-central1-docker.pkg.dev/jcwc-summarization-dev/jcwc-oss-dev/helloworld"
)
