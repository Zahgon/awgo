// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package aw

import (
	"go.deanishe.net/fuzzy"
)

// Valid modifier keys pressed by the user to run an alternate
// item action in Script Filters (in combination with ↩).
// Passed to Item.NewModifier().
//
// Alfred 3 only permits one modifier at a time, but in Alfred 4+
// you can combine them arbitrarily.
const (
	ModCmd   string = "cmd"   // Alternate action for ⌘↩
	ModAlt   string = "alt"   // Alternate action for ⌥↩
	ModOpt   string = "alt"   // Synonym for ModAlt
	ModCtrl  string = "ctrl"  // Alternate action for ^↩
	ModShift string = "shift" // Alternate action for ⇧↩
	ModFn    string = "fn"    // Alternate action for fn↩
)

// Types understood by Alfred's `action` API call and item field. Added in Alfred 4.5.
const (
	TypeFile = "file" // values are paths
	TypeURL  = "url"  // values are URLs
	TypeText = "text" // values are just text
)

// Item is a single Alfred Script Filter result.
// Together with Feedback & Modifier, Item generates Script Filter feedback
// for Alfred.
//
// Create Items via NewItem(), so they are bound to their parent Feedback.
type Item struct {
	title        string
	subtitle     *string
	match        *string
	uid          *string
	autocomplete *string
	arg          []string
	valid        bool
	file         bool
	copytext     *string
	largetype    *string
	ql           *string
	vars         map[string]string
	mods         map[string]*Modifier
	actions      map[string][]string
	icon         *Icon
	noUID        bool // Suppress UID in JSON
}

// Title sets the title of the item in Alfred's results.
func (it *Item) Title(s string) *Item { _ = "STUB: not implemented"; return nil }

// Subtitle sets the subtitle of the item in Alfred's results.
func (it *Item) Subtitle(s string) *Item { _ = "STUB: not implemented"; return nil }

// Match sets Item's match field for filtering.
// If present, this field is preferred over the item's title for fuzzy sorting
// via Feedback, and by Alfred's "Alfred filters results" feature.
func (it *Item) Match(s string) *Item { _ = "STUB: not implemented"; return nil }

// Arg sets Item's arg, the value(s) passed as {query} to the next workflow action.
// Multiple values are allowed in Alfred 4.1 and later.
func (it *Item) Arg(s ...string) *Item { _ = "STUB: not implemented"; return nil }

// UID sets Item's unique ID, which is used by Alfred to remember your choices.
// Use a blank string to force results to appear in the order you add them.
//
// You can also use the SuppressUIDs() Option to (temporarily) suppress output of UIDs.
func (it *Item) UID(s string) *Item { _ = "STUB: not implemented"; return nil }

// Autocomplete sets what Alfred's query expands to when the user TABs result.
// (or hits RETURN on a result where valid is false)
func (it *Item) Autocomplete(s string) *Item { _ = "STUB: not implemented"; return nil }

// Valid tells Alfred whether the result is "actionable", i.e. ENTER will
// pass Arg to subsequent action.
func (it *Item) Valid(b bool) *Item { _ = "STUB: not implemented"; return nil }

// IsFile tells Alfred that this Item is a file, i.e. Arg is a path
// and Alfred's File Actions should be made available.
func (it *Item) IsFile(b bool) *Item { _ = "STUB: not implemented"; return nil }

// Copytext is what CMD+C should copy instead of Arg (the default).
func (it *Item) Copytext(s string) *Item { _ = "STUB: not implemented"; return nil }

// Largetype is what is shown in Alfred's Large Text window on CMD+L
// instead of Arg (the default).
func (it *Item) Largetype(s string) *Item { _ = "STUB: not implemented"; return nil }

// Quicklook is a path or URL shown in a macOS Quicklook window on SHIFT
// or CMD+Y.
func (it *Item) Quicklook(s string) *Item { _ = "STUB: not implemented"; return nil }

// Icon sets the icon for the Item.
// Can point to an image file, a filepath of a file whose icon should be used,
// or a UTI.
//
// See the documentation for Icon for more details.
func (it *Item) Icon(icon *Icon) *Item { _ = "STUB: not implemented"; return nil }

// Action sets the value(s) to be passed to Alfred's Universal Actions if
// the user actions this item. Alfred will auto-detect the type of the value(s).
//
// Added in Alfred 4.5.
func (it *Item) Action(value ...string) *Item { _ = "STUB: not implemented"; return nil }

// ActionForType sets the value(s) to be passed to Alfred's Universal Actions if
// the user actions this item. Type may be one of "file", "url" or "text".
//
// Added in Alfred 4.5.
func (it *Item) ActionForType(typ string, value ...string) *Item {
	_ = "STUB: not implemented"
	return nil
}

// Var sets an Alfred variable for subsequent workflow elements.
func (it *Item) Var(k, v string) *Item { _ = "STUB: not implemented"; return nil }

