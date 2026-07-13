package cookbook

import (
	"github.com/golang-cookbook/cookbook/recipes/basics"
	"github.com/golang-cookbook/cookbook/recipes/concurrency"
	"github.com/golang-cookbook/cookbook/recipes/contextrec"
	"github.com/golang-cookbook/cookbook/recipes/embedrec"
	"github.com/golang-cookbook/cookbook/recipes/errors"
	"github.com/golang-cookbook/cookbook/recipes/functions"
	"github.com/golang-cookbook/cookbook/recipes/generics"
	"github.com/golang-cookbook/cookbook/recipes/httprec"
	"github.com/golang-cookbook/cookbook/recipes/interfaces"
	"github.com/golang-cookbook/cookbook/recipes/io"
	"github.com/golang-cookbook/cookbook/recipes/logging"
	"github.com/golang-cookbook/cookbook/recipes/pointers"
	"github.com/golang-cookbook/cookbook/recipes/project"
	"github.com/golang-cookbook/cookbook/recipes/slices"
	"github.com/golang-cookbook/cookbook/recipes/strings"
	"github.com/golang-cookbook/cookbook/recipes/structs"
	"github.com/golang-cookbook/cookbook/recipes/syncrec"
	"github.com/golang-cookbook/cookbook/recipes/testingrec"
	"github.com/golang-cookbook/cookbook/recipes/time"
)

// All returns every registered recipe in stable category order.
func All() []Recipe {
	return []Recipe{
		// Basics
		basics.Variables(),
		basics.ControlFlow(),
		basics.DeferPanicRecover(),

		// Functions
		functions.MultipleReturns(),
		functions.Variadic(),
		functions.Closures(),

		// Structs
		structs.Embedding(),
		structs.MethodReceivers(),

		// Pointers
		pointers.WhenToUse(),

		// Errors
		errors.Wrap(),
		errors.IsAndAs(),
		errors.Join(),
		errors.Sentinel(),
		errors.CustomType(),

		// Interfaces
		interfaces.SmallInterfaces(),
		interfaces.TypeAssertions(),
		interfaces.AcceptReturn(),

		// Slices & maps
		slices.AppendAndCapacity(),
		slices.FilterMap(),
		slices.MapPatterns(),

		// Strings & I/O
		strings.Builder(),
		strings.JSON(),
		io.ReadWriteFile(),
		io.JSONEncodeDecode(),

		// Concurrency
		concurrency.Goroutines(),
		concurrency.Channels(),
		concurrency.Select(),
		concurrency.Semaphore(),
		concurrency.ProducerConsumer(),
		concurrency.ErrGroup(),

		// Context
		contextrec.TimeoutCancel(),
		contextrec.Propagation(),

		// HTTP
		httprec.ServerRouting(),
		httprec.Client(),
		httprec.Middleware(),

		// Testing
		testingrec.TableDriven(),
		testingrec.SubtestsAndParallel(),

		// Generics
		generics.Functions(),
		generics.Constraints(),

		// Logging
		logging.StructuredSlog(),

		// Sync & time
		syncrec.MutexWaitGroup(),
		syncrec.Once(),
		time.TickersTimers(),

		// Project & embed
		project.ModuleLayout(),
		embedrec.EmbedFiles(),
	}
}

// ByID looks up a recipe by its stable identifier.
func ByID(id string) (Recipe, bool) {
	for _, r := range All() {
		if r.ID == id {
			return r, true
		}
	}
	return Recipe{}, false
}

// Categories returns recipes grouped by category name.
func Categories() map[string][]Recipe {
	out := make(map[string][]Recipe)
	for _, r := range All() {
		out[r.Category] = append(out[r.Category], r)
	}
	return out
}
