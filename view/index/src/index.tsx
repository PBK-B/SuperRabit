import './global.less'
import { createRoot } from 'react-dom/client'
import { RecoilRoot } from 'recoil'
import { ConfigProvider } from 'antd'
import Router from './router'


const root = createRoot(document.getElementById('root')!)

root.render(
  <RecoilRoot>
    <ConfigProvider>
      <Router />
    </ConfigProvider>
  </RecoilRoot>
)
