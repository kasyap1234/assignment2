<script>
  import { onMount } from "svelte";
  import { tasks, fetchTasks, deleteTask, editTask } from './TaskStore.js';

  onMount(() => {
    fetchTasks();
  });
</script>

<div class="py-4">
  <h1 class="text-2xl font-semibold mb-4">Task List</h1>
  
  {#each $tasks as task (task.id)}
    <div class="bg-white shadow-md rounded-lg p-4 mb-4">
      <h2 class="text-lg font-semibold">{task.title}</h2>
      <p class="text-gray-600">{task.description}</p>
      <p class="text-sm text-gray-500 mt-1">Due Date: {task.dueDate}</p>
      <div class="mt-2">
        <a href={`/tasks/${task.id}`} class="text-blue-600 hover:underline">View Details</a>
        <span class="mx-2 text-gray-400">|</span>
        <a href={`/tasks/edit/${task.id}`} class="text-green-600 hover:underline">Edit</a>
        <span class="mx-2 text-gray-400">|</span>
        <button class="text-red-600 hover:underline" on:click={() => deleteTask(task.id)}>Delete</button>
      </div>
    </div>
  {:else}
    <p>No tasks found.</p>
  {/each}
</div>

<style>
  /* Tailwind CSS or custom styles */
</style>
