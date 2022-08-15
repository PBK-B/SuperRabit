import Configs from '@/Configs';
import COS from 'cos-js-sdk-v5';
import { makeAutoObservable, runInAction } from 'mobx'

class Service {

    cosConfigs = {
        COS_BUCKET_NAME: '',
        COS_APP_ID: '',
        COS_REGION: '',
        COS_CDN_URL: '',
        COS_RES_DIR: '',
    }

    constructor() {
        makeAutoObservable(this)
        this.setCosConfig()
    }
    setCosConfig(onEnd?: (cf: any) => void) {
        let url = `${Configs.SERVER_ROOT}/user/cos`
        fetch(url, {
            method: 'GET'
        }).then(rs => rs.json())
            .then(rs => {
                let data = rs.data ?? {}
                if (rs.data) {
                    runInAction(() => {
                        this.cosConfigs = {
                            COS_BUCKET_NAME: data.COS_BUCKET_NAME,
                            COS_APP_ID: data.COS_APP_ID,
                            COS_REGION: data.COS_REGION,
                            COS_CDN_URL: data.COS_CDN_URL,
                            COS_RES_DIR: data.COS_RES_DIR
                        }
                    })
                    if(onEnd) onEnd(data)
                }
            })
    }

    getCOSClient() {
        return new COS({
            getAuthorization: function (options: any, callback: any) {
                let url = `${Configs.SERVER_ROOT}/user/cos`
                fetch(url, {
                    method: 'GET'
                })
                    .then(rs => rs.json())
                    .then(rs => {
                        let data = rs.data ?? {}
                        if (rs.data) {
                            callback({
                                TmpSecretId: data.Data.Credentials?.TmpSecretId,
                                TmpSecretKey: data.Data.Credentials?.TmpSecretKey,
                                SecurityToken: data.Data.Credentials?.Token,
                                StartTime: data.Data.StartTime,
                                ExpiredTime: data.Data.ExpiredTime,
                            })
                        }
                    })
            }
        })
    }

    _uploadFile(cf: any, key: string, file: any, onProgress: any, onFileFinish: any) {
        const cos = this.getCOSClient()
        cos.uploadFile({
            Bucket: `${cf.COS_BUCKET_NAME}-${cf.COS_APP_ID}`,
            Region: cf.COS_REGION,
            Key: key,
            Body: file,
            onProgress: function (progress: any) {
                if (onProgress) onProgress(progress)
            },
            onFileFinish: function (err: any, data: any, options: any) {
                if (onFileFinish) onFileFinish(err, data, options)
            }
        })
    }
    _deleteFile(cf:any,key:string) {
        const cos = this.getCOSClient()
        cos.deleteObject({
            Bucket: `${cf.COS_BUCKET_NAME}-${cf.COS_APP_ID}`,
            Region: cf.COS_REGION,
            Key: key,
        })
    }

    async uploadFile(
        file: any,
        objectPath: string,
        onProgress?: (progress: COS.ProgressInfo) => void,
        onFileFinish?: (err: Error, data: COS.UploadFileItemResult, options: COS.UploadFileItemParams) => void) {
        if (this.cosConfigs.COS_BUCKET_NAME == '') {
            this.setCosConfig((cf) => {
                this._uploadFile(cf, objectPath, file, onProgress, onFileFinish)
            })
        } else {
            this._uploadFile(this.cosConfigs, objectPath, file, onProgress, onFileFinish)
        }
    }
    async deleteFile(objectPath: string) {
        if (this.cosConfigs.COS_BUCKET_NAME == '') {
            this.setCosConfig((cf) => {
                this._deleteFile(cf,objectPath)
            })
        } else {
            this._deleteFile(this.cosConfigs,objectPath)
        }
    }

    async appList() {
        return fetch(`${Configs.SERVER_ROOT}/app/list`).then(rs => rs.json())
    }
    async appVersions(id: string) {
        return fetch(`${Configs.SERVER_ROOT}/app/versions?id=${id}`).then(rs => rs.json())
    }
    async publishApp(params: any) {
        let form = new FormData();
        Object.keys(params).map(key => {
            if (params[key] != undefined || params[key] != null) {
                form.append(key, params[key])
            }
        });
        return fetch(`${Configs.SERVER_ROOT}/app/create`, {
            method: 'POST',
            body: form
        }).then(rs => rs.json())
    }
}
const service = new Service()
export default service