// NewModifier returns an initialised Modifier bound to this Item.
// It also populates the Modifier with any workflow variables set in the Item.
//
// You must specify at least one modifier key. Alfred 3 only supports
// a single modifier, but Alfred 4+ allow them to be arbitrarily combined.
// Any invalid modifier keys are ignored. If you specify an unusable set of
// modifiers (i.e. they evaluate to ""), although a Modifier is returned,
// it is not retained by Item and will not be sent to Alfred. An error message
// is also logged.
func (it *Item) NewModifier(key ...string) *Modifier { _ = "STUB: not implemented"; return nil }

// Add Item variables to Modifier

// SetModifier sets a Modifier for a modifier key.
func (it *Item) SetModifier(m *Modifier) { _ = "STUB: not implemented"; return }

// Cmd returns an initialised Modifier bound to this Item and the CMD (⌘) key.
func (it *Item) Cmd() *Modifier { _ = "STUB: not implemented"; return nil }

// Alt returns an initialised Modifier bound to this Item and the ALT/OPT (⌥) key.
func (it *Item) Alt() *Modifier { _ = "STUB: not implemented"; return nil }

// Opt is a synonym for Alt().
func (it *Item) Opt() *Modifier {
	_ = "STUB: not implemented"

	// Ctrl returns an initialised Modifier bound to this Item and the CTRL (^) key.
	return nil
}

func (it *Item) Ctrl() *Modifier { _ = "STUB: not implemented"; return nil }

// Shift returns an initialised Modifier bound to this Item and the SHIFT (⇧) key.
func (it *Item) Shift() *Modifier { _ = "STUB: not implemented"; return nil }

// Fn returns an initialised Modifier bound to this Item and the fn key.
func (it *Item) Fn() *Modifier { _ = "STUB: not implemented"; return nil }

// Vars returns the Item's workflow variables.
func (it *Item) Vars() map[string]string {
	_ = "STUB: not implemented"

	// MarshalJSON serializes Item to Alfred's JSON format.
	// You shouldn't need to call this directly: use SendFeedback() instead.
	return nil
}

func (it *Item) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Serialise Item

// serialise single arg as string

// itemText encapsulates the copytext and largetext values for a result Item.
type itemText struct {
	// Copied to the clipboard on CMD+C
	Copy *string `json:"copy,omitempty"`
	// Shown in Alfred's Large Type window on CMD+L
	Large *string `json:"largetype,omitempty"`
}

// Modifier encapsulates alterations to Item when a modifier key is held when
// the user actions the item.
//
// Create new Modifiers via Item.NewModifier(). This binds the Modifier to the
// Item, initializes Modifier's map and inherits Item's workflow variables.
// Variables are inherited at creation time, so any Item variables you set
// after creating the Modifier are not inherited.
type Modifier struct {
	// The modifier key, e.g. "cmd", "alt".
	// With Alfred 4+, modifiers can be combined, e.g. "cmd+alt", "ctrl+shift+cmd"
	Key      string
	arg      []string
	subtitle *string
	valid    bool
	icon     *Icon
	vars     map[string]string
}

// newModifier creates a Modifier, validating key.
func newModifier(key ...string) *Modifier { _ = "STUB: not implemented"; return nil }

// Arg sets the arg for the Modifier. Multiple values are allowed in Alfred 4.1 and later.
func (m *Modifier) Arg(s ...string) *Modifier { _ = "STUB: not implemented"; return nil }

// Subtitle sets the subtitle for the Modifier.
func (m *Modifier) Subtitle(s string) *Modifier { _ = "STUB: not implemented"; return nil }

// Valid sets the valid status for the Modifier.
func (m *Modifier) Valid(v bool) *Modifier { _ = "STUB: not implemented"; return nil }

// Icon sets an icon for the Modifier.
func (m *Modifier) Icon(i *Icon) *Modifier { _ = "STUB: not implemented"; return nil }

// Var sets a variable for the Modifier.
func (m *Modifier) Var(k, v string) *Modifier { _ = "STUB: not implemented"; return nil }

// Vars returns all Modifier variables.
func (m *Modifier) Vars() map[string]string {
	_ = "STUB: not implemented"

	// MarshalJSON serializes Item to Alfred 3's JSON format.
	// You shouldn't need to call this directly: use SendFeedback() instead.
	return nil
}

func (m *Modifier) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// serialise single arg as string

// Feedback represents the results for an Alfred Script Filter.
//
// Normally, you won't use this struct directly, but via the Workflow methods
// NewItem(), SendFeedback(), etc. It is important to use the constructor
// functions for Feedback, Item and Modifier structs so they are properly
// initialised and bound to their parent.
type Feedback struct {
	Items  []*Item           // The results to be sent to Alfred.
	NoUIDs bool              // If true, suppress Item UIDs.
	rerun  float64           // Tell Alfred to re-run Script Filter.
	sent   bool              // Set to true when feedback has been sent.
	vars   map[string]string // Top-level feedback variables.
}

