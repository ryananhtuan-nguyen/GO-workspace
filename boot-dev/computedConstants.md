COMPUTED CONSTANTS
Constants must be known at compile time. More often than not they will be declared with a static value:

const myInt = 15
Copy icon
However, constants can be computed so long as the computation can happen at compile time.

For example, this is valid:

const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName
Copy icon
That said, you cannot declare a constant that can only be computed at run-time.

ASSIGNMENT
Keeping track of time in a message-sending application is critical. Imagine getting an appointment reminder an hour after your doctor's visit.

Complete the code using a computed constant to print the number of seconds in an hour.