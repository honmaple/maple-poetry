import { getCurrentInstance } from 'vue'
import useIcon from './icon'

export default function useApp() {
	const { proxy } = getCurrentInstance()

	return {
		app: proxy,
		icon: useIcon(),
	}
}