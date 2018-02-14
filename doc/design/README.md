# Design

- store project scope results in `.benchboard` in project
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
  - visualize in realtime is possible
- `benchboard list` or use `-g` flag
  - list all available benchmark (as tree?) 
    - just used to discover how many benchmark in a go repo
- `benchboard ui` to show a browser for existing results

