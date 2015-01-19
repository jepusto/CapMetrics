package gorethink

import p "github.com/scascketta/capmetro-log/Godeps/_workspace/src/github.com/dancannon/gorethink/ql2"

// Transform each element of the sequence by applying the given mapping function.
func (t Term) Map(args ...interface{}) Term {
	return constructMethodTerm(t, "Map", p.Term_MAP, funcWrapArgs(args), map[string]interface{}{})
}

// Takes a sequence of objects and a list of fields. If any objects in the
// sequence don't have all of the specified fields, they're dropped from the
// sequence. The remaining objects have the specified fields plucked out.
// (This is identical to `HasFields` followed by `Pluck` on a sequence.)
func (t Term) WithFields(args ...interface{}) Term {
	return constructMethodTerm(t, "WithFields", p.Term_WITH_FIELDS, args, map[string]interface{}{})
}

// Flattens a sequence of arrays returned by the mapping function into a single
// sequence.
func (t Term) ConcatMap(args ...interface{}) Term {
	return constructMethodTerm(t, "ConcatMap", p.Term_CONCATMAP, funcWrapArgs(args), map[string]interface{}{})
}

type OrderByOpts struct {
	Index interface{} `gorethink:"index,omitempty"`
}

func (o *OrderByOpts) toMap() map[string]interface{} {
	return optArgsToMap(o)
}

// Sort the sequence by document values of the given key(s).
// To specify the index to use for ordering us a last argument in the following form:
//
//	OrderByOpts{Index: "index-name"}
//
// OrderBy defaults to ascending ordering. To explicitly specify the ordering,
// wrap the attribute with either Asc or Desc.
//
//	query.OrderBy("name")
//	query.OrderBy(Asc("name"))
//	query.OrderBy(Desc("name"))
func (t Term) OrderBy(args ...interface{}) Term {
	var opts = map[string]interface{}{}

	// Look for options map
	if len(args) > 0 {
		if possibleOpts, ok := args[len(args)-1].(OrderByOpts); ok {
			opts = possibleOpts.toMap()
			args = args[:len(args)-1]
		}
	}

	for k, arg := range args {
		if t, ok := arg.(Term); !(ok && (t.termType == p.Term_DESC || t.termType == p.Term_ASC)) {
			args[k] = funcWrap(arg)
		}
	}

	return constructMethodTerm(t, "OrderBy", p.Term_ORDERBY, args, opts)
}

func Desc(args ...interface{}) Term {
	return constructRootTerm("Desc", p.Term_DESC, funcWrapArgs(args), map[string]interface{}{})
}

func Asc(args ...interface{}) Term {
	return constructRootTerm("Asc", p.Term_ASC, funcWrapArgs(args), map[string]interface{}{})
}

// Skip a number of elements from the head of the sequence.
func (t Term) Skip(args ...interface{}) Term {
	return constructMethodTerm(t, "Skip", p.Term_SKIP, args, map[string]interface{}{})
}

// End the sequence after the given number of elements.
func (t Term) Limit(args ...interface{}) Term {
	return constructMethodTerm(t, "Limit", p.Term_LIMIT, args, map[string]interface{}{})
}

type SliceOpts struct {
	LeftBound  interface{} `gorethink:"left_bound,omitempty"`
	RightBound interface{} `gorethink:"right_bound,omitempty"`
}

func (o *SliceOpts) toMap() map[string]interface{} {
	return optArgsToMap(o)
}

// Trim the sequence to within the bounds provided.
func (t Term) Slice(args ...interface{}) Term {
	var opts = map[string]interface{}{}

	// Look for options map
	if len(args) > 0 {
		if possibleOpts, ok := args[len(args)-1].(SliceOpts); ok {
			opts = possibleOpts.toMap()
			args = args[:len(args)-1]
		}
	}

	return constructMethodTerm(t, "Slice", p.Term_SLICE, args, opts)
}

// AtIndex gets a single field from an object or the nth element from a sequence.
func (t Term) AtIndex(args ...interface{}) Term {
	return constructMethodTerm(t, "AtIndex", p.Term_BRACKET, args, map[string]interface{}{})
}

// Nth gets the nth element from a sequence.
func (t Term) Nth(args ...interface{}) Term {
	return constructMethodTerm(t, "Nth", p.Term_NTH, args, map[string]interface{}{})
}

// Get the indexes of an element in a sequence. If the argument is a predicate,
// get the indexes of all elements matching it.
func (t Term) IndexesOf(args ...interface{}) Term {
	return constructMethodTerm(t, "IndexesOf", p.Term_INDEXES_OF, funcWrapArgs(args), map[string]interface{}{})
}

// Test if a sequence is empty.
func (t Term) IsEmpty(args ...interface{}) Term {
	return constructMethodTerm(t, "IsEmpty", p.Term_IS_EMPTY, args, map[string]interface{}{})
}

// Concatenate two sequences.
func (t Term) Union(args ...interface{}) Term {
	return constructMethodTerm(t, "Union", p.Term_UNION, args, map[string]interface{}{})
}

// Select a given number of elements from a sequence with uniform random
// distribution. Selection is done without replacement.
func (t Term) Sample(args ...interface{}) Term {
	return constructMethodTerm(t, "Sample", p.Term_SAMPLE, args, map[string]interface{}{})
}
