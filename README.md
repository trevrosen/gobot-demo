## Gobot Demo

[Gobot](http://gobot.io) is "a framework for robotics, computing, and
the Internet of Things, written in the Go programming language."

You can use it to control things like:

* [Sphero](http://www.gosphero.com)
* [BeagleBone](http://beagleboard.org/black)
* [ARDrone](http://ardrone2.parrot.com)

### The abstractions

#### Gobot/Robot
**Gobot** is not only the name of the library, it is also a key type in the lib. It is a logical container for Robots, which are abstractions of devices, connections, and functions. While there are many ways to use Gobot (the library), it really shines when used to control groups of things together. The Gobot type is a struct that basically contains a collection of Robots and of commands (a map of named functions).

**Robots** are tracked by the Gobot type and are what execute commands. Each has a name and is associated with a Connection predicated on an Adaptor. Adaptors are per Platform and are used to map abstract buckets of functionality represented by Drivers to the particular type of hardware in use.

#### Platform
Any given device that is supported directly is a Platform and has an Adaptor. You instantiate an Adaptor type and in order to use a driver with it.

#### Driver
Drivers represent genericized functionality that will be common to multiple platforms. As you can see in the example code, one passes an Adaptor to a Driver so that the unique aspects of the platform can be handled in a decoupled way. For instance: the BeagleBone Adaptor contains the logical representation of the board's pinout, and the GPIO driver knows how to handle the modern Linux notions of GPIO, which are common to all distros using kernel v4.3+

### The REST API
Gobot supports a web-based interface for interaction, data gathering, and running commands. It supports HTTP basic auth for security. There is [a great guide here](http://gobot.io/documentation/guides/api/) - just remember to import the API package as well.
