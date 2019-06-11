
#Start gin server

go run server.go

#Run load testing

wrk -t 1 -s scripts/increasement.lua http://localhost:3000

then you can see PostForm is not thread safe
Maybe because gin use a cache for post parameters
https://github.com/gin-gonic/gin/blob/08b52e5394099db4c2399357e060619c1545083e/context.go#L477
