echo "creating the ouput folder"
mkdir -p out

echo "inserting the README.md file"
[[ -e ./README.md ]] && cp ./README.md ./out/README.md

echo "linux"
env GOOS=linux;GOARCH=adm64; go build -ldflags="-X 'jim/utils.Version=$1'" -o ./out/jim -v ./cmd/ 
tar vcfz ./out/jim-linux-adm64.tar.gz ./out/README.md ./out/jim
rm ./out/jim

echo "darwin"
env GOOS=darwin;GOARCH=adm64; go build -ldflags="-X 'jim/utils.Version=$1'" -o ./out/jim -v ./cmd/
tar vcfz ./out/jim-darwin-adm64.tar.gz ./out/README.md ./out/jim
rm ./out/jim

echo "windows"
env GOOS=windows;GOARCH=adm64; go build -ldflags="-X 'jim/utils.Version=$1'" -o ./out/jim.exe -v ./cmd/
tar vcfz ./out/jim-windows-adm64.tar.gz ./out/README.md ./out/jim.exe
rm ./out/jim.exe

rm ./out/README.md

echo "output"
ls ./out/


