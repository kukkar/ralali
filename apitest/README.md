#apitest

ApiTest is a repository for ralali assignment.

# Modules
 
 
 
# Pre-requisites
1. Go 1.9+
    
* Set-up the application log directory. 
```bash
  sudo mkdir /var/log/ralali/apitest
  sudo chmod -R 777 /var/log/ralali/apitest
  ```
  The above directory will used to store all logs.  
    
* To build the application execute the below command:-
  ```bash
  cd $GOPATH/src/github.com/ralali/apitest
  go run main.go 
  ```
    
* To build the docker image :-
  ```bash
  cd $GOPATH/src/github.com/ralali/apitest
  docker build .


## TESTS 
```bash
ginkgo -r -v=true -cover=true  ./...
```  
## For single package TESTS 
```bash
ginkgo -r -v=true -cover=true  ./path/of/the/package (Example: ./service/brand/create)
```  