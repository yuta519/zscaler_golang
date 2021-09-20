# Zscaler Golang

## What is this repository?
---

- This repository is a library to use Zscaler API for GO.
- I'm developing ZIA SDK for GO in accordance with [Zsclaer help page](https://help.zscaler.com/zia/api/)


## Prerequisite
---

### Zsclaer (ZIA)
- You need ZIA API key, Adminuser ID and Password.
- ZIA API is not enabled initially, so you need send a support ticket.

### Go
- Go 
    - Version 1.6 or above
- Library
    - gopkg.in/ini.v1



## Quick Start
---

1. Prepare ZIA credentiasl
You need write your ZIA credential inforamtion at **config/config.ini**.
```
[credential]
USERNAME=yuta.kawamura@zscaler.net
PASSWORD=IloveZscaler
HOSTNAME=admin.zscaler.net
APIKEY=XXXXXXXXXXXXXXXXXXXXX
```

2. Using codes under ***pkg/XXX.go*** like src/main.go .