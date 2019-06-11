

function delay()
    return 1000
end

counter = 0
-- url = "http://localhost:3000"
--url = "https://wtc-go-be-staging-us-east-1.roosterteeth.com"
-- url = "https://wtc.staging.roosterteeth.com"
request = function ()
    wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"
    counter = counter + 1
    print(counter)
    r = wrk.format("POST", "/abc", wrk.headers, "value=" .. counter)
    return r
end
