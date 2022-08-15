import { useState, useEffect,useRef } from 'react';


export function useAdaptive(defaultValue: Record<string, number>) {
	// MediaQuery Matches (max-width: 1440px) :  true => //less than 1440px
	//Base 1920
	//Stop scaling when width's less than 1440px
	//Eg: num => 1200
	// greater than 1440 => `(1200*100/1920)%`		(v * 100/1920)%
	// less eq than 1440 => 1200 * 1440/1920		(v * 1440/1920)
	const media = useRef(window.matchMedia("(max-width: 1440px)")).current;
	const dimensionsValue = useRef(adopt(media.matches));
	const [dimensions, setDimensions] = useState(adopt(media.matches));
	// const lastMatches = useRef(false);
	function updateDimensions(dims: any) {
		dimensionsValue.current = dims;
		setDimensions(dims)
	}
	function adopt(matches: boolean) {
		let keys = Object.keys(defaultValue);
		let vs: any = {};
		if (matches) {
			for (let k of keys) {
				vs[k] = defaultValue[k] * 1440 / 1920
			}
		} else {
			for (let k of keys) {
				vs[k] = `${defaultValue[k] * 100 / 1920}vw`
			}
		}
		return vs;
	}
	useEffect(() => {
		media.addEventListener("change", (ev) => {
			
			let keys = Object.keys(defaultValue);
			if (keys.length == 0) return;
			if (ev.matches) {
				//USE => (v * 1440/1920)
				if (defaultValue[keys[0]] === dimensionsValue.current[0]) {
					//匹配上MediaQuery、但是值相等、不重复计算
					return;
				} else {
					updateDimensions(adopt(true));
				}
			} else {
				//USE => (v * 100/1920)%
				updateDimensions(adopt(false));
			}
		})
	}, [])
	return dimensions;
}

export function useFontAdaptive(defaultValue: Record<string, number>) {
	const media = useRef(window.matchMedia("(max-width: 1440px)")).current;
	const dimensionsValue = useRef(adopt(media.matches));
	const [dimensions, setDimensions] = useState(adopt(media.matches));
	function updateDimensions(dims: any) {
		dimensionsValue.current = dims;
		setDimensions(dims)
	}
	function adopt(matches: boolean) {
		let keys = Object.keys(defaultValue);
		let vs: any = {};
		if (matches) {
			//小于等于1440 用固定尺寸; 64px * 1440/1920
			for (let k of keys) {
				vs[k] = defaultValue[k] * 1440 / 1920
			}
		} else {
			for (let k of keys) {
				vs[k] = `calc(${defaultValue[k]} * 100vw / 1920)`
			}
		}
		return vs;
	}
	useEffect(() => {
		media.addEventListener("change", (ev) => {
			let keys = Object.keys(defaultValue);
			if (keys.length == 0) return;
			if (ev.matches) {
				//USE => (v * 1440/1920)
				updateDimensions(adopt(true));
			} else {
				//USE => (v * 100/1920)%
				updateDimensions(adopt(false));
			}
		})
	}, [])
	return dimensions;
}