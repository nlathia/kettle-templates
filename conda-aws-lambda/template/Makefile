VIRTUALENV_NAME=aws-lambda.{{ .ProjectName }}

.PHONY: install uninstall test clean invoke-local invoke activate

activate:
	@echo "\n➡️  Use: conda activate $(VIRTUALENV_NAME)"

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
