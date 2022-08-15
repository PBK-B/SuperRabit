import { Suspense, lazy } from 'react'
import { BrowserRouter, Routes, Route, Navigate, useLocation } from 'react-router-dom'

const NotFound = lazy(() => import(/* webpackChunkName: "errors.404" */ './pages/errors/404'))
const Home = lazy(() => import(/* webpackChunkName: "home" */ './pages/home/Home'))
const AppLayout = lazy(() => import(/* webpackChunkName: "home" */ './layout'))

function Router() {
	return (
		<BrowserRouter>
			<Routes>
				<Route path="/" element={<AppLayout />}>
					<Route index element={<Home />} />
				</Route>
				<Route path="*" element={<NotFound />} />
			</Routes>
		</BrowserRouter>
	)
}

export default Router
