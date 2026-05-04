local socket = require("socket")

local server = assert(socket.bind("127.0.0.1", 9788))
local _, port = server:getsockname()

local clients = {}
local listeners = {server}

local messages = {}

print("server running on port " .. port)

while true do
  local ready, _, err = socket.select(listeners, nil, 1)
  if err then
    goto continue
  end
  for _, input in ipairs(ready) do
    if input == server then
      local client = server:accept()
      if client then
        client:settimeout(10)
        table.insert(listeners, client)
        print("connection")
      end
    else
      local line, err = input:receive()
      if err then
        input:close()
        print("disconnection")
      else
        table.insert(messages, line)
      end
    end
  end
  ::continue::
end
