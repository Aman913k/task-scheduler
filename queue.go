package main

type TaskQueue chan Task

func enqueTask(queue TaskQueue, task Task) {
	queue <- task
}
