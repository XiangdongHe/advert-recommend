-- KEYS 空，全部用 ARGV
-- ARGV = {tag1, tag2, ..., tagN}

local result = {}

-- Step 1: 获取所有 planID
local planIDSet = {}
for i, tag in ipairs(ARGV) do
    local key = "interest:" .. tag
    local planIDs = redis.call("SMEMBERS", key)
    for _, pid in ipairs(planIDs) do
        planIDSet[pid] = true
    end
end

-- Step 2: 获取 plan JSON（每个 plan JSON 已包含 creatives）
for pid, _ in pairs(planIDSet) do
    local planJSON = redis.call("GET", "ad:plan:" .. pid)
    if planJSON then
        table.insert(result, planJSON)
    end
end

return result
