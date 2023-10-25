wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"

request = function()
    local id = math.random(1, 100000)
    local body = '{"id": "' .. id .. '"}'
    return wrk.format("POST", nil, nil, body)
end
