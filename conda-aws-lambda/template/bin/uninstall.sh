# Makefile command: make uninstall
# Uninstall the conda virtual environment
set -e

source $(dirname $0)/_config.sh

conda deactivate 
conda remove -y --name $1 --all
