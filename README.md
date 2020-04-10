# goget
`goget` is a mini tool that makes HTTP requests to given URL(s) and displays hashed response

## install:
You need to have [`Go`](https://golang.org/) installed in your system. Then clone the repo and run:
```
go install .
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
$GOPATH/bin/goget -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com 

# outputs
http://google.com 010c6ccacb2c948ef2842a86c0e56447
http://adjust.com 51b0f83c5612591d00f6912559498652
http://facebook.com d9cd5022f1ba405b50a6cce0b0806728
http://yandex.com 8af81f08500b3a5cc234cf5933c28f50
http://twitter.com e4afe711b20950599e5c48eedab11d6d
http://reddit.com/r/funny e3df1e9b50c1a3391ec8553e7a292a07
http://yahoo.com 927ca1fe09a7478039729d177c85c58a
http://reddit.com/r/notfunny 84cd2f402466be4bf9d2d685cbcd88ef
http://baroquemusiclibrary.com 24d185566c069eb686b31b81a6253515
```