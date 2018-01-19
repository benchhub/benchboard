# GoBenchUI

UI for overview of your package benchmarks progress. https://github.com/divan/gobenchui

- vcs
  - wrap around git cli, allow filter
  - we may use https://github.com/src-d/go-git
- websocket
  - used to update when running benchmarks in background
  - might use something similar in benchboard for realtime visualization as well
- use `x/tools/benchmark/parse`

main.go

````go
commits, err := vcs.Commits()
ch := RunBenchmarks(vcs, commits, *benchArgs) // start running benchmark in background
// bechmark.go 
vcs.SwitchTo(commit.Hash)
runBnechmark() // it supports both go and gb https://github.com/constabulary/gb
info.AddResult(result)
````