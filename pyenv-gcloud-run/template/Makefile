VIRTUALENV_NAME=gcloud-run.$(shell pwd | rev | cut -d '/' -f 1 | rev)
IMAGE_URL=gcr.io/{{ .ProjectID }}/{{ .ProjectName }}:latest

.PHONY: install uninstall test clean docker localhost

clean:
	@./bin/cleanup.sh

uninstall: clean
	@./bin/remove_pyenv.sh

install: uninstall
	@./bin/setup_pyenv.sh $(VIRTUALENV_NAME)

test:
	@pytest

docker: clean
	docker build -f Dockerfile -t $(IMAGE_URL) .

localhost: docker
	@./bin/launch.sh $(IMAGE_URL)
