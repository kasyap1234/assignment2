import adapter from '@sveltejs/adapter-node';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

const config = {
	kit: {
		adapter: adapter(),
		// Removed the preprocess option from here
		
			// other Vite configurations...
		},
		// other configurations...
	
};

export default config;