// NewFeedback creates a new, initialised Feedback struct.
func NewFeedback() *Feedback { _ = "STUB: not implemented"; return nil }

// Var sets an Alfred variable for subsequent workflow elements.
func (fb *Feedback) Var(k, v string) *Feedback { _ = "STUB: not implemented"; return nil }

// Rerun tells Alfred to re-run the Script Filter after `secs` seconds.
func (fb *Feedback) Rerun(secs float64) *Feedback { _ = "STUB: not implemented"; return nil }

// Vars returns the Feedback's workflow variables.
func (fb *Feedback) Vars() map[string]string {
	_ = "STUB: not implemented"

	// Clear removes any items.
	return nil
}

func (fb *Feedback) Clear() { _ = "STUB: not implemented"; return }

// IsEmpty returns true if Feedback contains no items.
func (fb *Feedback) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// NewItem adds a new Item and returns a pointer to it.
//
// The Item inherits any workflow variables set on the Feedback parent at
// time of creation.
func (fb *Feedback) NewItem(title string) *Item { _ = "STUB: not implemented"; return nil }

// Add top-level variables to Item. The reason for this is that
// (older versions of) Alfred drops all item- and top-level variables
// on the floor if a modifier has any variables set (i.e. only the
// modifier's variables are retained). So, add top-level variables to Item
// (and in turn to any Modifiers) to enforce more sensible behaviour.

// MarshalJSON serializes Feedback to Alfred's JSON format.
// You shouldn't need to call this: use Send() instead.
func (fb *Feedback) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Send generates JSON from this struct and sends it to Alfred
// (by writing the JSON to STDOUT).
//
// You shouldn't need to call this directly: use SendFeedback() instead.
func (fb *Feedback) Send() error { _ = "STUB: not implemented"; return nil }

// Sort sorts Items against query. Uses a fuzzy.Sorter with the specified
// options.
func (fb *Feedback) Sort(query string, opts ...fuzzy.Option) []*fuzzy.Result {
	_ = "STUB: not implemented"
	return nil
}

// Filter fuzzy-sorts Items against query and deletes Items that don't match.
// It returns a slice of Result structs, which contain the results of the
// fuzzy sorting.
func (fb *Feedback) Filter(query string, opts ...fuzzy.Option) []*fuzzy.Result {
	_ = "STUB: not implemented"
	return nil
}

// Keywords implements fuzzy.Sortable.
//
// Returns the match or title field for Item i.
func (fb *Feedback) Keywords(i int) string {
	_ = "STUB: not implemented"

	// Sort on title if match isn't set
	return ""
}

// Len implements sort.Interface.
func (fb *Feedback) Len() int { _ = "STUB: not implemented"; return 0 }

// Less implements sort.Interface.
func (fb *Feedback) Less(i, j int) bool {
	_ = "STUB: not implemented"
	// we only want to sort based on fuzzy match score
	return false
}

// Swap implements sort.Interface.
func (fb *Feedback) Swap(i, j int) { _ = "STUB: not implemented"; return }

// ArgVars lets you set workflow variables from Run Script actions.
// It emits the arg and variables you set in the format required by Alfred.
//
// Use ArgVars.Send() to pass variables to downstream workflow elements.
type ArgVars struct {
	arg  []string
	vars map[string]string
}

// NewArgVars returns an initialised ArgVars object.
func NewArgVars() *ArgVars { _ = "STUB: not implemented"; return nil }

// Arg sets the arg(s)/query to be passed to the next workflow action.
// Multiple values are allowed in Alfred 4.1 and later.
func (a *ArgVars) Arg(s ...string) *ArgVars { _ = "STUB: not implemented"; return nil }

// Vars returns ArgVars' variables.
// NOTE: This function only returns variables you have set with ArgVars.Var()
// for export to Alfred during this run. To read variables from the environment,
// use Workflow.Config.
func (a *ArgVars) Vars() map[string]string {
	_ = "STUB: not implemented"

	// Var sets the value of a workflow variable.
	return nil
}

func (a *ArgVars) Var(k, v string) *ArgVars { _ = "STUB: not implemented"; return nil }

// String returns a string representation.
//
// If any variables are set, JSON is returned. Otherwise, a plain string
// is returned.
func (a *ArgVars) String() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Vars set, so return as JSON

// Send outputs arg and variables to Alfred by printing a response to STDOUT.
func (a *ArgVars) Send() error { _ = "STUB: not implemented"; return nil }

// MarshalJSON serialises ArgVars to JSON.
// You probably don't need to call this: use ArgVars.Send() instead.
func (a *ArgVars) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// Return arg regardless of whether it's empty or not:
	// we have to return *something*
	return nil, nil
}

// Want empty string, i.e. "", not null
