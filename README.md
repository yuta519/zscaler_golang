# Zscaler Golang

## What is this repository?

- This repository is a library to use Zscaler API for GO.
- I'm developing ZIA SDK for GO in accordance with [Zsclaer help page](https://help.zscaler.com/zia/api/)


## Prerequisite

### Zsclaer (ZIA)
- You need ZIA API key, Adminuser ID and Password.
- ZIA API is not enabled initially, so you need send a support ticket.

### Stacks
- Go
    - Version 1.8 or above
- Library
    - gopkg.in/ini.v1
- API
  - Zscaler Internet Access


## Quick Start

1. Prepare ZIA credentiasl
You need write your ZIA credential inforamtion at **config/config.ini**.
```
[zia]
USERNAME=yuta.kawamura@zscaler.net
PASSWORD=IloveZscaler
HOSTNAME=admin.zscaler.net
APIKEY=XXXXXXXXXXXXXXXXXXXXX
```
<img src="https://raw.githubusercontent.com/yuta519/assets/master/zscaler_golang/setup.gif" alt="How to setup">

2. Build main.go
Run below command.
> make build

After running, you could find `zia` file on your current directory.

<img src="https://raw.githubusercontent.com/yuta519/assets/master/zscaler_golang/build.gif" alt="How to build">


3. Using `zia`
You could try some `zia` command. Just run below, you can see usages.
> ./zia

```
Commands:
  zia credential                      # Show credential info placed in config.ini

  zia auth exclude                    # Show exempted urls of cookie auth login

  zia adminuser COMMAND               # Run a command about adminusers
                                        ls
                                        create

  zia adminrole COMMAND               # Run a command about adminusers
                                        ls

  zia urlcategory COMMAND             # Run a command about urlcategory
                                        ls
                                        lookup [URLS]

  zia urlfilter COMMAND                 # Run a command about firewall
                                        ls
                                        --id [ID]
                                        --all

  zia firewall COMMAND                 # Run a command about firewall
                                        ls
                                        --id [ID]
                                        --all

  zia network COMMAND                 # Run a command about network
                                        ipdst ls
                                        ipsrc ls

```
<img src="https://raw.githubusercontent.com/yuta519/assets/master/zscaler_golang/initial_run.gif" alt="How to build">
