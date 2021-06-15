return function(name)
    local w = {}
    w.name = name
    w.systems = {}
    w.entities = {}

    w.addEntity = function(entity)
        w.entities[#w.entities + 1] = entity
    end

    w.addSystem = function(system)
        w.systems[#w.systems + 1] = system
    end

    w.sortEntities = function()
        for index = 1, #w.systems do
            w.systems[index].sortEntities()
        end
    end

    w.updateSystems = function()
        local system
        local entity
        local requiredComponent
        local component
        local hasComponent = false
        local hasAllComponents = false
        for systemIndex = 1, #w.systems do
            system = w.systems[systemIndex]
            for entityIndex = 1, #w.entities do
                entity = w.entities[entityIndex]
                for requiredComponentIndex = 1, #system.components do
                    requiredComponent = system.components[requiredComponentIndex]
                    if entity[requiredComponent] then
                        hasComponent = true
                    end
                    hasAllComponents = hasComponent
                    hasComponent = false
                end
                if hasAllComponents then
                    system.addEntity(entity)
                end
                hasAllComponents = false
            end
        end
    end

    w.load = function(arg)
        for index = 1, #w.systems do
            w.systems[index].load(arg)
        end
    end

    w.update = function(dt)
        for index = 1, #w.systems do
            w.systems[index].update(dt)
        end
    end

    w.draw = function()
        for index = 1, #w.systems do
            w.systems[index].draw()
        end
    end

    w.keypressed = function(key)
        for index = 1, #w.systems do
            w.systems[index].keypressed(key)
        end
    end

    w.keyreleased = function(key)
        for index = 1, #w.systems do
            w.systems[index].keyreleased(key)
        end
    end

    w.resize = function(width, height)
        for index = 1, #w.systems do
            w.systems[index].resize(width, height)
        end
    end

    return w
end
