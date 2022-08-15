package handler

import (
	"fmt"
	"math"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
	"yayar/internal/conf"
	"yayar/internal/data"
	"yayar/internal/data/ent"
	"yayar/internal/data/ent/app"
	"yayar/internal/data/ent/version"
	"yayar/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	Data *data.Data
	Conf *conf.AppConf
	Tmp  *string
	Cos  *utils.TTCos
}

type AppHandlerOptions func(handler *AppHandler)

func (handler *AppHandler) InitAppHandler(opts ...AppHandlerOptions) {
	for _, opt := range opts {
		opt(handler)
	}
}
func InitAppHandlerWithData(data *data.Data) AppHandlerOptions {
	return func(handler *AppHandler) {
		handler.Data = data
	}
}
func InitAppHandlerWithConf(conf *conf.AppConf) AppHandlerOptions {
	return func(handler *AppHandler) {
		handler.Conf = conf
	}
}
func InitAppHandlerWithTmpDir(dir string, absolutePath *string) AppHandlerOptions {
	return func(handler *AppHandler) {
		if absolutePath != nil {
			handler.Tmp = absolutePath
			os.MkdirAll(*absolutePath, 0775)
		} else if dir != "" {
			pwd, _ := os.Getwd()
			tmpdir := path.Join(pwd, dir)
			handler.Tmp = &tmpdir
			os.MkdirAll(tmpdir, 0775)
		}
	}
}
func InitAppHandlerWithCOS(per *utils.TTCOSConfig, tmp *utils.TTCOSTmpConfig) AppHandlerOptions {
	return func(handler *AppHandler) {
		ttcos := utils.TTCos{}
		if per != nil {
			ttcos.Init(utils.InitWithPersistKeys(*per))
			handler.Cos = &ttcos
		} else if tmp != nil {
			ttcos.Init(utils.InitWithTempKeys(*tmp))
			handler.Cos = &ttcos
		}
	}
}

//判断COS是否初始化完成
func (handler *AppHandler) isCosClientValide() bool {
	return handler != nil && handler.Cos.Client != nil
}

func (handler *AppHandler) genAppVersionDirPath(bundleID, version string) string {
	return path.Join(
		strings.TrimPrefix(handler.Conf.COS_RES_DIR, "/"),
		strings.ReplaceAll(bundleID, ".", "_"),
		strings.ReplaceAll(version, ".", "_"),
	)
}

//生成、上传plist文件
// return 访问链接、对象路径、错误
func (handler *AppHandler) genAndUploadPlist(ctx *gin.Context, ipa, bundleID, version, name string) (string, string, error) {
	plistContent := utils.Plist.GenPlistContent(ipa, bundleID, version, name)
	appVerDir := handler.genAppVersionDirPath(bundleID, version)
	tmpAppVerDir := path.Join(*handler.Tmp, appVerDir)
	os.MkdirAll(tmpAppVerDir, 0775)
	plistPath := path.Join(tmpAppVerDir, "manifest.plist")
	plistFile, _ := os.Create(plistPath)
	plistFile.WriteString(plistContent)
	plistFile.Close()
	objectPath := path.Join(appVerDir, "manifest.plist")
	result, _, err := handler.Cos.UploadObject(
		ctx,
		objectPath,
		plistPath,
		nil)
	os.Remove(plistPath)
	if err != nil {
		return "", "", err
	}
	return result.Location, objectPath, nil
}

//移动COS对象文件
func (handler *AppHandler) moveObjectFile(ctx *gin.Context, sourceURL, bundleID, version string) (string, string, error) {
	//例如: distribute/app/com_xxx/1_0_0
	appVerDir := handler.genAppVersionDirPath(bundleID, version)
	surl, _ := url.Parse(sourceURL)
	// sourcebucket-appid.cos.region.myqcloud/sourcekey
	hos := surl.Host //对应访问地址 https://xxx.xxx.myqcloud.com
	pat := surl.Path //对应objectPath、即cos上的存储路径 /tmp/xxx.apk
	spath := strings.SplitN(pat, "/", -1)
	filename := spath[len(spath)-1]
	objectPath := path.Join(appVerDir, filename)
	_, err := handler.Cos.MoveObject(ctx, sourceURL, pat, objectPath)
	if err != nil {
		return "", "", err
	}
	return path.Join("https://", hos, objectPath), objectPath, nil
}

//查询单个App信息
func (handler *AppHandler) App(ctx *gin.Context) {

}

