import React, { useMemo } from 'react';
import '../css/text.css';
import useChildren from '../hooks/useChildren';
import useRestProps from '../hooks/useRestProps';
import useViewStyles from '../hooks/useViewStyles';

interface TextProps extends Omit<React.HTMLProps<HTMLParagraphElement>, 'style' | 'fontSize'> {
    style?: React.CSSProperties | React.CSSProperties[];
    fontSize?: number;
    numOfLines?: number;
    disableSelectable?: boolean;
}
const Text = React.forwardRef((props: TextProps, ref: any) => {

    const fontSize = props.fontSize ?? 15;
    const numOfLines = props.numOfLines;

    const restProps = useRestProps(props, ['children', 'style', 'numOfLines', 'disableSelectable']);
    const children = useChildren(props.children)
    const viewStyles = useViewStyles(props.style, [
        {
            fontSize: `${fontSize}px`,
            // lineHeight: `${fontSize * 1.2}px`,
            // maxHeight: numOfLines === undefined ? undefined : `${fontSize * 1.2 * numOfLines}px`,
            // minHeight: `${fontSize * 1.2 * (numOfLines ?? 1)}px`,
            overflow: 'hidden',
            textOverflow: 'ellipsis',
            WebkitBoxOrient: 'vertical',
            WebkitLineClamp: numOfLines,
            cursor: 'default'
        }
    ])

    const className = useMemo(() => {
        return 'text ' + (props.className ?? '')
    }, [props.className])

    return <p ref={ref} className={className} style={viewStyles} {...restProps}>{children}</p>
})
export default Text;