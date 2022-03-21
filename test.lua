-- file: lua/backend-baidu.lua

local http = require 'http'
local backend = require 'backend'

local char = string.char
local byte = string.byte
local find = string.find
local sub = string.sub

local ADDRESS = backend.ADDRESS
local PROXY = backend.PROXY
local DIRECT_WRITE = backend.SUPPORT.DIRECT_WRITE

local SUCCESS = backend.RESULT.SUCCESS
local HANDSHAKE = backend.RESULT.HANDSHAKE
local DIRECT = backend.RESULT.DIRECT

local ctx_uuid = backend.get_uuid
local ctx_proxy_type = backend.get_proxy_type
local ctx_address_type = backend.get_address_type
local ctx_address_host = backend.get_address_host
local ctx_address_bytes = backend.get_address_bytes
local ctx_address_port = backend.get_address_port
local ctx_write = backend.write
local ctx_free = backend.free
local ctx_debug = backend.debug

local flags = {}
local kHttpHeaderSent = 1
local kHttpHeaderRecived = 2

function wa_lua_on_flags_cb(ctx)
    return DIRECT_WRITE
end

function wa_lua_on_handshake_cb(ctx)
    local uuid = ctx_uuid(ctx)

    if flags[uuid] == kHttpHeaderRecived then
        return true
    end

    if flags[uuid] ~= kHttpHeaderSent then
        local host = ctx_address_host(ctx)
        local port = ctx_address_port(ctx)
        local res = 'CONNECT ' .. host .. ':' .. port .. ' HTTP/1.1\r\n' ..
                    'Host: ' .. host .. '\r\n' ..
                    'Host: 113.200.89.85:809/fifa/?rrsip=20.24.98.39&videoid=18080196&PlayType=vod&apptype=app&spid=32121&pid=8031006300&preview=1&portalid=500&videoname=MDA377ya5peg5pqH6LW05q275Y6f5aOw54mI&userid=21569963447&userip=2408:856a:8110:4f0:dba9:658c:8167:7779&spip=20.24.98.39&spport=8080&freetag=1&ugpid=968&tradeid=21e32f188a2c465c80b254063394e4d1&lsttm=20220321020002&enkey=3b4f696b35dfd3311c651dcf83ad1124\r\n' ..
                    'Proxy-Connection: Keep-Alive\r\n'..
                    'X-T5-Auth: 9511095110\r\nUser-Agent: Dalvik/2.1.0\r\nHost: tms.dingtalk.com\r\n\r\n'
        ctx_write(ctx, res)
        flags[uuid] = kHttpHeaderSent
    end

    return false
end

function wa_lua_on_read_cb(ctx, buf)
    ctx_debug('wa_lua_on_read_cb')
    local uuid = ctx_uuid(ctx)
    if flags[uuid] == kHttpHeaderSent then
        flags[uuid] = kHttpHeaderRecived
        return HANDSHAKE, nil
    end
    return DIRECT, buf
end

function wa_lua_on_write_cb(ctx, buf)
    ctx_debug('wa_lua_on_write_cb')
    return DIRECT, buf
end

function wa_lua_on_close_cb(ctx)
    ctx_debug('wa_lua_on_close_cb')
    local uuid = ctx_uuid(ctx)
    flags[uuid] = nil
    ctx_free(ctx)
    return SUCCESS
end
