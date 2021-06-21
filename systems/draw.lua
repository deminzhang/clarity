local system = require('library.system')
local viewport = require('utilities.viewport')

local graphics = love.graphics
local renderWidth = love.c.renderWidth
local renderHeight = love.c.renderHeight

return function(state)
    local s = system({'appearance', 'position'})

    local cameraCanvas
    local uiCanvas
    local translateX = 0
    local translateY = 0
    local windowScale = 1

    local updateTransform = function()
        local w, h, scale = viewport.getWindowData()
        -- Floor to even numbers
        local width = w - (w % 2)
        local height = h - (h % 2)
        translateX = (width - renderWidth * scale) * 0.5
        translateY = (height - renderHeight * scale) * 0.5
        windowScale = scale
    end

    s.load = function(arg)
        updateTransform()
        cameraCanvas = graphics.newCanvas(renderWidth, renderHeight)
        uiCanvas = graphics.newCanvas(renderWidth, renderHeight)
    end

    s.draw = function()

        -- Clear canvases
        graphics.setCanvas(cameraCanvas)
        graphics.clear()
        graphics.setCanvas(uiCanvas)
        graphics.clear()
        graphics.setCanvas()

        for index = 1, #s.entities do
            local e = s.entities[index]

            local canvas = e.camera and cameraCanvas or uiCanvas
            graphics.setCanvas(canvas)
            if e.animation then
                graphics.draw(
                    e.appearance.image,
                    e.animation.getFrame(),
                    e.position.x,
                    e.position.y
                )
            elseif e.appearance.image then
                graphics.draw(
                    e.appearance.image,
                    e.appearance.frame,
                    e.position.x,
                    e.position.y
                )
            end
            if e.text then
                local width = e.appearance.width
                if e.animation then
                    local _, _, w, _ = e.animation.getFrame():getViewport()
                    width = w
                end
                graphics.printf(
                    e.text.content,
                    e.position.x + e.text.padding,
                    e.position.y + e.text.padding,
                    width - (2 * e.text.padding),
                    e.text.align
                )
            end
            graphics.setCanvas()
        end

        -- Begin rendering
        graphics.push()
        graphics.translate(translateX, translateY)
        graphics.scale(windowScale)

        -- Avoid incorrect colours https://love2d.org/forums/viewtopic.php?f=4&p=211418#p211418
        graphics.setColor(1, 1, 1, 1)
        graphics.setBlendMode('alpha', 'premultiplied')
        graphics.draw(cameraCanvas)
        graphics.draw(uiCanvas)
        graphics.setBlendMode('alpha')

        -- End rendering
        graphics.pop()
    end

    s.resize = function()
        updateTransform()
    end

    return s
end