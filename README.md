# gap

[![pipeline](https://github.com/acim/gap/actions/workflows/pipeline.yaml/badge.svg)](https://github.com/acim/gap/actions/workflows/pipeline.yaml)
[![Go Reference](https://pkg.go.dev/badge/go.acim.net/gap.svg)](https://pkg.go.dev/go.acim.net/gap)
[![Go Report](https://goreportcard.com/badge/go.acim.net/gap)](https://goreportcard.com/report/go.acim.net/gap)
![Go Coverage](https://img.shields.io/badge/coverage-91.4%25-brightgreen?style=flat&logo=go)

Library for easy implementation of REST API's in Go. Besides helpers to decode and encode JSON payloads using standard library _http.HandlerFunc_ type of handler, this library supports handler functions returning errors, e.g. *func(w http.ResponseWriter, req *http.Request) error)\* and makes errors handling much easier.

## Warning :construction:

This project is in an early stage so you can expect API breaking changes until the first major release.

## License

Licensed under either of

- Apache License, Version 2.0
  ([LICENSE-APACHE](LICENSE-APACHE) or http://www.apache.org/licenses/LICENSE-2.0)
- MIT license
  ([LICENSE-MIT](LICENSE-MIT) or http://opensource.org/licenses/MIT)

at your option.

## Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in the work by you, as defined in the Apache-2.0 license, shall be
dual licensed as above, without any additional terms or conditions.
