package controller

// Controller holds the available controllers of the entire application.
type Controller struct {
	Slide    interface{ Slide }
	User     interface{ User }
	Instance interface{ Instance }
	Link     interface{ Link }
}
