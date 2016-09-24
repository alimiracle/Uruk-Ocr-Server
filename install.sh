export GOPATH=`pwd`
go get github.com/otiai10/gosseract
go install ocr-server
cp -R bin /
cp -R etc /

cd /bin
chmod +x ocr-server