//发布App
//Method: POST
//
func (handler *AppHandler) Create(ctx *gin.Context) {
	maxReserveVer := handler.Conf.MAX_RESERVE_VER
	type CreateForm struct {
		//必传参数: ID 或 Name,Logo,BundleID
		ID       int    `form:"id"`
		Name     string `form:"name"`
		Logo     string `form:"logo"`
		BundleID string `form:"bundle_id"`
		//version fields
		//必传参数: Version , Build , Size
		Version     string  `form:"version"`
		Build       int     `form:"build"`
		Size        float64 `form:"size"`
		IpaURL      string  `form:"ipa_url"`
		ApkURL      string  `form:"apk_url"`
		PlistURL    string  `form:"plist_url"`
		Description string  `form:"description"`
		Access      string  `form:"access"`
		AccessCode  string  `form:"access_code"`
	}
	client := handler.Data.Client
	var form CreateForm
	resp := Response{}
	if err := ctx.Bind(&form); err == nil {
		var a *ent.App
		if form.ID != 0 {
			//存在App ID, 获取 App数据
			a, err = client.App.Get(ctx, form.ID)
			if err != nil && !ent.IsNotFound(err) {
				resp.Serve(ctx, WithError(err.Error()))
				return
			}
		} else {
			//根据bundle_id查询App
			a, err = client.App.Query().Where(app.BundleIDEQ(form.BundleID)).First(ctx)
			if err != nil && !ent.IsNotFound(err) {
				resp.Serve(ctx, WithError(err.Error()))
				return
			}
		}
		//不存在App ID 或无法根据ID查询到对应 App, 新创建App记录以及版本
		if form.BundleID == "" || form.Name == "" {
			resp.Serve(ctx, WithError("App名称、包名不能为空"))
			return
		}
		//创建App
		if a == nil {
			logo, _, err := handler.moveObjectFile(ctx, form.Logo, form.BundleID, form.Version)
			if err != nil {
				resp.Serve(ctx, WithError(err.Error()))
				return
			}
			a, err = client.App.Create().SetName(form.Name).SetBundleID(form.BundleID).SetLogo(logo).Save(ctx)
			if err != nil {
				resp.Serve(ctx, WithError(err.Error()))
				return
			}
		}
		//创建版本
		if form.Version == "" || form.Build == 0 || form.Size == 0 {
			resp.Serve(ctx, WithError("发布版本必须传入正确的版本号、编译号和包大小"))
			return
		}
		if form.ApkURL == "" && form.IpaURL == "" {
			resp.Serve(ctx, WithError("发布版本必须传入安卓或iOS应用下载地址"))
			return
		}
		reg, _ := regexp.Compile(`^([0-9]+)\.([0-9]+)\.([0-9]+)$`)
		if !reg.MatchString(form.Version) {
			resp.Serve(ctx, WithError("版本号格式错误、必须遵循 x.x.x 格式"))
			return
		}
		a_vers, err := client.Version.Query().Where(version.HasAppWith(app.IDEQ(a.ID))).Order(ent.Desc(version.FieldVersion)).All(ctx)
		if err != nil {
			resp.Serve(ctx, WithError(err.Error()))
			return
		}
		needDeleteLast := false
		if len(a_vers) > 0 {
			if form.Version > a_vers[0].Version {
				//版本号有效、可以发布版本
				if len(a_vers) >= maxReserveVer {
					needDeleteLast = true
				} //可以直接发布最新版本, needDeleteLast = false
			} else if form.Version == a_vers[0].Version {
				if form.Build > a_vers[0].Build {
					needDeleteLast = true
				} else if form.Build < a_vers[0].Build {
					resp.Serve(ctx, WithError("版本需要大于已上传的最大版本、版本号相同时需要编译号大于已上传的版本号对应编译号"))
					return
				} else {
					//上传的应用版本号等于数据库中最新版本、判断是否是上传不同平台的安装包
					var uperr string = ""
					var platform string = ""
					if a_vers[0].ApkURL != "" && form.IpaURL != "" {
						platform = "ios"
					} else if a_vers[0].IpaURL != "" && form.ApkURL != "" {
						platform = "android"
					} else {
						uperr = "请提高版本号、发布新版本来代替替换操作"
					}
					if uperr != "" {
						resp.Serve(ctx, WithError(uperr))
					} else {
						tx, _ := client.BeginTx(ctx, nil)
						upt := tx.Version.UpdateOneID(a_vers[0].ID)
						if platform == "ios" {
							ipaurl, _, err := handler.moveObjectFile(ctx, form.IpaURL, form.BundleID, form.Version)
							if err != nil {
								resp.Serve(ctx, WithError(err.Error()))
								return
							}
							upt.SetIpaURL(ipaurl).SetIpaSize(form.Size)
							plistURL, objectPath, err := handler.genAndUploadPlist(ctx, form.IpaURL, form.BundleID, form.Version, form.Name)
							if err != nil {
								resp.Serve(ctx, WithError(err.Error()))
								return
							}
							if strings.Contains(handler.Conf.COS_CDN_URL, "https") {
								//使用CDN加速地址进行访问
								plistURL = path.Join(handler.Conf.COS_CDN_URL, objectPath)
							}
							upt.SetPlistURL(plistURL)
						} else {
							apkurl, _, err := handler.moveObjectFile(ctx, form.ApkURL, form.BundleID, form.Version)
							if err != nil {
								resp.Serve(ctx, WithError(err.Error()))
								return
							}
							upt.SetApkURL(apkurl).SetApkSize(form.Size)
						}
						if form.Access == "public" || form.Access == "private" {
							upt.SetAccess(form.Access)
						}
						upt.SetAccessCode(form.AccessCode)
						v, err := upt.Save(ctx)
						if err != nil {
							tx.Rollback()
							resp.Serve(ctx, WithError(err.Error()))
						} else {
							tx.Commit()
							resp.Serve(ctx, WithData(v, nil))
						}
					}
					return
				}
			} else {
				resp.Serve(ctx, WithError("版本需要大于已上传的最大版本、版本号相同时需要编译号大于已上传的版本号对应编译号"))
				return
			}
		}
		tx, _ := client.BeginTx(ctx, nil)
		if needDeleteLast {
			tx.Version.DeleteOne(a_vers[len(a_vers)-1]).Exec(ctx)
		}
		crt := tx.Version.Create().SetVersion(form.Version).SetBuild(int(math.Max(1, float64(form.Build)))).SetAccessCode(form.AccessCode).SetDescription(form.Description).AddApp(a)
		if form.Access == "public" || form.Access == "private" {
			crt.SetAccess(form.Access)
		}
		regAppUrl, _ := regexp.Compile(`^https://(.*)\.(apk|ipa)$`)
		if form.ApkURL != "" {
			//发布安卓应用
			if regAppUrl.MatchString(form.ApkURL) {
				apkurl, _, err := handler.moveObjectFile(ctx, form.ApkURL, form.BundleID, form.Version)
				if err != nil {
					resp.Serve(ctx, WithError(err.Error()))
					return
				}
				crt.SetApkURL(apkurl).SetApkSize(form.Size)
			} else {
				resp.Serve(ctx, WithError("安卓应用下载链接异常、非APK下载链接"))
				return
			}
		} else {
			//发布iOS应用
			if regAppUrl.MatchString(form.IpaURL) {
				ipaurl, _, err := handler.moveObjectFile(ctx, form.IpaURL, form.BundleID, form.Version)
				if err != nil {
					resp.Serve(ctx, WithError(err.Error()))
					return
				}
				crt.SetIpaURL(ipaurl).SetIpaSize(form.Size)
			} else {
				resp.Serve(ctx, WithError("苹果应用下载链接异常、非ipa下载链接"))
				return
			}
			//生成 manifest.plist 文件、上传到cos
			// https://xxx.xx/app/com_xxx/1_0_0/manifest.plist
			plistURL, objectPath, err := handler.genAndUploadPlist(ctx, form.IpaURL, form.BundleID, form.Version, form.Name)
			if err != nil {
				resp.Serve(ctx, WithError(err.Error()))
				return
			}
			if strings.Contains(handler.Conf.COS_CDN_URL, "https") {
				//使用CDN加速地址进行访问
				plistURL = path.Join(handler.Conf.COS_CDN_URL, objectPath)
			}
			crt.SetPlistURL(plistURL)
		}
		ver, err := crt.Save(ctx)
		if err != nil {
			tx.Rollback()
			resp.Serve(ctx, WithError(err.Error()))
		} else {
			tx.Commit()
			resp.Serve(ctx, WithData(ver, nil))
		}
	} else {
		resp.Serve(ctx, WithError("Params Error"))
	}
}

