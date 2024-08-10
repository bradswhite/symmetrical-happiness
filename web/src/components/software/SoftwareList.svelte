<script lang='ts'>
  import SoftwareCard from './SoftwareCard.svelte';

  interface SoftwareType {
    id: string;
    name: string;
    title: string;
    description: string;
    image: string;
    url: string;
    username: string;
    createdAt: string;
  }
  
  const getSoftwareList = async (): SoftwareType => {
    const url = import.meta.env.PUBLIC_API_URL;
    const res = await fetch(`${url}/software`, {
      headers: {'Access-Control-Allow-Origin': '*'}
    });
    return await res.json();
  };

</script>

<section class='w-full justify-center py-4'>
  <a href='/#software' class='flex inline-flex items-center px-8 py-4'>
    <svg class='stroke-gray-800 dark:stroke-gray-200 mr-2' width="2em" height="2em" viewBox="0 0 14 14" >
      <g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3 5V1.5a1 1 0 0 1 1-1h8.5a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H11m-8-7h10.5"/><rect width="8" height="6" x=".5" y="7.5" rx="1"/><path d="M.5 10h8"/></g>
    </svg>
    <h3 id='software' class='text-xl font-crenzo text-gray-800 dark:text-gray-200 py-4'>Software List</h3>
  </a>

  <div class='grid gap-4'>
    {#await getSoftwareList()}
      <p>Loading software data...</p>
    {:then softwareList}
      {#each softwareList as software}
        <SoftwareCard {software} />
      {/each}
    {:catch error}
      <p>Cannot load software data</p>
      <p>{error.message}</p>
    {/await}
  </div>
</section>
