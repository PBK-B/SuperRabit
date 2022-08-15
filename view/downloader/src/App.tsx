import React, { useEffect, useMemo, useRef, useState } from "react";
import { View, Text, Image, StyleSheet, ScrollView, Pressable, TextInputer } from "./core/components";
import Android from "./assets/android.png"
import Apple from "./assets/apple.png"
import TBG from "./assets/tbg.png"
import Configs from "./Configs";

const DownloadingTipCard = React.memo((props: {

}) => {
	const [bottom, setBottom] = useState('-40vh')
	const [opacity,setOpacity] = useState(0)

	useEffect(() => {
		setTimeout(() => {
			setBottom('0')
			setOpacity(1.0)
		}, 2000);
	}, [])
	return (
		<>
		<View style={[StyleSheet.absoluteFill,{
			zIndex:1,
			backgroundColor:'#00000055',
			opacity,
			transition: 'opacity 0.26s linear',
			}]}/>
			<View style={{
				position: 'absolute',
				zIndex: 2,
				bottom,
				left: 0,
				right: 0,
				height: '40vh',
				borderTopLeftRadius: 16,
				borderTopRightRadius: 16,
				backgroundColor: 'white',
				transition: 'bottom 0.38s ease-out',
				boxShadow: '0 -3px 12px #00000020',
			}}>

			</View>
		</>
	)
})

const PlatformIcon = React.memo((props: {
	size?: number;
	platform?: 'Apple' | 'Android';
	active?: boolean;
}) => {
	const { size = 23, platform = 'Android', active = true } = props
	const IconSize = useMemo(() => ({ width: size, height: size }), [size])
	const ImgSize = useMemo(() => ({ width: size * 0.5, height: size * 0.5 }), [size])
	return (
		<View style={[IconSize, { borderRadius: IconSize.height / 2, opacity: active ? 1.0 : 0.36 }, styles.icon]}>
			<Image src={platform == 'Android' ? Android : Apple} style={[ImgSize, { objectFit: 'contain' }]} />
		</View>
	)
})
const MIN = 60
const HOUR = MIN * 60
const DAY = HOUR * 24
function timeAgo(time: string) {
	let tt = new Date(time).getTime()
	let now = Date.now()
	let sec = (now - tt) / 1000 //过去多少秒了
	if (sec < MIN) {
		return `${sec}秒`
	} else if (sec < HOUR) {
		return `${Math.round(sec / MIN)}分钟`
	} else if (sec < DAY) {
		return `${Math.round(sec / HOUR)}小时`
	}
	return `${Math.round(sec / DAY)}天`
}
const TimeAgo = React.memo((props: {
	time: string;
}) => {
	const { time } = props
	const [v, setV] = useState(timeAgo(time) + '以前')
	useEffect(() => {
		var t = setInterval(() => {
			setV(timeAgo(time) + '以前')
		}, 1000)
		return () => {
			clearTimeout(t)
		}
	}, [])
	return <Text style={[styles.sub, { marginLeft: 10 }]}>更新时间: {v}</Text>
})

