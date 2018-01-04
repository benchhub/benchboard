from shutil import copytree, ignore_patterns

# hard coded path to sync local gommon to vendor folder https://github.com/benchhub/benchboard/issues/4
copytree('/home/at15/workspace/src/github.com/dyweb/gommon',
         '/home/at15/workspace/src/github.com/benchhub/benchboard/agent/vendor/github.com/dyweb/gommon',
         ignore=ignore_patterns('.git', 'vendor', '.idea'))
