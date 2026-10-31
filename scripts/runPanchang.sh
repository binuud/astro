export DATA_PATH="./data/"
export ASTRO_EPHIMERIS_ENDPOINT="localhost:9090"

## when debugging, we can run ephimeris seperately
#export ASTRO_EPHIMERIS_ENDPOINT="localhost:9090"

source ~/Code/local/dev-tools/python-env/3.14/bin/activate
cd ./scripts/pachangJson/ && python3 getPanchang.py