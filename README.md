# library i used

 1. calculate geo distance: https://github.com/kellydunn/golang-geo
 2. for project struct:  https://github.com/spf13/cobra

# invite
this project has package called "invite". this package has all functionality to read customer json and generate customer list based on provided filter.
this package works like library. any one can use this package as their project library

# how to run
by default this programs default maximum distance is 100km and use ./customer.json to get customer list
- run with default parameter
```go run main.go list```

- run with custom file location
```go run main.go list -f=./data/test.txt```

- run with custom distance
	```go run main.go list -f=./data/test.txt -d=100```
