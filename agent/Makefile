.PHONY: sync-gommon
sync-gommon:
# TODO: ignore .git and vendor
	rm -rf vendor/github.com/dyweb/gommon
	python script/sync_gommon.py
# cp -r ../../../dyweb/gommon ./vendor/github.com/dyweb/gommon