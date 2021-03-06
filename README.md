# TEMPLATE_SUBSTITUTE_PROJECT_NAME


## To write

- README.md
- docs
- examples
- test

- cmd
- pkg
- internal

(if not golang)
- Dockerfile
- Makefile: vendor, fmt, lint

(vars)
OKD_PRODUCTION_SERVER_HOSTNAME
OKD_STAGING_SERVER_HOSTNAME

## Gitlab Configurations

- creat API token
- CICD Variables: GITLAB_API_TOKEN, GITLAB_PRODUCTION_OKD_TOKEN, GITLAB_STAGING_OKD_TOKEN
- master: protected: no one
- pipeline must suceed
- Jira integration?

## Development

You can use docker-compose to run a local copy of the application. Here are useful Make commands:
* `make build` builds Docker image locally.
* `make run` spins up local Docker containers so you can access via web or command line.
* `make test` runs all tests locally.
* `make rsh` spins up a temporary container based on the built image and gives you a bash inside that container.
* `make debug` spins up a temporary container based on the build image with sleep entrypoint, and gives you a bash inside that container to debug in case container stops immediately.


## Libraries

https://oxozle.com/awetop/avelino-awesome-go/

```
log:
    sirupsen/logrus 17.5k
    kubernetes/klog 280
    log

cli:
    command:
        spf13/cobra 21k
        urfave/cli 15k
    flag:
        flag
        spf13/pflag 1k
configuration:
    viper 15k
error:
    errors
http:
    net
    valyala/fasthttp 14k
web-framework:
    gin 46k
    beego/beego 26k
    echo 20k
    revel 12k
    goa 4k
routers:
    gorilla/mux 14k
    httprouter 12k
dns:
    miekg/dns 5k
```

## License

GNU General Public License v3.0, see [LICENSE](LICENSE).
