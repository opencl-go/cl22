# Go wrapper library for OpenCL 2.2

[![Go version of Go module](https://img.shields.io/github/go-mod/go-version/opencl-go/cl22.svg)](https://github.com/opencl-go/cl22)
[![GoDoc reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/opencl-go/cl22)
[![GoReportCard](https://goreportcard.com/badge/github.com/opencl-go/cl22)](https://goreportcard.com/report/github.com/opencl-go/cl22)
[![License](https://img.shields.io/github/license/opencl-go/cl22.svg)](https://github.com/opencl-go/cl22/blob/main/LICENSE)
[![OpenCL 2.2](https://img.shields.io/badge/OpenCL-2.2-green.svg)][opencl-api]

This library provides a complete wrapper for the OpenCL 2.2 API.
If you require a different API level, refer to [the opencl-go project][opencl-go] to see which versions are available.

**This is work-in-progress. The wrapper is not yet in a state to provide useful functionality**

## Usage

To build and work with this library, you need an OpenCL SDK installed on your system.
Refer to [the documentation on opencl-go][opencl-go] on how to do this.

The API requires knowledge of the [OpenCL API][opencl-api]. While the wrapper hides some low-level C-API details,
there is still heavy use of `unsafe.Pointer` and the potential for memory access-violations if used wrong.

[opencl-api]: https://registry.khronos.org/OpenCL/sdk/2.2/docs/man/html/
[opencl-go]: https://opencl-go.github.com

## License

This project is based on the MIT License. See `LICENSE` file.

The API documentation is, in part, based on the official asciidoctor source files from https://github.com/KhronosGroup/OpenCL-Docs,
licensed under the Creative Commons Attribution 4.0 International License; see https://creativecommons.org/licenses/by/4.0/ .
