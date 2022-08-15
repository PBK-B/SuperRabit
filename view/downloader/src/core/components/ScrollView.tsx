import React, { useEffect, useCallback, useState, useLayoutEffect, useMemo, useRef } from 'react';
import useChildren from '../hooks/useChildren';
import useViewStyles from '../hooks/useViewStyles';
import _ from 'lodash';
import Visible from './Visible';
interface ContentViewProps {
    children?: any;
    contentStyles: any;
    onContentSizeChanged: any;
    onScroll?: (offset: number) => void;
    onEndReached?: () => void;
    onEndReachThreshold?: number;
    onMomentumScrollEnd: (offset: number) => void;
    enableMomentumScrollListener?: boolean;
}
interface ScrollViewProps extends Omit<ContentViewProps, 'contentStyles' | 'onContentSizeChanged' | 'onMomentumScrollEnd'> {
    style?: React.CSSProperties;
    contentStyle?: React.CSSProperties;
    horizontal?: boolean;
    disableAdaptiveWidth?: boolean;
    onContentSizeChanged?: (size: { width: number, height: number }) => void;
    onMomentumScrollEnd?: (event: {
        offset: number;
        contentSize: { width: number, height: number };
        containerSize: { width: number, height: number };
    }) => void;
}

const ScrollView = React.forwardRef((props: ScrollViewProps, ref) => {

    const { onContentSizeChanged } = props;
    const cnref: any = useRef(); // ScrollView Container Div reference
    const isHorizontal = props.horizontal ?? false;
    /**
     * 是否禁用ScrollView容器宽度跟随浏览器窗口大小变化。
     * 在横向滚动时需要指定容器的宽度、从而使得内容可滚动。因此浏览器窗口变化时ScrollView
     * 容器宽度也需要跟随变化、才能使得滚动视图时刻保持窗口宽度比。
     */
    const disableAdaptiveWidth = props.disableAdaptiveWidth ?? false;

    const children = useChildren(props.children)

    const [windowRect, setWindowRect] = useState({ width: window.innerWidth, height: window.innerHeight })
    const [containerSize, setContainerRect]: [any, any] = useState({});
    const containerSizeValue = useRef({ width: 0, height: 0 })
    const contentSizeValue = useRef({ width: 0, height: 0 })
    const isContentVisible = useMemo(() => containerSize.height !== undefined, [containerSize])

    const containerStyle: React.CSSProperties = useMemo(() => ({
        overflowX: isHorizontal ? 'scroll' : 'hidden',
        overflowY: isHorizontal ? 'hidden' : 'scroll',
        maxHeight: containerSize.height,
        maxWidth: containerSize.width,
    }), [containerSize, isHorizontal])
    const contentStyle: React.CSSProperties = useMemo(() => {
        return ({
            display: 'flex',
            overflow: 'scroll',
            flexDirection: isHorizontal ? 'row' : 'column',
            height: isHorizontal ? '100%' : undefined,
            width: isHorizontal ? containerSize.width : '100%',
        })
    }, [containerSize, isHorizontal])
    const styles = useViewStyles(props.style, [containerStyle])
    const contentStyles = useViewStyles(props.contentStyle, [contentStyle])

    const _onContentSizeChanged = useCallback(({ width, height }: { width: number, height: number }) => {
        onContentSizeChanged && onContentSizeChanged({ width, height })
        contentSizeValue.current = { width, height };
    }, [onContentSizeChanged])

    useLayoutEffect(() => {
        if (cnref.current) {
            let nrect = {
                width: cnref.current.clientWidth,
                height: cnref.current.clientHeight,
            };
            if (containerSizeValue.current != nrect) {
                containerSizeValue.current = nrect;
                setContainerRect(nrect)
            }
        }
        if (!disableAdaptiveWidth && props.style?.width === undefined) {
            window.onresize = () => {
                setWindowRect({ width: window.innerWidth, height: window.innerHeight })
            }
        }
        // eslint-disable-next-line
    }, [])

    useEffect(() => {
        if (typeof containerSize.width === 'number' && containerSize.width !== windowRect.width) {
            let c = { ...containerSize };
            c.width = windowRect.width;
            setContainerRect(c);
        }
        // eslint-disable-next-line
    }, [windowRect])

    const _onMomentumScrollEnd = (offset: number) => {
        if (props.onMomentumScrollEnd) {
            props.onMomentumScrollEnd({
                offset,
                contentSize: contentSizeValue.current,
                containerSize: containerSizeValue.current
            })
        }
    }

    return (
        <div ref={cnref} className="f1 fill col-fs-fs" style={styles}>
            <Visible show={isContentVisible}>
                <ContentView
                    contentStyles={contentStyles}
                    onContentSizeChanged={_onContentSizeChanged}
                    onScroll={props.onScroll}
                    onEndReached={props.onEndReached}
                    onEndReachThreshold={props.onEndReachThreshold}
                    onMomentumScrollEnd={_onMomentumScrollEnd}
                    enableMomentumScrollListener={props.onMomentumScrollEnd !== undefined}
                >
                    {children}
                </ContentView>
            </Visible>
        </div>
    )
})

export default ScrollView;

const ContentView = (props: ContentViewProps) => {
    const ctref: any = useRef(); // ScrollView Content Div reference
    const { onContentSizeChanged, contentStyles } = props;
    const children = useMemo(() => props.children, [props.children])
    const onEndReachThreshold = props.onEndReachThreshold ?? 0;
    const onEndReachedRecord = useRef(false);
    const onScroll = props.onScroll ?? ((v) => {});

    useLayoutEffect(() => {
        if (ctref.current) {
            onContentSizeChanged({
                width: ctref.current.scrollWidth,
                height: ctref.current.scrollHeight
            })
        }
    }, [children, onContentSizeChanged])

    function _onScroll(t: any) {
        onScroll(t.target.scrollTop)
        if (props.onEndReached) {
            let diff = Math.floor(t.target.scrollHeight - t.target.scrollTop)
            if (diff <= (t.target.clientHeight + onEndReachThreshold)) { //距离满足了
                if (onEndReachedRecord.current === false) {
                    onEndReachedRecord.current = true;
                    props.onEndReached();
                }
            } else {//在触底阈值之上
                if (onEndReachedRecord.current) onEndReachedRecord.current = false;
            }
        }
    }
    const _onMomentumScrollEnd = _.debounce((event) => {
        props.onMomentumScrollEnd(ctref.current?.scrollTop ?? 0)
    }, 100)
    useEffect(() => {
        let tt: any;
        if (ctref.current) {
            if (props.onEndReached || props.onScroll) {
                ctref.current.addEventListener('scroll', _onScroll)
            }
            if (props.enableMomentumScrollListener) {
                ctref.current.addEventListener('scroll', _onMomentumScrollEnd)
            }
        } else {
            tt = setTimeout(() => {
                if (props.onEndReached || props.onScroll) {
                    ctref.current.addEventListener('scroll', _onScroll)
                }
                if (props.enableMomentumScrollListener) {
                    ctref.current.addEventListener('scroll', _onMomentumScrollEnd)
                }
                clearTimeout(tt);
            }, 200);
        }
        return () => {
            if (tt) clearTimeout(tt);
            if (props.onScroll && ctref.current) ctref.current.removeEventListener('scroll', _onScroll);
            if (props.enableMomentumScrollListener && ctref.current) ctref.current.removeEventListener('scroll', _onMomentumScrollEnd);
        }
    }, [])
    return (
        <div ref={ctref} className="scrollbar-hide" style={contentStyles}>
            {children}
        </div>
    )
}