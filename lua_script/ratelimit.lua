local cur = tonumber(redis.call('get', KEYS) or 0)
if cur >= tonumber(ARGV) then return 0 end
redis.call('incr', KEYS); redis.call('expire', KEYS, ARGV)
return cur + 1
