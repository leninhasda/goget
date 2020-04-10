# goget
`goget` is a mini tool that makes HTTP requests to given URL(s) and displays hashed response

## install:
You need to have [`Go`](https://golang.org/) installed in your system. Then clone the repo and run:
```
go install 
```
This will build the binary in `$GOPATH/bin` directory. 

OR run:
```
go build -o ./dist/goget .
```

to building the binary in local `dist` directory.

## usage: 
From source code:
```
go run main.go google.com
```

From binary
```
$GOPATH/bin/goget google.com
```

### Example
```
./goget -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com 

# outputs
http://google.com 44d7b8035b5c796c40f48653fffddff5
http://adjust.com 51b0f83c5612591d00f6912559498652
http://facebook.com 7ea781b0a3701721dc41fcaba595f83e
http://yandex.com 2a53c33ad8baac0274128ee3b6959e04
http://twitter.com 6fea0869722ebfb7ca023056feda4c7b
http://yahoo.com 996e7f8f517b00eacdfd835f92ed9b38
http://reddit.com/r/funny f420bfd1d3947253488698662af15cf7
http://reddit.com/r/notfunny 977f46e7e5f8892160a116626d310b3b
http://baroquemusiclibrary.com 90ae3b48812f3a109715ac1224e7125d
```