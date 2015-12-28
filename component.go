package component

// System ...
type System struct {
	components map[string]*Component
}

// Component ...
type Component struct {
	name         string
	constructor  func(args ...interface{}) Lifecycle
	args         []interface{}
	dependencies []string
	entity       Lifecycle
}

// Lifecycle ...
type Lifecycle interface {
	Start(dependencies ...interface{}) error
	Stop() error
}

// NewSystem ...
func NewSystem() *System {
	components := make(map[string]*Component)
	system := System{
		components: components,
	}

	return &system
}

// NewComponent ...
func (s *System) NewComponent(name string) *Component {
	comp := Component{
		name: name,
	}
	s.components[name] = &comp

	return &comp
}

// Constructor ...
func (c *Component) Constructor(f func(args ...interface{}) Lifecycle) *Component {
	c.constructor = f

	return c
}

// Args ...
func (c *Component) Args(args ...interface{}) *Component {
	c.args = args

	return c
}

// Dependencies ...
func (c *Component) Dependencies(args ...string) *Component {
	c.dependencies = args

	return c
}

// Start ...
func (s *System) Start() error {
	graph := NewGraph()
	for dep, component := range s.components {
		for _, d := range component.dependencies {
			graph.AddEdge(dep, d)
		}
	}
	vertices, err := (*graph).TopologicalSort()
	if err != nil {
		return err
	}
	for _, vertice := range vertices {
		component := s.components[vertice]
		entity := component.constructor(component.args...)
		dependencies := make([]interface{}, len(component.dependencies))
		for i, dep := range component.dependencies {
			dependencies[i] = s.components[dep].entity
		}
		entity.Start(dependencies...)
		component.entity = entity
	}

	return nil
}

// Stop ...
func (s *System) Stop() error {
	for _, c := range s.components {
		err := c.entity.Stop()
		if err != nil {
			return err
		}
	}
	return nil
}
