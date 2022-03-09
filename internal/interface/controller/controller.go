package controller

type Controller struct {
	User interface{ User }
	Post interface{ Post }
}
