package tasklist

import "time"

// Task is an struct for the todo information.
// content is a content of a task.
// date is a date of a task.
// done is a boolean whether task is finished of not.
type Task struct {
	content string
	date    time.Time
	done    bool
}

// If the date is before today, return date for next year.
// Otherwise, return date for this year.
func calcDate(month, day int) time.Time {
	today := time.Now().Local()
	date := time.Date(today.Year(), time.Month(month), day, 23, 59, 59, 0, time.Local)
	if date.Before(today) {
		date = date.AddDate(1, 0, 0)
	}

	return date
}

// NewTodo creates a item by the information of the task.
func NewTodo(content string, month, day int) *Task {
	t := Task{
		content: content,
		date:    calcDate(month, day),
		done:    false,
	}
	return &t
}

// Content returns the content of the item.
func (t Task) Content() string {
	return t.content
}

// Date returns the date of the item.
func (t Task) Date() time.Time {
	return t.date
}

// Done reports whether the item is finished or not.
func (t Task) Done() bool {
	return t.done
}

// Toggle changes the "done" state of the todo.
// If the "done" value is true, Toggle changes the "done" value to false.
func (t *Task) Toggle() {
	t.done = !t.done
}

// Days returns the number of days from today to the task date.
func (t *Task) Days() int {
	today := time.Now().Local()
	return int(time.Duration.Hours(t.date.Sub(today)) / 24)
}
