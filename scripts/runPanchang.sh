export DATA_PATH="./data/"
export ASTRO_EPHIMERIS_ENDPOINT="localhost:9090"

## when debugging, we can run ephimeris seperately
#export ASTRO_EPHIMERIS_ENDPOINT="localhost:9090"
export ASTRO_BIN='./astroBin/astroBin'

source ~/Code/local/dev-tools/python-env/3.14/bin/activate
cd ./scripts/ && python3 getPanchang.py