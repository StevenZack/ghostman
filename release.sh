replace -f vars/mode.go debug release
htmltostring
go build -o release/mac/ghostman.app/Contents/MacOS/ghostman main.go
cp /Users/stevenzacker/go/bin/dylib/sciter-osx-64.dylib release/mac/ghostman.app/Contents/MacOS/