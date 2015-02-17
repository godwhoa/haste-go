Hastebin cli in golang
======================

Install 
=======
```
$ go get github.com/godwhoa/haste-go
```
```
$ cd $GOPATH/src/github.com/godwhoa/haste-go
```
```
$ go build;sudo mv haste /usr/bin/
```

Usage
=====
From stdin<br/>
Link to the paste will be in your clipboard.
```
$ cat <file> | haste | xclip
```

TODO:<br/>
Take filename as an argument 



