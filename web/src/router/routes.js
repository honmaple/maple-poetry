const routes = [
	{
		path: '/',
		component: () => import('layouts/Layout.vue'),
		children: [
			{
				path: '',
				component: () => import('pages/Index.vue'),
			},
			{
				path: '/authors',
				component: () => import('pages/Authors.vue'),
				meta: {
					title: "诗人",
				}
			},
			{
				path: '/authors/:id',
				component: () => import('pages/Author.vue'),
				meta: {
					title: "诗人详情",
				}
			},
			{
				path: '/poems',
				component: () => import('pages/Poems.vue'),
				meta: {
					title: "诗词",
				}
			},
			{
				path: '/poems/:id',
				component: () => import('pages/Poem.vue'),
				meta: {
					title: "诗词详情",
				}
			},
			{
				path: '/collections',
				component: () => import('pages/Collections.vue'),
				meta: {
					title: "诗集",
				}
			},
			{
				path: '/about',
				component: () => import('pages/About.vue'),
				meta: {
					title: "关于",
				}
			}
		]
	},

	// Always leave this as last one,
	// but you can also remove it
	{
		path: '/:catchAll(.*)*',
		component: () => import('pages/404.vue')
	}
]

export default routes
