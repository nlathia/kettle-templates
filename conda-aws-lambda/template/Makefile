VIRTUALENV_NAME=aws-lambda.$(shell pwd | rev | cut -d '/' -f 1 | rev)

.PHONY: install uninstall test clean invoke-local invoke

clean:
	@./bin/cleanup.sh

uninstall: clean
	@./bin/uninstall.sh $(VIRTUALENV_NAME)

install: uninstall
	@./bin/install.sh $(VIRTUALENV_NAME)

test:
	@pytest

invoke-local:
	@./bin/invoke_local.sh 

invoke:
	@./bin/invoke.sh 
