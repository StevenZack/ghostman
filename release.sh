replace -f vars/mode.go debug release
htmltostring
go install
cp ~/go/bin/ghostman ~/Desktop/ghostman