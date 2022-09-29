# golang-fix-lib

Golang FIX library for PDAX Equities

### Installation

To install golang-fix-lib library, use ```go get```:
```
$ go get github.com/contrerasarf03/golang-fix-lib
```

### Get Latest update

To update to the latest version, use
```
$ go get -u github.com/contrerasarf03/golang-fix-lib
```



## TO use

```
func main() {
	fixAPI, _ := NewFixAPI("config/session.conf", 40)

	if err := fixAPI.Run(); err != nil {
		panic("Failed to start initiatior")
	}
	err := f.SecurityList("SENDER_COMP_ID", "TARGET_COMP_ID")
	logrus.Error(err)

	for {
		logrus.Info(<- f.RespChan)
	}
}
```
