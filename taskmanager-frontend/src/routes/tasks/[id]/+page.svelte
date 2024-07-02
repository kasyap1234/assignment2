<script>
  import { page } from '$app/stores';
  import { onMount } from 'svelte';
  import TaskDetails from '$lib/components/TaskDetails.svelte';

  let taskId;

  onMount(() => {
    taskId = $page.params.id;
  });
</script>

<TaskDetails {taskId} />

<style>
  /* Tailwind CSS or custom styles */
</style>

