
echo "Welcome to the OptProv type generator!"
echo "Pruning all existing files..."
mkdir -p ./src/models
mkdir -p ../pkg/api/types/
rm ./src/models/* || true 2> /dev/null
rm ../pkg/api/types/* || true 2> /dev/null

for FILE_PATH in $(find ./types/*.json); do
  FILE_NAME=$(basename $FILE_PATH)
  MODEL=${FILE_NAME%.*}

  echo "Generating typescript model: $MODEL"

  `npm bin`/quicktype \
    --src $FILE_PATH \
    --src-lang schema \
    --lang typescript \
    --out src/models/$MODEL.ts

  `npm bin`/quicktype \
    --src $FILE_PATH \
    --src-lang schema \
    --lang golang \
    --package types \
    --out ../pkg/api/types/$MODEL.go
done;

echo "Done!"