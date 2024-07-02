import { writable } from 'svelte/store';

export const tasks = writable([]);

export async function fetchTasks() {
    const res = await fetch("http://localhost:8000/api/tasks");
    if (res.ok) {
        const data = await res.json();
        tasks.set(data);
    }
}

export async function deleteTask(taskId) {
    const res = await fetch(`http://localhost:8000/api/tasks/${taskId}`, {
        method: "DELETE",
    });
    if (res.ok) {
        fetchTasks(); // Refresh task list after deletion
    }
}
export async function editTask(taskId, newTask) {
    const res = await fetch(`http://localhost:8000/api/tasks/${taskId}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(newTask),
    });
    if (res.ok) {
        fetchTasks(); // Refresh task list after deletion
    }
}
