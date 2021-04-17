## How to run

https://github.com/opencv/opencv/tree/master/data/haarcascades
から haarcascade_frontalface_default.xml を取得して使う

```bash
$ wget https://raw.githubusercontent.com/opencv/opencv/master/data/haarcascades/haarcascade_frontalface_default.xml
```

```bash
$ go run main.go 0 haarcascade_frontalface_default.xml 
```
