<script>
  /*import DarkModeLink from '@components/DarkModeLink.astro';
  import LogoLink from '@components/LogoLink.astro';
  import NavLinkBase from './NavLinkBase.svelte';
  import NavLink from './NavLink.astro';
  import ContactLink from '@components/contact/ContactLink.astro';
  import ProductsLink from '@components/products/ProductsLink.astro';
  import AddProductLink from '@components/add-product/AddProductLink.astro';
  import MoreLink from '@components/MoreLink.astro';

  import SignInBtn from '@components/auth/SignInBtn.astro';
  import SignUpBtn from '@components/auth/SignUpBtn.astro';*/
</script>

<div class="flex items-center justify-start w-1/4 h-full pr-4">
  <LogoLink />
</div>

<NavLinkBase>
  <div class="flex flex-col items-start justify-center w-full space-x-4 text-center md:w-2/3 md:mt-0 md:flex-row md:items-center">
      <NavLink url='/' text='Home'>
        <svg width="1em" height="1em" viewBox="0 0 14 14">
          <title>Home</title>
          <path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 6.94a.999.999 0 0 0-.32-.74L7 .5L.82 6.2a1 1 0 0 0-.32.74v5.56a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1zM7 13.5v-4"/>
        </svg>
      </NavLink>

      <ProductsLink />

      <ContactLink />

      <AddProductLink />

      <MoreLink />

      <a href="#_" class="absolute top-0 left-0 hidden py-2 mt-6 ml-10 mr-2 text-gray-600 lg:inline-block md:mt-0 md:ml-2 lg:mx-3 md:relative">
          <svg class="inline w-5 h-5" fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" stroke="currentColor"><path d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
      </a>
  </div>
  <div class="flex flex-col items-start justify-end w-full pt-4 md:items-center md:w-1/3 md:flex-row md:py-0">
    <DarkModeLink />
    <SignInBtn />
    <SignUpBtn />
  </div>
</NavLinkBase>
