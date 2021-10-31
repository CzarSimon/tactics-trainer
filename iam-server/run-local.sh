cd cmd/server
echo "Building iam-server"
go build
cd ../..
mv cmd/server/server iam-server

export MIGRATIONS_PATH='./resources/db/mysql'
export JWT_SECRET='084415a2891cd35485f690dc19bbcedb22a9432bc962932726f89be77bf56bd7'
export DB_TYPE='mysql'
export DB_HOST='127.0.0.1'
export DB_PORT='13306'
export DB_NAME='iamserver'
export DB_USERNAME='iamserver'
export DB_PASSWORD='7535807ef23504ca84c7200671611ebc'
# export DB_FILENAME='./resources/testing/test.db'
export KEY_ENCRYPTION_KEYS_PATH='./resources/testing/key-encryption-keys.txt'

./iam-server