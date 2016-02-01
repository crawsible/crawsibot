package mocks

import "reflect"

type FakeFlagSet struct {
	StringVarCalls int
	DefinedFlags   [][]string
	BoundVarPtrs   map[string]uintptr

	ParseCalls            int
	ParseArgs             []string
	StringVarCallsAtParse int
}

func NewFakeFlagSet() *FakeFlagSet {
	return &FakeFlagSet{
		BoundVarPtrs: make(map[string]uintptr),
	}
}

func (f *FakeFlagSet) StringVar(p *string, name, value, usage string) {
	f.DefinedFlags = append(f.DefinedFlags, []string{name, value, usage})
	f.BoundVarPtrs[name] = reflect.ValueOf(p).Pointer()
	f.StringVarCalls += 1
}

func (f *FakeFlagSet) Parse(args []string) error {
	f.ParseCalls += 1
	f.ParseArgs = args
	f.StringVarCallsAtParse = f.StringVarCalls

	return nil
}
