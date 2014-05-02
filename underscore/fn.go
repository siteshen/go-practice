//     Underscore.js 1.6.0
//     http://underscorejs.org
//     (c) 2009-2014 Jeremy Ashkenas, DocumentCloud and Investigative Reporters & Editors
//     Underscore may be freely distributed under the MIT license.
package underscore

import "errors"

type (
	IntSlice    []int
	IntEachFn   func(int)
	IntMapFn    func(int) int
	IntReduceFn func(int, int) int
	IntFilterFn func(int) bool
)

// Collection Functions
// --------------------

// The cornerstone, an `each` implementation, aka `forEach`.
//
// var each = _.each = _.forEach = function(obj, iterator, context) {
func (self IntSlice) Each(fn IntEachFn) {
	for _, item := range self {
		fn(item)
	}
}

// Return the results of applying the iterator to each element.
//
// _.map = _.collect = function(obj, iterator, context) {
func (self IntSlice) Map(fn IntMapFn) (result IntSlice) {
	for _, item := range self {
		result = append(result, fn(item))
	}
	return
}

// Reduce builds up a single result from a list of values.
//
// _.reduce = _.foldl = _.inject = function(obj, iterator, memo, context) {
func (self IntSlice) Reduce(fn IntReduceFn, initial int) (result int) {
	result = initial
	for _, item := range self {
		result = fn(result, item)
	}
	return
}

// The right-associative version of reduce, also known as `foldr`.
//
// _.reduceRight = _.foldr = function(obj, iterator, memo, context) {

// Return the first value which passes a truth test. Aliased as `detect`.
//
// _.find = _.detect = function(obj, predicate, context) {
func (self IntSlice) Find(fn IntFilterFn) (int, error) {
	for _, item := range self {
		if fn(item) {
			return item, nil
		}
	}
	return 0, errors.New("element not found")
}

// Return all the elements that pass a truth test.
//
// _.filter = _.select = function(obj, predicate, context) {
func (self IntSlice) Filter(fn IntFilterFn) (result IntSlice) {
	for _, item := range self {
		if fn(item) {
			result = append(result, item)
		}
	}
	return
}

// Return all the elements for which a truth test fails.
//
// _.reject = function(obj, predicate, context) {
func (self IntSlice) Reject(fn IntFilterFn) (result IntSlice) {
	for _, item := range self {
		if !fn(item) {
			result = append(result, item)
		}
	}
	return
}

// Determine whether all of the elements match a truth test.
//
// _.every = _.all = function(obj, predicate, context) {
func (self IntSlice) Every(fn IntFilterFn) bool {
	for _, item := range self {
		if !fn(item) {
			return false
		}
	}
	return true
}

// Determine if at least one element in the object matches a truth test.
//
// var any = _.some = _.any = function(obj, predicate, context) {
func (self IntSlice) Some(fn IntFilterFn) bool {
	for _, item := range self {
		if fn(item) {
			return true
		}
	}
	return false
}

// Determine if the array or object contains a given value.
//
// _.contains = _.include = function(obj, target) {
func (self IntSlice) Contains(item int) bool {
	return self.Some(func(x int) bool { return x == item })
}

// Invoke a method (with arguments) on every item in a collection.
// _.invoke = function(obj, method) {

// Convenience version of a common use case of `map`: fetching a property.
// _.pluck = function(obj, key) {

// Convenience version of a common use case of `filter`: selecting only objects
// containing specific `key:value` pairs.
// _.where = function(obj, attrs) {

// Convenience version of a common use case of `find`: getting the first object
// containing specific `key:value` pairs.
// _.findWhere = function(obj, attrs) {

