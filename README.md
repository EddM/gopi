# gopi

I'm learning [go](https://golang.org) with a view to making a game with it one day.

So I thought a good way to get to know go and its game development ecosystem was to create an OpenGL-based visualisation of a Monte Carlo method for approximating pi.

This program imagines a square, and a circle inside the square with a diameter equal to the square's width. It plots random points inside the square. The ratio of points inside the circle to the total number of points is a very rough approximation of `pi/4`.

In addition, this program renders a visualisation of this Monte Carlo simulation.

![Screenshot](https://i.imgur.com/UeDeVJa.png)
