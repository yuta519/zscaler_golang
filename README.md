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

2. Build main.go
Run below command.
> make build

After running, you could find `zia` file on your current directory.

3. Using `zia`
You could try some `zia` command. Just run below, you can see usages.
> ./zia

```
Commands:
        zia credential                      # Show credential info placed in config.ini

        zia adminuser COMMAND               # Run a command about adminusers
                                              ls
                                              create
                                                <role_name>
                                                <login_name>
                                                <user_name>
                                                <password>
                                                <admin_scope>

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
```
