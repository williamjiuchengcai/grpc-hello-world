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