const App = React.memo((props: any) => {

	const [isLocked, setLocked] = useState(false)
	const [error, setError] = useState<any>(null)
	const [data, setData] = useState<any>(null)
	const [urls, setURLS] = useState<any>({})
	const isIosOk = useMemo(() => typeof (urls.ios) == 'string' && urls.ios.length > 0, [urls])
	const isAndroidOk = useMemo(() => typeof (urls.android) == 'string' && urls.android.length > 0, [urls])

	useEffect(() => {
		const global:any = globalThis
		let url = new URL(window.location.href)
		let id = global.id
		if (id == null || id == undefined) {
			setError("404")
		} else {
			fetch(`${Configs.SERVER_ROOT}/app/installdetail?id=${id}`)
				.then(rs => rs.json())
				.then((rs: any) => {
					console.log(rs);
					if (rs.message) {
						setError(rs.message)
					} else {
						if (rs.data.access != 'public') {
							setLocked(true)
						}
						if (rs.data.urls) {
							setURLS(rs.data.urls)
						}
						setData(rs.data)
					}
				})
				.catch((e: any) => {
					setError(e.message)
				})
		}
	}, [])

	const downloadIOS = () => {
		if (isIosOk) {
			let sv = "itms-services:///?action=download-manifest&url=" + urls.ios
			window.open(sv, '_self')
		}
	}
	const downloadAndroid = () => {
		if (isAndroidOk) {
			window.open(urls.android, '_blank')
		}
	}
	if (error != null) {
		return (
			<View className="fill" style={styles.page}>
				<Text style={styles.err_label}>Oops! Error Happens</Text>
				<Text style={[styles.err_label, { fontSize: 18, fontWeight: 'normal' }]}>{error ?? "Unknown Reason"}</Text>
			</View>
		)
	}
	return (
		<>
			{/* <DownloadingTipCard /> */}
			<View className="fill" style={styles.page}>
				<Image src={TBG} style={styles.topbg} />
				{data != null && (
					<>
						<Image
							src={data.logo}
							style={styles.logo}
						/>
						<View style={[styles.center, { margin: '15px 0' }]}>
							<Text style={styles.name}>{data.name}</Text>
							<PlatformIcon platform="Apple" active={isIosOk} />
							<PlatformIcon platform="Android" active={isAndroidOk} />
						</View>
						<View style={[styles.center]}>
							<Text style={styles.sub}>版本: {data.version}(build:{data.build})</Text>
							<TimeAgo time={data.updatedAt} />
						</View>
						<View style={[styles.center, { marginTop: 55, height: 36 }]}>
							{
								isLocked ? (
									<>
										<View style={{ width: 160, height: 32, backgroundColor: '#f3f3f5' }}>
											<TextInputer
												placeholder="输入访问密码"
												style={{ flex: 1, margin: '0 8px', color: '#222', fontSize: 14 }}
											/>
										</View>
										<Pressable style={{ padding: '0 15px' }}>
											<Text style={styles.btn_label2}>解锁</Text>
										</Pressable>
									</>
								) : (
									<>
										<Pressable onPress={downloadIOS} style={[styles.btn, { opacity: isIosOk ? 1.0 : 0.38 }]}>
											<Text style={styles.btn_label}>下载</Text>
											<Image src={Apple} style={styles.btn_icon} />
										</Pressable>
										<Pressable onPress={downloadAndroid} style={[styles.btn, { opacity: isAndroidOk ? 1.0 : 0.38 }]}>
											<Text style={styles.btn_label}>下载</Text>
											<Image src={Android} style={styles.btn_icon} />
										</Pressable>
									</>
								)
							}
						</View>
					</>
				)}
			</View >
		</>
	)
})
export default App;

const styles = StyleSheet.create({
	page: {
		justifyContent: 'center',
		alignItems: 'center',
	},
	topbg: {
		width: '100vw',
		height: 30,
		position: 'absolute',
		top: 0,
		left: 0, right: 0
	},
	logo: {
		width: 110,
		height: 110,
		borderRadius: 26,
		overflow: 'hidden',
		backgroundColor: 'white',
		boxShadow: '0px 3px 12px #00000066'
	},
	icon: {
		justifyContent: 'center',
		alignItems: 'center',
		backgroundColor: '#353538',
		margin: '0 5px'
	},
	name: {
		fontSize: 24,
		color: '#353538',
		fontWeight: '500'
	},
	sub: {
		color: '#788099',
		fontSize: 14,
	},
	center: {
		flexDirection: 'row',
		justifyContent: 'center',
		alignItems: 'center'
	},
	btn: {
		width: 100,
		height: 36,
		borderRadius: 18,
		flexDirection: 'row',
		justifyContent: 'center',
		alignItems: 'center',
		backgroundColor: '#F74783',
		margin: '0 6px'
	},
	btn_label: {
		fontSize: 15,
		color: '#FFFFFF',
		cursor: 'pointer'
	},
	btn_label2: {
		fontSize: 15,
		fontWeight: 'bold',
		color: '#F74783',
		cursor: 'pointer'
	},
	pwd_tip: {
		color: '#455055',
		fontSize: 14,
	},
	btn_icon: {
		width: 16,
		height: 16,
		objectFit: 'contain',
		marginLeft: 6
	},
	err_label: {
		display: 'flex',
		color: '#788099',
		fontSize: '7vw',
		fontWeight: '500',
		marginBottom: 20,
	}
})