//更新某个App的信息
func (handler *AppHandler) Update(ctx *gin.Context) {

}

//获取App列表
//Method: GET
func (handler *AppHandler) List(ctx *gin.Context) {
	resp := Response{}
	apps, err := handler.Data.Client.App.Query().Order(ent.Desc(app.FieldUpdatedAt)).All(ctx)
	if err != nil {
		resp.Serve(ctx, WithError(err.Error()))
		return
	}
	resp.Serve(ctx, WithData(apps, nil))
}

//获取App的版本列表
//Method: GET
//Params: id => int[app的id]
func (handler *AppHandler) ListVersions(ctx *gin.Context) {
	resp := Response{}
	id_str := ctx.Query("id")
	id, err := strconv.Atoi(id_str)
	if id == 0 || err != nil {
		resp.Serve(ctx, WithError("App不存在"))
		return
	}
	vers, err := handler.Data.Client.Version.Query().Where(version.HasAppWith(app.IDEQ(id))).All(ctx)
	if err != nil {
		resp.Serve(ctx, WithError(err.Error()))
		return
	}
	resp.Serve(ctx, WithData(vers, nil))
}

//重定向到下载页
//Method: GET
//"/install/:name/*version"
func (handler *AppHandler) Install(ctx *gin.Context) {
	// resp := Response{}
	reg, _ := regexp.Compile(`^([0-9]+)\.([0-9]+)\.([0-9]+)$`)
	n := strings.ReplaceAll(ctx.Param("name"), "_", ".")
	v := strings.ReplaceAll(ctx.Param("version"), "_", ".")
	v = strings.TrimPrefix(v, "/")
	//是否指定版本
	useVer := reg.MatchString(v)
	//是否使用bundle id作为app查询依据
	useBun := strings.Contains(n, ".")
	appqy := handler.Data.Client.App.Query()
	if useBun {
		appqy = appqy.Where(app.BundleIDEQ(n))
	} else {
		appqy = appqy.Where(app.NameEQ(n))
	}
	_, err := appqy.First(ctx)
	if err != nil {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"message": fmt.Sprintf("The Target App(%s) May Not Exist", n),
		})
		return
	}
	verqy := handler.Data.Client.Version.Query()
	if useVer {
		verqy = verqy.Where(version.VersionEQ(v))
	}
	ver, err := verqy.Order(ent.Desc(version.FieldVersion)).First(ctx)
	if err != nil {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"message": fmt.Sprintf("The Target App(%s) Version(%s) May Not Exist", n, v),
		})
		return
	}
	//正式环境使用模板html来代替重定向方式
	if handler.Conf.APP_MODE == "debug" {
		ctx.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("http://localhost:3333?id=%d", ver.ID))
	} else {
		ctx.HTML(http.StatusOK, "downloader.html", gin.H{
			"id": ver.ID,
		})
	}
}

