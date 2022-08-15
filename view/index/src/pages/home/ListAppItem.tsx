import React, { useState, useEffect, useMemo, useRef, useCallback } from 'react';
import { List, Image, Col, Collapse, BackTop, Badge, Divider, Row, Typography, message } from 'antd';
import { DownOutlined, UpOutlined, AppleFilled, AndroidFilled } from '@ant-design/icons'
import styles from './index.module.less'
import service from '@/service';
import dayjs from 'dayjs';
import Configs from '@/Configs';
import { observer } from 'mobx-react'

const ListAppItem = observer((props: {
  item: any;
  index: any;
}) => {
  const { item, index } = props
  const [expanded, setExpanded] = useState(false)
  const collapseActiveKey = useMemo(() => expanded ? index : undefined, [expanded])
  const [versions, setVersions] = useState<any[]>([])
  const toggleExpand = useCallback(async () => {
    let t = !expanded
    setExpanded(t)
    if (t) {
      try {
        let rs: any = await service.appVersions(item.id)
        if (rs.message) {
          message.error(rs.message)
        } else {
          setVersions(rs.data)
        }
      } catch (error: any) {
        message.error(error.message)
      }
    }
  }, [expanded])
  return (
    <List.Item onClick={toggleExpand} className={styles.item}>
      <Row style={{ width: '100%' }}>
        <Col style={{ width: '100%' }}>
          <Row justify='space-between' align='middle' style={{ width: '100%', padding: '0px 15px' }}>
            <Col>
              <Row gutter={12}>
                <Col>
                  <Image src={item.logo} preview={false} className={styles.icon} />
                </Col>
                <Col>
                  <Typography.Title level={5}>{item.name}</Typography.Title>
                  <Typography.Text>{item.bundle_id}</Typography.Text>
                </Col>
              </Row>
            </Col>
            <Col>
              {expanded ? <UpOutlined className={styles.c_icon} /> : <DownOutlined className={styles.c_icon} />}
            </Col>
          </Row>
          <Row style={{ width: '100%' }}>
            <Collapse
              ghost
              activeKey={collapseActiveKey}
              style={{ width: '100%' }}>
              <Collapse.Panel key={index} header={""} showArrow={false} style={{ width: '100%' }}>
                {versions.map((ver, ind) => {
                  let hasIPA = ver.ipa_url ?? '' != ''
                  let hasAPK = ver.apk_url ?? '' != ''
                  let upDate = dayjs(ver.updatedAt).format('YYYY/MM/DD HH:mm:ss')
                  let version = ver.version ?? '未知版本'
                  return (
                    <Row
                      key={ver.id}
                      justify='space-between'
                      align='middle'
                      onClick={() => {
                        let u = `${Configs.SERVER_ROOT}/app/install/${item.bundle_id}/${version}`
                        window.open(u,'_blank')
                      }}
                      style={{ backgroundColor: '#fafafc', padding: '6px 5px', marginBottom: 5 }}>
                      <Col>
                        <Row>
                          <Col>
                            <Badge status='success' />
                          </Col>
                          <Col>
                            <Typography.Text>
                              版本号
                              <Typography.Text type='secondary' style={{ paddingLeft: 5 }}>{version}</Typography.Text>
                            </Typography.Text>
                          </Col>
                        </Row>
                        <Row>
                          <Col>
                            <Badge status='success' style={{ opacity: 0 }} />
                          </Col>
                          <Col>
                            <Typography.Text type='secondary' style={{ fontSize: 11 }}>
                              上次更新 {upDate}
                            </Typography.Text>
                          </Col>
                        </Row>
                      </Col>
                      <Col>
                        <Row gutter={16}>
                          <Col>
                            <div className={styles.p_btn} style={{ opacity: hasAPK ? 1.0 : 0.38 }}>
                              <AndroidFilled style={{ color: '#565657' }} />
                            </div>
                          </Col>
                          <Col>
                            <div className={styles.p_btn} style={{ opacity: hasIPA ? 1.0 : 0.38 }}>
                              <AppleFilled style={{ color: '#565657' }} />
                            </div>
                          </Col>
                        </Row>
                      </Col>
                    </Row>
                  )
                })
                }
              </Collapse.Panel>
            </Collapse>
          </Row>
        </Col>
      </Row>
    </List.Item>
  )
})
export default ListAppItem;