import React, { useState, useRef, useMemo, useEffect, MutableRefObject } from 'react';
import { useSpring, animated } from 'react-spring'
import View from './View';
import StyleSheet from './StyleSheet';

interface ModalProps {
    children?: any;
    zIndex?: number;
    transparent?: boolean;
    animated?:boolean;
    opacity?:number;
}

const Modal = React.memo((props: ModalProps) => {

    const [animStyle,api] = useSpring(() => ({opacity: props.animated ? 0 : 1}))

    useEffect(() => {
        if(props.animated){
            api.start({opacity: props.opacity ?? 0.2});
        }
    },[])

    return (
        <animated.div style={animStyle}>
            <View className='f1 fill' style={[
                styles.container,
                { backgroundColor: props.transparent ? '#FFFFFF00' : '#000000', zIndex: props.zIndex },
            ]}>
                {props.children}
            </View>
        </animated.div>
    )
})
export default Modal;

const ModalEmitter = {
    generateKey: () => {
        let t = new Date().getTime().toString()
        return t.substring(5, t.length) + Math.round(Math.random() * 100000)
    },
    hide: (key: string) => {
        window.dispatchEvent(new CustomEvent('modal-remove', {
            detail: { key: key }
        }))
    },
    show: (modalValue: { key: string; props: ModalProps; renderModal: (props: any) => any; }) => {
        window.dispatchEvent(new CustomEvent('modal-append', {
            detail: modalValue
        }))
    }
}
const ModalProvider = React.memo(() => {

    const globalKey = useRef(2000);

    const modalsValue: MutableRefObject<any> = useRef({})
    const [modals, setModals]: [any, any] = useState({})
    const modalKeys = useMemo(() => Object.keys(modals), [modals])

    /**
     * {
     *  key: string;
     *  props: any;
     *  renderModal: (props) => any;
     * }
     */
    useEffect(() => {
        window.addEventListener('modal-append', (event: any) => {
            let detail = event.detail;
            let existedSameKey = false;
            for (let k of Object.keys(modalsValue.current)) {
                if (modalsValue.current[k].key === detail.key) {
                    existedSameKey = true;
                    break;
                }
            }
            if (!existedSameKey) {
                globalKey.current += 1;
                detail.props = Object.assign(detail.props, { zIndex: globalKey.current })
                modalsValue.current[detail.key] = detail;
                setModals(modalsValue.current)
            }
        })
        window.addEventListener('modal-remove', (event: any) => {
            let detail = event.detail;
            if (modalsValue.current[detail.key] !== undefined) {
                delete modalsValue.current[detail.key]
            }
            setModals(modalsValue.current)
        })
    }, [])

    return (
        <>
            {modalKeys.map((item: any,index:number) => {
                let modalProps = modals[item].props;
                return modals[item].renderModal(Object.assign(modalProps,{key: index.toString()}));
            })}
        </>
    )
});

export { ModalProvider, ModalEmitter }
const styles = StyleSheet.create({
    container: {
        position: 'absolute',
        top: 0,
        right: 0,
        bottom: 0,
        left: 0
    }
})