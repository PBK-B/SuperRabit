import { useState, useEffect, useMemo, useRef } from 'react';
import { List, Image, Col, Spin, Skeleton, Collapse, BackTop, Badge, Divider, Row, Typography, Grid, Form, Input, message } from 'antd';
import InfiniteScroll from 'react-infinite-scroll-component';
import { observer } from 'mobx-react'
import ListAppItem from './ListAppItem';
import styles from './index.module.less'
import SearchBar from './SearchBar'
import Publish from './Publish'
import service from '@/service';

const Loading = () => {
  return (
    <Row gutter={12} align='middle' justify='center' className={styles.loading_box}>
      <Col>
        <Spin size="small" style={{ transform: 'translateY(1px)' }} />
      </Col>
      <Col>
        <Typography.Text style={{ color: '#78787a', fontSize: 16 }}>奋力加载中...</Typography.Text>
      </Col>
    </Row>
  )
}
const Home = observer((props:any) => {
  const [hasMore, setHasMore] = useState(false);
  const [loading, setLoading] = useState(false);
  const [data, setData] = useState<any[]>([]);
  const [publishModalVisible, setPublishModalVisible] = useState(false)
  const loadMoreData = async () => {
    if (loading) {
      return;
    }
    setLoading(true);
    try {
      let rs = await service.appList()
      // console.log(rs);
      if (rs.message) {
        message.error(rs.message)
      } else {
        setData(rs.data)
      }
    } catch (error: any) {
      message.error(error.message)
    }
    setLoading(false)
  };
  useEffect(() => {
    loadMoreData();
  }, []);
  useEffect(() => {
    var count = 0
    window.onkeydown = (ev) => {
      if (publishModalVisible) return
      if (ev.code == "KeyN") {
        count += 1
        if (count == 3) {
          count = 0
          setPublishModalVisible(true)
        }
      } else {
        if (count > 0) count = 0
      }
    }
  }, [publishModalVisible])
  const onPublishSuccess = () => {
    loadMoreData()
  }

  if (loading && data.length == 0) return <Loading />
  return (
    <>
      <Publish
        visible={publishModalVisible}
        hide={() => setPublishModalVisible(false)}
        onPublishSuccess={onPublishSuccess}
      />
      <div id={'scrollable'} className={styles.list}>
        <SearchBar />
        <InfiniteScroll
          dataLength={data.length}
          next={loadMoreData}
          hasMore={hasMore}
          loader={loading && <Skeleton paragraph={{ rows: 1 }} active />}
          endMessage={
            <Divider plain>没有更多了🤐</Divider>
          }
          scrollableTarget="scrollable"
          style={{ marginTop: 56 }}
        >
          <List
            dataSource={data}
            renderItem={(item, index) => <ListAppItem item={item} index={index} />}
          />
        </InfiniteScroll>
      </div>
    </>
  );
});

export default Home;

// grid={{
//     xs: 1,
//     sm: 2,
//     md: 3,
//     lg: 4,
//     xl: 4,
//     xxl: 5  
// }}