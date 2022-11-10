echo "creating the ouput folder"
mkdir -p out

echo "inserting the README.md file"
[[ -e ./README.md ]] && cp ./README.md ./out/README.md

echo "linux"
env GOOS=linux;GOARCH=amd64; go build -ldflags="-X 'jim/utils.Version=$1'" -o ./out/jim ./cmd/
cd ./out && tar -cvzf jim-linux-amd64.tar.gz README.md jim && cd -
rm ./out/jim

echo "darwin"
env GOOS=darwin;GOARCH=amd64; go build -ldflags="-X 'jim/utils.Version=$1'" -o ./out/jim ./cmd/
cd ./out && tar -cvzf jim-darwin-amd64.tar.gz README.md jim && cd -
rm ./out/jim

echo "windows"
env GOOS=windows;GOARCH=amd64; go build -ldflags="-X 'jim/utils.Version=$1'" -o ./out/jim.exe ./cmd/
cd ./out && tar -cvzf jim-windows-amd64.tar.gz README.md jim.exe && cd -
rm ./out/jim.exe

rm ./out/README.md

echo "output"
ls ./out/


