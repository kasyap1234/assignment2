<script>
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";
  import { writable, get } from "svelte/store";
  import { fetchTasks } from './TaskStore.js';

  export let taskId = null;

  const task = writable({
    title: "",
    description: "",
    dueDate: "",
  });
 let formMode = taskId ? "Edit Task" : "Add New Task";
  async function fetchTaskDetails() {
    if (taskId) {
      const res = await fetch(`http://localhost:8000/api/tasks/${taskId}`);
      if (res.ok) {
        const data = await res.json();
        task.set(data);
      } else {
        console.error('Failed to fetch task details', res.statusText);
      }
    }
  }

  async function handleSubmit() {
    const currentTask = get(task);
    let res;
    if (taskId) {
      res = await fetch(`http://localhost:8000/api/tasks/${taskId}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(currentTask),
      });
    } else {
      res = await fetch("http://localhost:8000/api/tasks", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(currentTask),
      });
    }

    if (res.ok) {
      fetchTasks(); // Refresh task list after update or creation
      goto("/tasks"); // Redirect to task list after successful update or creation
    } else {
      console.error('Failed to submit task', res.statusText);
    }
  }

  onMount(fetchTaskDetails);
</script>

<div class="py-4">
  <h1 class="text-2xl font-semibold mb-4">{formMode}</h1>
  
  <form on:submit|preventDefault={handleSubmit}>
    <div class="mb-4">
      <label for="title" class="block text-sm font-medium text-gray-700">Title</label>
      <input type="text" id="title" bind:value={$task.title} class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
    </div>

    <div class="mb-4">
      <label for="description" class="block text-sm font-medium text-gray-700">Description</label>
      <textarea id="description" bind:value={$task.description} rows="3" class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200 focus:ring-opacity-50"></textarea>
    </div>

    <div class="mb-4">
      <label for="dueDate" class="block text-sm font-medium text-gray-700">Due Date</label>
      <input type="date" id="dueDate" bind:value={$task.dueDate} class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
    </div>

    <div class="mt-6">
      <button type="submit" class="inline-block bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-600">{taskId ? "Update Task" : "Add Task"}</button>
    </div>
  </form>
</div>

<style>
  /* Tailwind CSS or custom styles */
</style>