//获取下载页的详细信息
//Method: GET
//Params: id => int[对应version的id]
func (handler *AppHandler) InstallDetail(ctx *gin.Context) {
	id_str := ctx.Query("id")
	id, err := strconv.Atoi(id_str)
	resp := Response{}
	if err != nil {
		resp.Serve(ctx, WithError("Invalid Param ID"))
		return
	}
	type AppInstallURLS struct {
		Ios     string `json:"ios"`
		Android string `json:"android"`
	}
	type AppInstallData struct {
		Name      string          `json:"name"`
		Logo      string          `json:"logo"`
		Version   string          `json:"version"`
		Build     int             `json:"build"`
		Access    string          `json:"access"`
		UpdatedAt time.Time       `json:"updatedAt"`
		URLS      *AppInstallURLS `json:"urls,omitempty"`
	}
	ver, err := handler.Data.Client.Version.Query().Where(version.IDEQ(id)).WithApp().First(ctx)
	if err != nil {
		resp.Serve(ctx, WithError(err.Error()))
		return
	}
	installData := AppInstallData{
		Name:      ver.Edges.App[0].Name,
		Logo:      ver.Edges.App[0].Logo,
		Version:   ver.Version,
		Build:     ver.Build,
		Access:    ver.Access,
		UpdatedAt: ver.UpdatedAt,
		URLS:      nil,
	}
	if ver.Access == "public" {
		installData.URLS = &AppInstallURLS{
			Android: ver.ApkURL,
			Ios:     ver.PlistURL,
		}
	}
	resp.Serve(ctx, WithData(installData, nil))
}

//校对访问码
//Method: GET
//Params: code => string , id => int[对应version的id]
func (handler *AppHandler) VerifyAccessCode(ctx *gin.Context) {
	code := ctx.Query("code")
	id_str := ctx.Query("id")
	id, err := strconv.Atoi(id_str)
	resp := Response{}
	if err != nil {
		resp.Serve(ctx, WithError("Invalid Param ID"))
		return
	}
	ver, err := handler.Data.Client.Version.Get(ctx, id)
	if err != nil {
		resp.Serve(ctx, WithError(err.Error()))
		return
	}
	if ver.AccessCode == code {
		resp.Serve(ctx, WithData(struct {
			Ios     string `json:"ios"`
			Android string `json:"android"`
		}{
			Ios:     ver.PlistURL,
			Android: ver.ApkURL,
		}, nil))
		return
	}
	resp.Serve(ctx, WithError("密码错误"))
}

//删除App
func (handler *AppHandler) Delete(ctx *gin.Context) {

}
