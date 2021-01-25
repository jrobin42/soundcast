# SoundCast technical test  

This test was about a creation of Golang micro service API that listen for  `[GET]/info?ua={url_encoded_ua}` and return a JSON object declared as mention below filled with information queried from a JSON database  


```json
Json Response :

{
    "app": string,
    "device": string,
    "bot": bool
}
```

## Usage

To run this project, execute the following lines

```shell
git clone https://github.com/jrobin42/soundcast.git $GOPATH/src/soundcast
cd soundcast/

go get -d ./...

cd api/
go build api.go

./api
```