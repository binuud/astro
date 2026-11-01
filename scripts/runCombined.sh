export MONGO_ENDPOINT="localhost:27017"
export MONGO_DB="astro"
export MONGO_USER=""
export MONGO_PASSWORD=""
export MONGO_ENCRYPTION_KEY=""
export TRANCE_PRODUCTION="false"
export TRANCE_LOG_VERBOSITY="1"
export DATA_PATH="./data/"
#export ASTRO_EPHIMERIS_ENDPOINT="localhost:9090"


export FILE_STORE_PATH="./dataVolume"

## when debugging, we can run ephimeris seperately
export ASTRO_EPHIMERIS_ENDPOINT="localhost:9090"

go run ./cmd/combined/main.go