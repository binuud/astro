export TRANCE_PRODUCTION="false"
export TRANCE_LOG_VERBOSITY="1"
export DATA_PATH="./data/"
#export ASTRO_EPHIMERIS_ENDPOINT="localhost:9090"


export FILE_STORE_PATH="./dataVolume"
export ASTRO_BIN='./astroBin/astroBin'

## when debugging, we can run ephimeris seperately
export ASTRO_EPHIMERIS_ENDPOINT="localhost:9090"

go run ./cmd/mcpServer/main.go