# gotorch

### A learning project!

## Version 1:

ability to create a basic neural network for classification on MNIST digits

things i need:
- data management
  - tensor object
  - split train test
- activation functions -- relu for all hidden and softmax at output
- model object
  - hidden layers
    - dense
    - activation
    - dropout
  - optimiser
  - loss -- categorical crossentropy
  - optimisation -- batch gradient descent
  - metric -- accuracy
  - training parameters
    - batch size
    - learning rate
    - epochs


# Go Notes

> Arguments are always copied in Go. If we attempted to modify one of the fields inside of the circleArea function, it would not modify the original variable. Because of this we would typically write the function like this:

```golang
c := Circle{0, 0, 5}
func circleArea(c *Circle) float64 {
  return math.Pi * c.r*c.r
}
```

And when calling the function we use

```golang
fmt.Println(circleArea(&c))
```

Instead of using a function we can use a method

```golang
func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}
```

> In between the keyword func and the name of the function we've added a “receiver”. The receiver is like a parameter – it has a name and a type – but by creating the function in this way it allows us to call the function using the . operator:

`fmt.Println(c.area())`

> This is much easier to read, we no longer need the & operator (Go automatically knows to pass a pointer to the circle for this method) and because this function can only be used with Circles we can rename the function to just area.


- Great Go resource: https://www.golang-book.com/books/intro
- Linear Layer Torch: https://pytorch.org/docs/stable/_modules/torch/nn/modules/linear.html#Linear
