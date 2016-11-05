## stream-api

#### Tech Stack
* Go1.5.1
* Echo - Micro Web framework
* sqlx - sql layer extensions to golang's database/sql
* sqlite3 support
* go-validator - model/struct validations
* jwt-go - JWT implementation library


### Go Setup (for local development)
* [Install Go](https://golang.org/dl/)
* [Install GVM](https://github.com/moovweb/gvm)
* Setup GVM/Go environment
```
gvm install go1.4.3
gvm use go1.4.3
gvm install go1.5.1
gvm use go1.5.1
```


### Project Setup (for local development)

* Setup GOPATH project structure
```
cd ~/Source/Go
git clone stream-api src
gvm pkgset use --local
go build stream-api && ./stream-api
```
```
Running with 8 CPUs
Starting server on port 4000
```

* Setup auto-reload (optional)
```
go get github.com/codegangsta/gin
cd src/stream-api
gin -a 4000 -p 4001
```

### Database Setup
* Create DB (sqlite3)


### Docker Setup

Create your Docker Machine, or use the default env.

`eval "$(docker-machine env default)"`

```
git clone stream-api
docker build -t golang-app .
docker run -p 8080:4000 -d golang-app
```

Load `localhost:8080` into your browser. If you are using a Docker VM, use the VM's IP address instead. i.e. `192.168.99.100:8080`

You can view the active VM by running `docker-machine ls`

Load `http://localhost:4001/` in the browser. Code will hot-reload on browser refresh when code changes.



### Go IDE Resources
* https://github.com/fatih/vim-go
* https://github.com/farazdagi/vim-go-ide
* http://farazdagi.com/blog/2015/vim-as-golang-ide/
* https://atom.io/packages/go-plus
* https://github.com/joefitzgerald/go-plus
* https://github.com/nsf/gocode
