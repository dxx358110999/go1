local stock = tonumber(redis.call('get', KEYS))
if stock <= 0 then return 0 end
redis.call('decr', KEYS)
return 1
