// Package tasklist provides methods to manage a todolist.
package tasklist

// TaskList listes tasks.
type TaskList struct {
	list []Task
}

// Add method adds a item to the list.
func (l *TaskList) Add(t Task) {
	l.list = append(l.list, t)
}

// Remove method removes a item in the specified index.
// This index starts 0.
func (l *TaskList) Remove(idx int) (ok bool) {
	if 0 <= idx && idx < len(l.list) {
		l.list = append(l.list[:idx], l.list[idx+1:]...)
		return true
	}
	return false
}

// Toggle method toggles whether the task is finished or not.
// This index starts 0.
func (l *TaskList) Toggle(idx int) (ok bool) {
	if 0 <= idx && idx < len(l.list) {
		l.list[idx].Toggle()
		return true
	}
	return false
}
