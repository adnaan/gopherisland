package main

import (
	"fmt"
	"time"
)

type Vec2 struct {
	x int
	y int
}

/*interfaces*/

type Movement interface {
	Pos() Vec2
	Velocity() float64
}

type Health interface {
	CurrHealth() int
	MaxHealth() int
}

type Render interface {
	Drawable() string
}

type Component interface {
	Update(float64)
}

type Entity interface {
	UpdateEntity(float64)
	GetComponents() []interface{}
}

/*game components*/

/*movement*/

type MovementComponent struct {
	pos Vec2
	v   float64
}

func (m *MovementComponent) Pos() Vec2 {
	return m.pos

}

func (m *MovementComponent) Velocity() float64 {

	return m.v

}

func (m *MovementComponent) Update(delta float64) {

}

/*health*/

type HealthComponent struct {
	currHealth int
	maxHealth  int
}

func (h *HealthComponent) CurrHealth() int {
	return h.currHealth

}

func (h *HealthComponent) MaxHealth() int {
	return h.maxHealth

}

func (h *HealthComponent) Update(delta float64) {

}

/*render*/

type RenderComponent struct {
	image string //for demo
}

func (r *RenderComponent) Drawable() string {

	return r.image

}

func (r *RenderComponent) Update(delta float64) {

}

/*Entities*/
/*Player Entity*/

type Player struct {
	score      int
	components []interface{}
}

func (p *Player) add(c Component) {
	p.components = append(p.components, c)

}

func (p *Player) UpdateEntity(delta float64) {
	for i := range p.components {
		p.components[i].(Component).Update(delta)
	}

}

func (p *Player) GetComponents() []interface{} {

	return p.components
}

/*Enemy Entity*/
type Enemy struct {
	components []interface{}
}

func (e *Enemy) add(c Component) {
	e.components = append(e.components, c)

}

func (e *Enemy) UpdateEntity(delta float64) {
	for i := range e.components {
		e.components[i].(Component).Update(delta)
	}

}

func (e *Enemy) GetComponents() []interface{} {
	return e.components

}

type World struct {
	size     int
	entities []interface{}
	grid     [][]string
	endGame  bool
}

func (w *World) create() {

	N := w.size

	for i := 0; i < N; i++ {
		row := make([]string, N*2)
		w.grid = append(w.grid, row)

	}

	for y := 0; y < N; y++ {
		for x := 0; x < N*2; x++ {
			w.grid[y][x] = "*" //fill the world with cookies

		}
	}

	w.renderWorld("Created World!")

}

func (w *World) renderWorld(msg string) {
	fmt.Printf("%v...\n\n", msg)
	for i := range w.grid {
		fmt.Printf("%v\n", w.grid[i])
	}
	fmt.Printf(".....................................\n\n")

}

func (w *World) start() {
	for i := range w.entities {

		for j := range w.entities[i].(Entity).GetComponents() {

			if v, ok := w.entities[i].(Entity).GetComponents()[j].(*MovementComponent); ok {

				if _, ok := w.entities[i].(*Player); ok {
					//if player
					w.grid[v.Pos().y][v.Pos().x] = "p"
				} else {
					//if enemy
					w.grid[v.Pos().y][v.Pos().x] = "e"
				}

			}

		}

	}

	w.renderWorld("Game Initialized!")

}

func (w *World) add(et Entity) {
	w.entities = append(w.entities, et)

}

func (w *World) update(delta float64) {

	for i := range w.entities {

		w.entities[i].(Entity).UpdateEntity(delta)
	}

	//get data from the components and render the world

	w.renderWorld("Updated Game!")

	//apply ai or input
	//check collision with cookie, enemy
	//if cookie: +score, hide cookie with "-"
	//if enemy: -health
	//render

}

func (w *World) isEndGame() bool {

	return w.endGame

}

func main() {

	//create world
	w := World{size: 9}
	w.create()

	//create player
	p := Player{}
	p.add(&MovementComponent{Vec2{0, 0}, 2})
	p.add(&HealthComponent{10, 10})
	p.add(&RenderComponent{"p"})

	//create enemy
	e := Enemy{}
	e.add(&MovementComponent{Vec2{17, 8}, 2})
	e.add(&HealthComponent{1, 1})
	e.add(&RenderComponent{"e"})

	//add player and enemy to world
	w.add(&p)
	w.add(&e)

	w.start()

	for {
		time.Sleep(500 * time.Millisecond)
		w.update(1)
		if w.isEndGame() {
			break
		}

	}

}
