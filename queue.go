package main

type TaskQueue chan TaskQueue

func enqueTask(queue TaskQueue, task Task) {
	queue <- task
}

