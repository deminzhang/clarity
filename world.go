package main

type world struct {
	name     string
	systems  []runner
	entities []*entity
}

func (w *world) run() {
	for _, system := range w.systems {
		system.run()
	}
}

func (w *world) addEntity(e *entity) {
	w.entities = append(w.entities, e)
}

func (w *world) addSystem(s runner) {
	w.systems = append(w.systems, s)
}

func (w *world) sortEntities() {
	for _, system := range w.systems {
		system.sortEntities()
	}
}

// Update all systems with the entities they will be run on
func (w *world) updateSystems() {
	hasComponents := true
	for _, system := range w.systems {
		for _, entity := range w.entities {
			for _, component := range system.getComponents() {
				switch component {
				case "appearance":
					if entity.appearance == nil {
						hasComponents = false
					}
				case "ui":
					if entity.ui == nil {
						hasComponents = false
					}
				case "camera":
					if entity.camera == nil {
						hasComponents = false
					}
				case "collision":
					if entity.collision == nil {
						hasComponents = false
					}
				case "controls":
					if entity.controls == nil {
						hasComponents = false
					}
				case "velocity":
					if entity.velocity == nil {
						hasComponents = false
					}
				case "gravity":
					if entity.gravity == nil {
						hasComponents = false
					}
				case "position":
					if entity.position == nil {
						hasComponents = false
					}
				}
			}
			if hasComponents {
				system.addEntity(entity)
			}
			hasComponents = true
		}
	}
}

func newMainWorld() *world {
	w := &world{name: "main"}

	player := newPlayerEntity()
	draw := newDrawSystem().(*drawSystem)
	draw.cameraTarget = player
	w.addSystem(draw)
	w.addSystem(newPlayerSytem())
	w.addSystem(newUISystem())
	w.addEntity(player)
	w.addEntity(newBoxEntity())

	w.updateSystems()
	w.sortEntities()
	return w
}
