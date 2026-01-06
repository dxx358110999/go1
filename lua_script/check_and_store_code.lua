-- 说明：每天不超过指定的验证码短信条数，并且60秒内没有发过知信，
-- 返回0，表示可以发
-- 返回1：表示上一条短信发完还没超过60秒
-- 返回2：表示条数已超

local verify_key = KEYS[1]
local verify_code = tonumber(KEYS[2])
local keyseconds = tonumber(KEYS[3])
local daycount = tonumber(KEYS[4])
local keycount = verify_key .. 'count'
--redis.log(redis.LOG_NOTICE,' keyseconds: '..keyseconds..';daycount:'..daycount)
local current = redis.call('GET', verify_key)
--redis.log(redis.LOG_NOTICE,' current: verify_key:'..current)
if current == false then
    --redis.log(redis.LOG_NOTICE,key..' is nil ')
    local count = redis.call('GET', keycount)
    if count == false then
        redis.call('SET', keycount, 1)
        redis.call('EXPIRE', keycount, 86400)

        redis.call('SET', verify_key, verify_code)
        redis.call('EXPIRE', verify_key, keyseconds)
        return '0'
    else
        local num_count = tonumber(count)
        if num_count + 1 > daycount then
            return '2'
        else
            redis.call('INCRBY', keycount, 1)
            redis.call('SET', verify_key, verify_code)
            redis.call('EXPIRE', verify_key, keyseconds)
            return '0'
        end
    end
else
    --redis.log(redis.LOG_NOTICE,key..' is not nil ')
    return '1'
end
