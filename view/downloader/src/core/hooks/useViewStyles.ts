import React, { useMemo } from 'react';
import StyleSheet from '../components/StyleSheet';

export default function useViewStyles(style:any,defaultStyles?:React.CSSProperties[]) {

    const viewStyles = useMemo(() => {
        if (Array.isArray(style)) {
            return StyleSheet.combine([...(defaultStyles ?? []),...(style as React.CSSProperties[])])
        } else if (style !== undefined) {
            return StyleSheet.combine([...(defaultStyles ?? []),style]);
        }
        return defaultStyles === undefined ? undefined : StyleSheet.combine(defaultStyles);
    }, [style,defaultStyles])

    return viewStyles;
}