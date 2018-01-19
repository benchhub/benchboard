# Design

Draft

- store stuff in `.benchboard` in project
  - might need to tell IDE and editor to not track it
  - ignore it....
- use a lock file to avoid having multiple microbenchmark running a local project
  - ask user to `rm .benchboard/lock.json`
- runner (might be implemented in `benchhub/lib/runner`)
  - kind of like travis ...
- `benchboard serve` to start a visualization server of previous benchmarks
  - maybe some public data? ...
- `benchboard run [benchmark-name]...`
  - lock
  - make a copy of config
  - store stdout and stderr to file (might stream it as well)
    - [ ] checkout sockeye's python wrapper https://github.com/awslabs/sockeye/blob/master/sockeye/train.py 
      - seems just using python builtin logger to log to both destination https://github.com/awslabs/sockeye/blob/master/sockeye/log.py#L67
  - store raw benchmark results, i.e. csv file, go benchmark output
  - parse the files to reformat, index and aggregate data
  - visualize
- `benchboard list`
  - list all available benchmark (as tree?) 
    - just used to discover how many benchmark in a go repo
  
micro benchmark

````yaml
type: micro
language: go
framework: go-bultin-bench # not every language ships w/ a default microbenchmark framework
params:
    p1: p1
    p2: p2
test: AllocFreeFnv64
tests:
    AllocFreeFnv64:
      comment: alloca free should use less bytes
      expects:
        - ? < ? etc....
    StringToBytes:
      comment: the conversion should not bee free
````

marco benchmark

````yaml
type: database
framework: YCSB # TODO: there are many fork of YCSB though ... dialect of YCSB
framework-dialect: github.com/benchhub/YCSB
workload: tsworkload
params:
  a: b 
````

## Go

- https://godoc.org/golang.org/x/perf/cmd/benchstat allow compare benchmark result files
- https://godoc.org/golang.org/x/tools/cmd/compilebench benchmark go compiler ...