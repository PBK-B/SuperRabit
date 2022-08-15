import React, { useEffect, useMemo, useRef, useState } from "react";
import { Modal, Upload, message, AutoComplete, Select, Form, Input } from "antd";
import { InboxOutlined } from '@ant-design/icons';
import CoverUpload from "@/components/CoverUpload";
import service from "@/service";
import { observer } from 'mobx-react'
import _ from 'lodash'
const { Dragger } = Upload;

type FileUploaderValue = {
    location: string;
    objectPath: string;
    size: number;
}
const FileUploader = observer((props: {
    value?: FileUploaderValue;
    onChange?: (value: FileUploaderValue) => void;
}) => (

    <Dragger
        name='file'
        multiple={false}
        accept={".apk,.ipa"}
        customRequest={async (options) => {
            // console.log('上传的文件: ', options);
            const { onSuccess, onProgress } = options
            const file: any = options.file
            service.uploadFile(
                file,
                `${service.cosConfigs.COS_RES_DIR}/tmp/${file.name}`, //service.cosConfigs.COS_RES_DIR => /distribute/app
                (progress) => {
                    // console.log('上传进度: ', progress);
                    if (onProgress) {
                        onProgress({ percent: progress.percent * 100 })
                    }
                },
                (err, data, opts) => {
                    // console.log('上传结束: ', err, data, opts);
                    if (err == null) {
                        //上传文件完成、拿到 Location 和 key(objectPath)
                        let result = {
                            location: data.Location,
                            objectPath: opts.Key,
                            size: file.size
                        }
                        if (onSuccess) onSuccess(result)
                        if (props.onChange) props.onChange(result)
                    } else {
                        message.error(err.message)
                    }
                }
            )
        }}
    >
        <p className="ant-upload-drag-icon">
            <InboxOutlined />
        </p>
        <p className="ant-upload-text">点击或拖动文件到这里来上传应用文件</p>
        <p className="ant-upload-hint">
            仅支持上传 .apk 和 .ipa 格式的应用文件；请勿上传其它非相关文件；
        </p>
    </Dragger >
));

function amendHttpPrefix(url: string) {
    if (!url.startsWith('http')) {
        return 'https://' + url
    }
    return url
}

const Publish = observer((props: {
    hide: () => void;
    visible: boolean;
    onPublishSuccess?: () => void;
}) => {
    const { hide, visible } = props
    const [form] = Form.useForm()
    const [apps, setApps] = useState<any[]>([])
    const bundleids = useMemo(() => apps.map((v, i) => ({ label: v.bundle_id, value: v.bundle_id })), [apps])
    const [selectedApp, setSelectedApp] = useState<any>()
    const bundleid = useRef('')
    const tmpFiles = useRef<any>({}).current
    const onFinish = (values: any) => {
        let ac = ((values.access_code ?? '') as string).trim()
        let formValue: any = {
            logo: amendHttpPrefix(values.logo.location),
            bundle_id: bundleid.current,
            ...(_.pick(values, ['name', 'version', 'build'])),
            size: values.app.size,
            access: ac == '' ? 'public' : 'private',
            access_code: ac
        }
        if (values.app.location.indexOf('.apk') != -1) {
            formValue.apk_url = amendHttpPrefix(values.app.location)
        } else {
            formValue.ipa_url = amendHttpPrefix(values.app.location)
        }
        if (selectedApp) {
            formValue.id = selectedApp.id
        }
        // console.log(formValue);
        let remove = message.loading("正在发布更新中、请勿操作...")
        service.publishApp(formValue)
            .then(rs => {
                remove()
                hide()
                if (rs.message) {
                    message.error(rs.message)
                } else {
                    message.success("发布成功")
                    if (props.onPublishSuccess) props.onPublishSuccess()
                }
            })
            .catch(e => {
                hide()
                remove()
                message.error(e.message)
            })
    }
    const onBundleCompleteChange = (value: string, option: {
        label: string;
        value: string;
    } | {
        label: string;
        value: string;
    }[]) => {
        let foundApp;
        for (let a of apps) {
            if (a.bundle_id == value) {
                foundApp = a;
                break;
            }
        }
        bundleid.current = value
        if (foundApp) {
            form.setFieldsValue({
                name: foundApp.name,
                logo: { location: foundApp.logo, objectPath: '' },
            })
        }
    }
    useEffect(() => {
        service.appList().then(rs => {
            if (rs.message) {
                message.error(rs.message)
            } else {
                setApps(rs.data)
            }
        }).catch(e => {
            message.error(e.message)
        })
    }, [])
    return (
        <Modal
            visible={visible}
            width={500}
            centered
            okText="发布"
            cancelText="取消"
            bodyStyle={{
                padding: 0,
                margin: 0,
                height: '80vh',
            }}
            onOk={() => {
                form.submit()
            }}
            onCancel={() => {
                for (let k of Object.keys(tmpFiles)) {
                    let obj = tmpFiles[k]
                    if (obj) {
                        service.deleteFile(obj)
                    }
                }
                form.resetFields()
                hide()
            }}
        >
            <div style={{ height: '100%', width: '100%', overflow: 'scroll' }}>
                <Form
                    form={form}
                    name={'publish-app'}
                    onFinish={onFinish}
                    initialValues={{
                        build: 1
                    }}
                    onFieldsChange={(changed, all) => {
                        changed.map((v: any, i) => {
                            let n = v.name.toString()
                            if (n.indexOf('app') != -1) {
                                if ((v.value?.objectPath ?? '').length > 0) {
                                    tmpFiles[n] = v.value.objectPath
                                }
                            } else if (n.indexOf('logo') != -1) {
                                if ((v.value?.objectPath ?? '').length > 0) {
                                    tmpFiles[n] = v.value.objectPath
                                }
                            }
                        })
                    }}
                    labelCol={{ span: 4 }}
                    style={{ padding: '0 15px', paddingTop: 50 }}>
                    <Form.Item name={'app'} required>
                        <FileUploader />
                    </Form.Item>
                    <Form.Item name={'bundle_id'} label={'应用包名'} required>
                        <Input.Group compact>
                            <AutoComplete
                                placeholder="应用的唯一标识、例如: com.xxx"
                                options={bundleids}
                                onChange={onBundleCompleteChange}
                                style={{ width: '100%' }}
                            />
                        </Input.Group>
                    </Form.Item>
                    <Form.Item name={'name'} label={'应用名称'} required>
                        <Input disabled={selectedApp?.name != undefined} />
                    </Form.Item>
                    <Form.Item name={'access_code'} label={'访问密码'} tooltip="不设置则默认所有人访问下载链接均可直接下载">
                        <Input placeholder="设置App下载密码" />
                    </Form.Item>
                    <Form.Item name={'logo'} label={'应用图标'} required>
                        <CoverUpload />
                    </Form.Item>
                    <Form.Item name={'version'} label={'版本号'} required>
                        <Input placeholder="版本号、例如1.0.0" style={{ width: '50%' }} />
                    </Form.Item>
                    <Form.Item name={'build'} label={'编译号'} required>
                        <Input type={'number'} placeholder="当前版本第几次编译" style={{ width: '30%' }} />
                    </Form.Item>
                </Form>
            </div>
        </Modal>
    )
})
export default Publish