
# start gin server
go run server.go

# run load testing
wrk -t 1 -s scripts/increasement.lua http://localhost:3000

# then you can see PostForm is not thread safe