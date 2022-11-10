echo "creating the ouput folder"
mkdir -p out

echo "inserting the README.md file"
[[ -e ./README.md ]] && cp ./README.md ./out/README.md

echo "linux\n"
env GOOS=linux;GOARCH=adm64; go build -ldflags="-X 'jim/utils.Version=$1'" -o ./out/jim ./cmd/
cd ./out && tar vcfz jim-linux-adm64.tar.gz README.md jim && cd -
rm ./out/jim

echo "darwin\n"
env GOOS=darwin;GOARCH=adm64; go build -ldflags="-X 'jim/utils.Version=$1'" -o ./out/jim ./cmd/
cd ./out && tar vcfz jim-darwin-adm64.tar.gz README.md jim && cd -
rm ./out/jim

echo "windows\n"
env GOOS=windows;GOARCH=adm64; go build -ldflags="-X 'jim/utils.Version=$1'" -o ./out/jim.exe ./cmd/
cd ./out && tar vcfz jim-windows-adm64.tar.gz README.md jim.exe && cd -
rm ./out/jim.exe

rm ./out/README.md

echo "output"
ls ./out/


