import COS from 'cos-js-sdk-v5'
import { useState, memo, useRef, useMemo, useEffect } from 'react'
import { Upload, message } from 'antd'
import { LoadingOutlined, UploadOutlined } from '@ant-design/icons'

import styles from './style.module.less'
import _ from 'lodash'
import service from '@/service'
import { observer } from 'mobx-react'

function getFileExtension(filename: string): string {
  return filename.slice(((filename.lastIndexOf('.') - 1) >>> 0) + 2)
}

function beforeUpload(file: File) {
  const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png'
  if (!isJpgOrPng) {
    message.error('只允许上传 JPG/PNG 文件!')
  }
  const isLt2M = file.size / 1024 / 1024 < 2
  if (!isLt2M) {
    message.error('图片必须小于 2MB!')
  }
  return isJpgOrPng && isLt2M
}
type FileUploaderValue = {
  location: string;
  objectPath: string;
}
interface FCProps {
  path?: string
  value?: FileUploaderValue
  onChange?: (value: FileUploaderValue) => void
}

function CoverUpload(props: FCProps) {
  const { value } = props
  const [loading, setLoading] = useState(false)
  const [v, setV] = useState(props.value)
  const img = useMemo(() => {
    if (v?.location) {
      if (v?.location?.startsWith('http')) {
        return v.location as string
      }
      return 'https://' + v.location
    }
    return ''
  }, [v])

  useEffect(() => {
    if(value?.location) {
      setV(value)
    }
  },[value])

  const uploadButton = (
    <div className={styles.button}>
      {loading ? <LoadingOutlined /> : <UploadOutlined />}
      <div className={styles.text}>点击上传应用图标</div>
    </div>
  )
  return (
    <Upload
      name="logo"
      className={styles.upload}
      listType="picture-card"
      accept={'.jpg,.png,.jpeg'}
      showUploadList={false}
      maxCount={1}
      beforeUpload={beforeUpload}
      customRequest={(options) => {
        setLoading(true)
        // console.log('上传的文件: ', options);
        const file: any = options.file
        service.uploadFile(
          file,
          `${service.cosConfigs.COS_RES_DIR}/tmp/${file.name}`, //service.cosConfigs.COS_RES_DIR => /distribute/app
          (progress) => {
            // console.log('上传进度: ', progress);
          },
          (err, data, opts) => {
            // console.log('上传结束: ', err, data, opts);
            setLoading(false)
            if (err == null) {
              //上传文件完成、拿到 Location 和 key(objectPath)
              let result = {
                location: data.Location,
                objectPath: opts.Key
              }
              setV(result)
              if (props.onChange) props.onChange(result)
            } else {
              message.error(err.message)
            }
          }
        )
      }}
    >
      {v?.location ? <img className={styles.cover} src={img} alt="cover" /> : uploadButton}
    </Upload>
  )
}

export default observer(CoverUpload)
