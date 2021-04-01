# Makefile command: make install
# This script will create a conda virtual environment 
# And install all of the dependencies defined in the
# requirements file(s).

set -e

source $(dirname $0)/_config.sh
conda create -y -n $1 python=$PYTHON_VERSION

conda activate $1

pip install -r requirements.txt
pip install -r requirements-dev.txt

echo "\n ✅  Done"

echo "\n 🚀  Activate with: conda activate $1"
