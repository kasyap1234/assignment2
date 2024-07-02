<script>
  import { onMount } from 'svelte';
  import "../app.css"
</script>

<nav class="bg-gray-800 p-4 text-white">
  <a href="/" class="mr-4">Home</a>
  <a href="/tasks" class="mr-4">Tasks</a>
  <a href="/tasks/new" class="mr-4">Add Task</a>
</nav>

<main class="container mx-auto p-4">
  <slot></slot>
</main>

<style>
  /* Tailwind CSS or custom styles */
</style>
