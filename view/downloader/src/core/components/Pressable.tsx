import React, { useMemo,MutableRefObject,useLayoutEffect } from 'react';
import useChildren from '../hooks/useChildren';
import useRestProps from '../hooks/useRestProps';
import useViewStyles from '../hooks/useViewStyles';

interface ViewProps extends Omit<React.HTMLProps<HTMLDivElement>, 'style'> {
    style?: React.CSSProperties | React.CSSProperties[];
    onPress?: () => void;
    onDoublePress?: () => void;
    onLayout?: (event: {width: number, height: number,left:number,top:number}) => void;
}

const Pressable = React.forwardRef((props: ViewProps, ref?:any) => {
    
    const vref:MutableRefObject<HTMLDivElement|any> = ref ?? React.createRef();
    const restProps = useRestProps(props,['children', 'style','className', 'onPress', 'onDoublePress','onLayout']);
    const children = useChildren(props.children)
    const viewStyles = useViewStyles(props.style)

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
    },[])

    return React.createElement(
        'div',
        {
            ref: vref,
            className,
            onClick: props.onPress,
            onDoubleClick: props.onDoublePress, ...restProps, style: viewStyles
        },
        ...children
    )

})
export default Pressable;