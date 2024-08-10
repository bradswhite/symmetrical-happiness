<script lang='ts'>
  import HeaderLink from './HeaderLink.astro';

  let href: string;
  let text: string;
</script>

<HeaderLink {href} class='relative inline-block w-full h-full px-4 py-5 mx-2 font-medium leading-tight text-center md:py-2 group md:w-auto md:px-2 lg:mx-3 md:text-center duration-300 ease-out text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100'>
  <div class='flex items-center space-x-2'>
    <slot />
    <span>{text}</span>
  </div>
</HeaderLink>

<style>
.active > div{
  @apply text-gray-900 dark:text-gray-100;
}
</style>
