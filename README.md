# Component

Dependency Injection for golang.

Inspired by [stuartsierra/component](https://github.com/stuartsierra/component).
It is hardly be as elegant as in Clojure though.

It is type unsafe and I don't know how it could be fixed with current golang state.

# Usage

```golang
package main

import component
// imaginable postgresql driver
import pg

type Database {
  host string
  port int
  database string
  connection *pg.connection
}

type Scheduler {
  threadPoolSize int
}

type App {
  logLevel string
}

func (d *Database) NewDatabase(args ...interface{}) component.Lifecicle {
  return &Database{
    host: args[0].(string),
    port: args[1].(int),
    database: args[2].(string),
  }
}

func (d *Database) Start(dependencies ...interface{}) error {
  d.connection, err := pg.Connect(d.host, d.port, d.database)
  return err
}

func (d *Database) Stop() error {
  err := d.connection.Close()
  d.connection = nil
  return err
}

// ...
// same for Scheduler and App types
//

func main() {
  system = component.New()

  system.
    NewComponent("database").
    Constructor(NewDatabase).
    Args("localhost", 5432, "dev_database")

  system.
    NewComponent("scheduler").
    Constructor(NewScheduler).
    Args(12).
    Dependencies("database")

  system.
    NewComponent("app").
    Constructor(NewApp).
    Args("Error").
    Dependencies("database", "scheduler")

  system.Start()

  defer system.Stop()
}
```
