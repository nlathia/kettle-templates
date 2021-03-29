# Makefile command: make uninstall
# Uninstall the conda virtual environment
set -e

conda remove --name $1 --all
