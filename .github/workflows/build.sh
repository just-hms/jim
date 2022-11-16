echo "creating the ouput folder"
mkdir -p out

echo "inserting the README.md file"
[[ -e ./README.md ]] && cp ./README.md ./out/README.md

echo "linux"
env GOARCH=amd64 GOOS=linux go build -ldflags="-X 'jim/constants.Version=$1'" -o ./out/jim ./cmd/
cd ./out && tar -cvzf jim-linux-amd64.tar.gz README.md jim && cd -
rm ./out/jim

echo "darwin"
env GOARCH=amd64 GOOS=darwin go build -ldflags="-X 'jim/constants.Version=$1'" -o ./out/jim ./cmd/
cd ./out && tar -cvzf jim-darwin-amd64.tar.gz README.md jim && cd -
rm ./out/jim

echo "windows"
env GOARCH=amd64 GOOS=windows go build -ldflags="-X 'jim/constants.Version=$1'" -o ./out/jim.exe ./cmd/
cd ./out && tar -cvzf jim-windows-amd64.tar.gz README.md jim.exe && cd -
rm ./out/jim.exe

rm ./out/README.md

echo "output"
ls ./out/


