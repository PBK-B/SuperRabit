import { Suspense, useState, memo } from 'react'
import { Outlet } from 'react-router-dom'
import { Layout } from 'antd'

const { Header, Sider } = Layout

function AppLayout() {

  return (
    <Layout>
        <Outlet />
    </Layout>
  )
}

export default memo(AppLayout)
