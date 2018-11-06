graphdb-presentation
===========

Golang seed application with yaml config file + cli

Default command provided:
- start : to start the process.
- version : to display the version (based to git commit / branch / build time)

Package management by [https://github.com/Masterminds/glide](https://github.com/Masterminds/glide)
- install dependency
```bash
glide get package_name
```

Initialize dev:
- install glide
- Replace all occurences of `graphdb-presentation` by your project name.
- install dependencies.
```bash
glide install
```
- run project :
```bash
go run main.go start
```
- run with hot reload (http server)
```
go get github.com/codegangsta/gin
gin --appPort 8080 --buildArgs main.go -i run start
```

Docker
------

- Generate container :
```bash
make build
```

- Run container :
```bash
docker run -d [-v .my-config-file.yml:/.config.yml] project-name:version
```

TODO
----

- Build as CLI
- More documentation
- ci/cd, docker-hub integration ?
- Templating of the seed ?
  - GRPC
  - REST
  - ...

