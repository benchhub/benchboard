# Config

config file format

  
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