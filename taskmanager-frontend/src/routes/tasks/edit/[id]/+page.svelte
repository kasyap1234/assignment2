<script context="module">
  export async function load({ params }) {
    return { props: { taskId: params.id } };
  }
</script>

<script>
  import EditTask from '$lib/components/EditTask.svelte';
  export let taskId;
</script>

<EditTask {taskId} />
