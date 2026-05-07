# Privacy Desensitization Benchmark Comparison

## Scope

This document records the benchmark comparison between the legacy reflection-based implementation (V1) and the handwritten recursive implementation (manual).

## Environment

- OS: macOS
- CPU: Apple M2
- Go benchmark command used during comparison:

```zsh
go test ./privacy -run '^$' -bench 'BenchmarkDesensitize_CompareReflectVsManualV2$' -benchmem
```

## Result Snapshot

| Case | ns/op | B/op | allocs/op |
| --- | ---: | ---: | ---: |
| reflect (V1) | 1939 | 2984 | 49 |
| manual (handwritten) | 338.7 | 1024 | 4 |

## Gap

- Latency: manual is about **5.72x faster** (`~82.5%` lower ns/op).
- Memory bytes: manual uses about **65.7% less** memory (`B/op`).
- Allocation count: manual uses about **91.8% fewer** allocations (`allocs/op`).

## Final Decision

- Adopt the handwritten desensitization implementation as the final solution in `privacy`.
- Remove legacy reflection-based V1 code from the package.
- Remove `V2` suffix naming and keep a single unified contract and model naming.

## Current Benchmark Entry

Use the following command to track the current handwritten baseline:

```zsh
go test ./privacy -run '^$' -bench 'BenchmarkDesensitize_Manual$' -benchmem
```

