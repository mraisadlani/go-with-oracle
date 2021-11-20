# Go with Oracle database
Implementasi database oracle kedalam golang

# How to using swagger
1. Install generator swagger menggunakan perintah :
```sh
go get -u github.com/swaggo/swag/cmd/swag
```
2. tambahkan module swagger menggunakan perintah :
```sh
go get -u github.com/swaggo/http-swagger
go get -u github.com/alecthomas/template
```

3. tambahkan general API di swaggo:
```sh
docs.SwaggerInfo.Title = "Go Oracle"
docs.SwaggerInfo.Description = "Gin CRUD with Database Oracle"
docs.SwaggerInfo.Version = "1.0"
docs.SwaggerInfo.Host = fmt.Sprintf("%s:%v", config.C.App.HOST, config.C.App.PORT)
docs.SwaggerInfo.BasePath = "/api/v1"
docs.SwaggerInfo.Schemes = []string{"http", "https"}
```

4. Untuk menjalankan generator gin-swagger menggunakan perintah :
```sh
swag init -g cmd/main.go
```

5. Jalankan kembali aplikasi menggunakan perintah :
```sh
go run main.go
```

6. Kemudian bisa diakses melalui url :
```sh
http://localhost:8804/swagger/index.html
```

# How to install oracle client in windows
1. Download file oracle client and extract in 1 folder
   - Instant Client Package - Basic
   - Instant Client Package - SDK
2. Download and install Msys2
    - link : https://www.msys2.org/ and set path C:\msys64
```sh
# Update pacman
pacman -Su
# Close terminal and open a new terminal
# Update all other packages
pacman -Su
# Install pkg-config and gcc
pacman -S mingw64/mingw-w64-x86_64-pkg-config mingw64/mingw-w64-x86_64-gcc
```
3. Set path in operating system
```sh
- add to Path :
    > C:\msys64
    > C:\msys64\mingw64\bin
- create path :
    > CGO_CFLAGS : -IC:\instantclient_12_1\sdk\include
    > CGO_LDFLAGS : -LC:\instantclient_12_1 -loci
    > GOBIN : %GOPATH%\bin
    > GOPATH : D:\Workspace
    > GOROOT : C:\Go
    > OCI_INC_DIR : C:\instantclient_12_1
    > OCI_LIB_DIR : %OCI_INC_DIR%\sdk\lib\msvc
    > ORACLE_HOME : %OCI_INC_DIR%\sdk
    > PKG_CONFIG_PATH : C:\msys64\mingw64\lib\pkgconfig

Save and restart your computer
```

# How to install oracle client in Linux
1. Download file oracle client
    - oracle-instantclient12.1-basic-12.1.0.2.0-1.x86_64.rpm 
    - oracle-instantclient12.1-sqlplus-12.1.0.2.0-1.x86_64.rpm
    - oracle-instantclient12.1-devel-12.1.0.2.0-1.x86_64.rpm

install alien for running file rpm :
```sh
sudo apt-get install alien
```
convert the rpm file and install :
```sh
- sudo alien -i oracle-instantclinet*-basic-*.rpm
- sudo alien -i oracle-instantclinet*-devel-*.rpm
- sudo alien -i oracle-instantclinet*-sqlplus-*.rpm
```
Install libaio1 :
```sh
sudo apt-get install libaio1
```

add variable in .bashrc/ .bash_profile :
```sh
- export ORACLE_HOME=/usr/lib/oracle/12.1/client64
- export LD_LIBRARY_PATH=$ORACLE_HOME/lib
- export PATH=$PATH:$ORACLE_HOME/bin 
```

running in command line :
```sh
- sudo nano /etc/ld.so.conf.d/oracle.conf && chmod o+r /etc/ld.so.conf.d/oracle.conf
insert and save : export LD_LIBRARY_PATH=/usr/lib/oracle/12.1/client64/lib/${LD_LIBRARY_PATH:+:$LD_LIBRARY_PATH}

- sudo nano /etc/profile.d/oracle.sh && chmod o+r /etc/profile.d/oracle.sh
insert and save : export ORACLE_HOME=/usr/lib/oracle/12.1/client64
```

and running in command and restart
```sh
sudo ldconfig
```

*note: If libsqlplus.so: cannot open shared object file: No such file or directory is displayed when execute sqlplus
```sh
> sudo sh -c "echo /usr/lib/oracle/12.1/client64/lib > /etc/ld.so.conf.d/oracle.conf"
> sudo ldconfig
```

# Other Tools
| Tools | Version |
| ----- | ----- |
| [Oracle Client](https://www.oracle.com/database/technologies/instant-client/downloads.html) | Version 12.1.0.2.0  |
| GCC | for windows : https://www.msys2.org/ , for linux : sudo apt install build-essential |

# Tools for build App
| Tools | Version | Description |
| ----- | ----- | ----- |
| [gin](https://github.com/gin-gonic/gin) | v1.7.4 | Framework Golang |
| [godror](https://github.com/godror/godror) | v0.28.1 | Library database oracle |
| [logrus](https://github.com/sirupsen/logrus) | v1.8.1 | Library logging in golang |
| [viper](https://github.com/spf13/viper) | v1.9.0 | For manage file .yml |
| [swaggo](https://github.com/swaggo/swag) | v1.9.0 | Library swagger in golang |

# Todolist Application
- [x] CRUD with Oracle
- [x] Implement Golang for get data from Function
- [x] Implement Golang for get data from View
- [x] Implement Golang for set data to procedure

# Contribute
Support saya agar lebih banyak berkontribusi dalam membuat sebuah project sederhana menggunakan bahasa pemrograman golang
- Saweria : https://saweria.co/mraisadlani
