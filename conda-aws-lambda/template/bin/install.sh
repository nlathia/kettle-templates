# Makefile command: make install
# This script will create a conda virtual environment 
# And install all of the dependencies defined in the
# requirements file(s).

set -e

source $(dirname $0)/_config.sh
conda create -y $1 python=$PYTHON_VERSION
conda install --force-reinstall -y --name $1 -c conda-forge --file requirements.txt
conda install --force-reinstall -y --name $1 -c conda-forge --file requirements-dev.txt
conda activate $1

echo "\n ✅  Done: $1 activated."
