# `gyaml-benchmarks`

Benchmarks for [gyaml](https://github.com/m4l1c1ou5/gyaml) alongside [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3)

```
BenchmarkGYAMLGet-8                       615477              1932 ns/op            2882 B/op         36 allocs/op
BenchmarkGYAMLGetBytes-8                  614497              2001 ns/op            2882 B/op         36 allocs/op
BenchmarkGYAMLGetComplex-8                 76069             15730 ns/op           17504 B/op        267 allocs/op
BenchmarkGYAMLGetArray-8                   35044             34759 ns/op           36509 B/op        569 allocs/op
BenchmarkGYAMLParse-8                   1000000000               0.3381 ns/op          0 B/op          0 allocs/op
BenchmarkGYAMLValid-8                      71376             16656 ns/op           18640 B/op        273 allocs/op
BenchmarkYAMLv3UnmarshalMap-8              70812             17210 ns/op           18632 B/op        273 allocs/op
BenchmarkYAMLv3UnmarshalStruct-8           73639             17369 ns/op           16608 B/op        230 allocs/op
BenchmarkGYAMLMultipleGets-8              198769              6816 ns/op            8648 B/op        110 allocs/op
BenchmarkYAMLv3MultipleGets-8              72682             17468 ns/op           18632 B/op        273 allocs/op
BenchmarkGYAMLDeepPath-8                   27262             44245 ns/op           47480 B/op        748 allocs/op
BenchmarkYAMLv3DeepPath-8                  32686             36935 ns/op           37096 B/op        613 allocs/op
BenchmarkGYAMLQueryAll-8                   32304             35570 ns/op           40040 B/op        575 allocs/op
BenchmarkGYAMLQueryConditional-8           35884             37808 ns/op           37255 B/op        548 allocs/op
BenchmarkGYAMLForEach-8                    39108             30795 ns/op           57048 B/op        402 allocs/op
BenchmarkGYAMLArray-8                      10000            114170 ns/op          226442 B/op       1484 allocs/op
BenchmarkGYAMLMap-8                        44298             27203 ns/op           38856 B/op        445 allocs/op
BenchmarkGYAMLGetString-8                 616413              1894 ns/op            2800 B/op         36 allocs/op
BenchmarkGYAMLGetInt-8                    613545              2004 ns/op            2800 B/op         36 allocs/op
BenchmarkGYAMLGetBool-8                    35438             34354 ns/op           34744 B/op        567 allocs/op
BenchmarkGYAMLLargeDocument-8              28072             42936 ns/op           43584 B/op        680 allocs/op
BenchmarkYAMLv3LargeDocument-8             32971             37913 ns/op           33560 B/op        564 allocs/op
```


System Information on which benchmarks were recorded:
```
Go Version: go version go1.25.1 darwin/arm64
System: macOS
Model Name: MacBook Air
Chip: Apple M3
Memory: 8 GB
```

Last run: Nov 22, 2024

## Usage

```sh
make quick
```
