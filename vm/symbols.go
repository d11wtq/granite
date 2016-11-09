package vm

// Storage for the mapping of names to registers
type SymbolTable struct {
	// Scopes are nested as linked lists
	scope *scope
}

// Create a new symbol table.
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{newScope(nil)}
}

// Locate the register for a given name.
// If it does not exist in the current scope, lookup traverses into the parent.
// If the name does not exist at any level, the second return value is false.
func (s *SymbolTable) Get(name string) (Operand, bool) {
	return s.scope.get(name)
}

// Set the register for a given name.
// This only affects the current scope.
// It isimpossible to mutate the parent scope.
func (s *SymbolTable) Set(name string, v Operand) {
	s.scope.set(name, v)
}

// Start a new scope beneath the current scope.
func (s *SymbolTable) BeginScope() {
	s.scope = newScope(s.scope)
}

// End the current scope and return to the parent scope.
func (s *SymbolTable) EndScope() {
	if s.scope.parent == nil {
		panic("attempted to terminate the root scope")
	}

	s.scope = s.scope.parent
}

// Scope is the underlying storage.
// Scopes are nested using a linked list, where the head is the current scope.
type scope struct {
	parent  *scope
	symbols map[string]Operand
}

func newScope(parent *scope) *scope {
	return &scope{parent, make(map[string]Operand)}
}

func (s *scope) get(name string) (Operand, bool) {
	v, exists := s.symbols[name]
	if !exists && s.parent != nil {
		return s.parent.get(name)
	}
	return v, exists
}

func (s *scope) set(name string, v Operand) {
	s.symbols[name] = v
}