// Return the maximum element or (element-based computation).
//
// _.max = function(obj, iterator, context) {
func (self IntSlice) Max() int {
	if len(self) == 0 {
		panic("empty")
	}
	return self.Reduce(max, self[0])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Return the minimum element (or element-based computation).
//
// _.min = function(obj, iterator, context) {
func (self IntSlice) Min() int {
	if len(self) == 0 {
		panic("empty")
	}
	return self.Reduce(min, self[0])
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Shuffle an array, using the modern version of the
// [Fisher-Yates shuffle](http://en.wikipedia.org/wiki/Fisher–Yates_shuffle).
//
// _.shuffle = function(obj) {

// Sample N random values from a collection.
//
// _.sample = function(obj, n, guard) {

// Sort the object's values by a criterion produced by an iterator.
//
// _.sortBy = function(obj, iterator, context) {

// An internal function used for aggregate "group by" operations.
//
// var group = function(behavior) {

// Groups the object's values by a criterion. Pass either a string attribute
// to group by, or a function that returns the criterion.
//
// _.groupBy = group(function(result, key, value) {

// Indexes the object's values by a criterion, similar to `groupBy`, but for
// when you know that your index values will be unique.
//
// _.indexBy = group(function(result, key, value) {

// Counts instances of an object that group by a certain criterion. Pass
// either a string attribute to count by, or a function that returns the
// criterion.
//
// _.countBy = group(function(result, key) {

// Use a comparator function to figure out the smallest index at which
// an object should be inserted so as to maintain order. Uses binary search.
//
// _.sortedIndex = function(array, obj, iterator, context) {

// Safely create a real, live array from anything iterable.
// _.toArray = function(obj) {

// Return the number of elements in an object.
//
// _.size = function(obj) {
func (self IntSlice) Size() int { return len(self) }

// Array Functions
// ---------------

// Get the first N elements in the array.
//
// _.first = _.head = _.take = function(array, n, guard) {
func (self IntSlice) Take(n int) IntSlice { return self[:n] }

// Returns all the values in the array, excluding the last N.
//
// _.initial = function(array, n, guard) {
func (self IntSlice) Initial(n int) IntSlice { return self[:len(self)-n] }

// Get the last N elements in the array.
//
// _.last = function(array, n, guard) {
func (self IntSlice) Last(n int) IntSlice { return self[len(self)-n : len(self)] }

// Returns all the values in the array, excluding the first N.
//
// _.rest = _.tail = _.drop = function(array, n, guard) {
func (self IntSlice) Rest(n int) IntSlice { return self[n:] }

// Trim out all falsy values from an array.
//
// _.compact = function(array) {
func (self IntSlice) Compact() IntSlice {
	return self.Filter(func(item int) bool { return item != 0 })
}

// Internal implementation of a recursive `flatten` function.
//
// var flatten = function(input, shallow, output) {

// Flatten out an array, either recursively (by default), or just one level.
//
// _.flatten = function(array, shallow) {

// Return a version of the array that does not contain the specified value(s).
//
// _.without = function(array) {

// Split an array into two arrays: one whose elements all satisfy the given
// predicate, and one whose elements all do not satisfy the predicate.
//
// _.partition = function(array, predicate) {

// Produce a duplicate-free version of the array. If the array has already
// been sorted, you have the option of using a faster algorithm.
// Aliased as `unique`.
//
// _.uniq = _.unique = function(array, isSorted, iterator, context) {

// Produce an array that contains the union: each distinct element from all of
// the passed-in arrays.
//
// _.union = function() {

// Produce an array that contains every item shared between all the
// passed-in arrays.
// _.intersection = function(array) {

// Take the difference between one array and a number of other arrays.
// Only the elements present in just the first array will remain.
// _.difference = function(array) {

// Zip together multiple lists into a single array -- elements that share
// an index go together.
// _.zip = function() {

// Converts lists into objects. Pass either a single array of `[key, value]`
// pairs, or two parallel arrays of the same length -- one of keys, and one of
// the corresponding values.
// _.object = function(list, values) {

// Return the position of the first occurrence of an item in an array, or -1 if
// the item is not included in the array.
//
// _.indexOf = function(array, item, isSorted) {
func (self IntSlice) IndexOf(item int) int {
	for index, val := range self {
		if val == item {
			return index
		}
	}
	return -1
}

// Return the position of the last occurrence of an item in an array, or -1 if
// the item is not included in the array.
//
// _.lastIndexOf = function(array, item, from) {
func (self IntSlice) LastIndexOf(item int) int {
	length := self.Size()
	for idx := length - 1; idx >= 0; idx-- {
		if self[idx] == item {
			return idx
		}
	}
	return -1
}

// Generate an integer Array containing an arithmetic progression.
//
// A port of the native Python `range()` function. See [the Python
// documentation](http://docs.python.org/library/functions.html#range).
//
// _.range = function(start, stop, step) {
func Range(args ...int) (result IntSlice) {
	var start, stop, step int
	switch len(args) {
	case 1:
		start, stop, step = 0, args[0], 1
	case 2:
		start, stop, step = args[0], args[1], 1
	case 3:
		start, stop, step = args[0], args[1], args[2]
	default:
		return nil
	}
	// poolman's ceil
	length := (stop - start + step - 1) / step
	for i := 0; i < length; i++ {
		result = append(result, start+i*step)
	}

	return
}

// Customize functions
func (self IntSlice) Equal(other IntSlice) bool {
	// check length
	if len(self) != len(other) {
		return false
	}

	// check all value
	for index := range self {
		if self[index] != other[index] {
			return false
		}
	}

	return true
}
