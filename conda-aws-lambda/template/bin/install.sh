# Makefile command: make install
# This script will create a conda virtual environment 
# And install all of the dependencies defined in the
# requirements file(s).

set -e

source $(dirname $0)/_config.sh
conda env create -f environment.yml

echo "\n ✅  Done"

echo "\n 🚀  Activate with: conda activate $1"
