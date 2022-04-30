-include gomk/main.mk
-include local/Makefile

clean: clean-default
ifeq ($(unameS),Windows)
ifneq ($(wildcard testdata),)
	@powershell -c Remove-Item -Force -Recurse ./testdata
endif
else
	@rm -fr ./testdata
endif
