import adapter from '@sveltejs/adapter-auto';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';
/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter()
	},
	hydrate: true,
	// enable SSR and statically generate pages by default, see https://kit.svelte.dev/docs#routing-pages
	ssr: true,
	// enable client-side navigation
	router: true,
	// enable end-to-end caching via HTTP headers
	preprocess: vitePreprocess()
};
export default config;
