# Todo API Acceptance Tests

This is an executable specification file. This file follows markdown syntax.
Every heading in this file denotes a scenario. Every bulleted point denotes a step.

## Add first todo

* Clear all todos
* Add todo "buy some milk"
* See "buy some milk" in list

## Add second todo

* Todo list with "buy some milk"
* Add todo "enjoy the assignment"
* See "buy some milk" in list
* See "enjoy the assignment" in list

## Mark as completed

* Todo list with "buy some milk"
* Click "buy some milk"
* Check "buy some milk" is completed

## Mark as uncompleted

* Todo list with completed "buy some milk"
* Click "buy some milk"
* Check "buy some milk" is not completed

## Delete todo

* Todo list with "rest for a while"
* Remove todo "rest for a while"
* Check todo list is empty

## List with multiple todo
* Todo list with "rest for a while" and "drink water"
* Remove todo "rest for a while"
* See "drink water" in list