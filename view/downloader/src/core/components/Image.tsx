import React, { MutableRefObject, useLayoutEffect, useMemo, useRef } from 'react';
import useChildren from '../hooks/useChildren';
import useRestProps from '../hooks/useRestProps';
import useViewStyles from '../hooks/useViewStyles';

interface ViewProps extends Omit<React.HTMLProps<HTMLDivElement>, 'style'> {
    style?: React.CSSProperties | React.CSSProperties[];
    resizeMode?: 'fill' | 'contain' | 'cover' | 'none' | 'scale-down';
    onLayout?: (event: { width: number, height: number, left: number, top: number }) => void;
}

const Image = React.forwardRef((props: ViewProps, ref?: any) => {

    const { children, resizeMode } = props;
    const vref: MutableRefObject<HTMLDivElement | any> = ref ?? React.createRef();
    const restProps = useRestProps(props, ['children', 'style', 'className', 'onLayout', 'resizeMode']);
    const _children = useChildren(children)
    const viewStyles = useViewStyles(props.style, [{ objectFit: resizeMode}])

    const className = useMemo(() => {
        return 'col-fs-fs ' + (props.className ?? '')
    }, [props.className])

    useLayoutEffect(() => {
        if (vref.current && props.onLayout) {
            props.onLayout({
                width: vref.current.clientWidth,
                height: vref.current.clientHeight,
                top: vref.current.offsetTop,
                left: vref.current.offsetLeft
            })
        }
    }, [])

    return React.createElement(
        'img',
        { ref: vref, className, ...restProps, style: viewStyles },
        ..._children
    )

})
export default Image;