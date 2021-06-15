return function()
    local s = {}
    s.worlds = {}
    s.activateWorld = nil
    s.controls = {}

    local keyMap = {
        ['space'] = 'jump',
        ['up'] = 'up',
        ['w'] = 'up',
        ['left'] = 'left',
        ['a'] = 'left',
        ['right'] = 'right',
        ['d'] = 'right',
        ['down'] = 'down',
        ['s'] = 'down',
        ['escape'] = 'menu',
        ['`'] = 'debug'
    }
    for _, control in pairs(keyMap) do
        s.controls[control] = false
    end

    s.addWorld = function(world)
        s.worlds[world.name] = world
    end

    s.activateWorld = function(worldName)
        s.activateWorld = s.worlds[worldName]
    end

    s.load = function(arg)
        s.activateWorld.load(arg)
    end

    s.update = function(dt)
        s.activateWorld.update(dt)
    end

    s.draw = function()
        s.activateWorld.draw()
    end

    s.keypressed = function(key)
        local control = keyMap[key]
        if control then
            s.controls[control] = true
        end
        s.activateWorld.keypressed(key)
    end

    s.keyreleased = function(key)
        local control = keyMap[key]
        if control then
            s.controls[control] = false
        end
        s.activateWorld.keyreleased(key)
    end

    s.resize = function(w, h)
        s.activateWorld.resize(w, h)
    end

    return s
end