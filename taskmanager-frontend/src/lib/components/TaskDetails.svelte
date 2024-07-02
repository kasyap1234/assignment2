<script>
  import { onMount } from "svelte";
  import { goto } from '$app/navigation';
  

  export let taskId;
  let task = {};

  onMount(async () => {
    const res = await fetch(`http://localhost:8000/api/tasks/${taskId}`);
    if (res.ok) {
      task = await res.json();
    }
  });

  async function deleteTask() {
    const res = await fetch(`http://localhost:8000/api/tasks/${taskId}`, {
      method: "DELETE",
    });
    if (res.ok) {
      goto('/tasks'); // Redirect to task list after successful deletion
    }
  }
</script>

<div class="py-4">
  <h1 class="text-2xl font-semibold mb-4">{task.title}</h1>
  
  <div class="bg-white shadow-md rounded-lg p-4 mb-4">
    <p class="text-gray-600">{task.description}</p>
    <p class="text-sm text-gray-500 mt-1">Due Date: {task.dueDate}</p>
    <div class="mt-2">
      <a href={`/tasks/edit/${task.id}`} class="text-green-600 hover:underline">Edit</a>
      <span class="mx-2 text-gray-400">|</span>
      <button class="text-red-600 hover:underline" on:click={deleteTask}>Delete</button>
    </div>
  </div>
</div>

<style>
  /* Tailwind CSS or custom styles */
</style>

