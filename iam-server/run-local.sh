cd cmd/server
echo "Building iam-server"
go build
cd ../..
mv cmd/server/server iam-server

export MIGRATIONS_PATH='./resources/db/sqlite'
export JWT_SECRET='084415a2891cd35485f690dc19bbcedb22a9432bc962932726f89be77bf56bd7'
export DB_TYPE='sqlite'
export DB_FILENAME='./resources/testing/test.db'
export KEY_ENCRYPTION_KEYS_PATH='./resources/testing/key-encryption-keys.txt'

./iam-server