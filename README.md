## SwagIP (GO Edition)

A simple way to get your public IP address and other connection related information.

 This project was forked from our other project [https://github.com/StackFocus/SwagIP](https://github.com/StackFocus/SwagIP) to help us learn go.


#### Examples of retrieving your public IP address from Linux/Unix CLI:
```
wget -qO - ip.swagger.pro
curl ip.swagger.pro
fetch -qo - ip.swagger.pro
```

#### Example of retrieving your public IP address from PowerShell 3+:
```
Invoke-RestMethod http://ip.swagger.pro
```

### Authors
- [Kevin Law](https://github.com/thatarchguy)

### Installation
- Install the binary from the Releases or compile yourself:
```
$ go build .
$ ./swagip
```
### Testing
```
$ go test
```

### Docker
Docker is used to scale the application
```
$ docker build -t swagip .
$ docker run -p 0.0.0.0:80:8080 swagip
```

### Screenshots
#### Browser Information:
![Browser Information](screenshots/browser.png?raw=true)

#### Wget Commands:
![Wget Commands](screenshots/wget.png?raw=true)

#### Curl Commands:
![Curl Commands](screenshots/curl.png?raw=true)

#### Fetch Commands:
![Fetch Commands](screenshots/fetch.png?raw=true)

#### PowerShell Commands:
![PowerShell Commands](screenshots/powershell.png?raw=true)